package main

//#region Header
import (
	"./config"
	"./handlers"
	"fmt"
	dgo "github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	bot Bot
)

type Bot struct {
	Session *dgo.Session
	User    *dgo.User
}

//#endregion

func isErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Initialize and connect bot to discord
func (d *Bot) createBot() {
	tokenPrefix := ""
	if !config.Data.User {
		tokenPrefix = "Bot "
	}

	// Create new session
	s, err := dgo.New(tokenPrefix + os.Getenv("TOKEN"))
	isErr(err)
	s.State.MaxMessageCount = 100
	d.Session = s

	// Fetch information about bot
	u, err := s.User("@me")
	isErr(err)
	d.User = u

	fmt.Println("Logged in as", u.Username)

	// Open connection to discord
	err = s.Open()
	isErr(err)
}

// Add event handlers (listeners) to the bot
func (d *Bot) addHandlers() {
	d.Session.AddHandler(handlers.MessageCreated)
	d.Session.AddHandler(handlers.MemberAdded)
	d.Session.AddHandler(handlers.MemberRemoved)
}

// Entry point
func main() {
	config.LoadConfig()

	// Initialize bot
	bot.createBot()
	bot.addHandlers()

	// Catch interrupt signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	// Close connection to discord
	bot.Session.Close()
	os.Exit(0)
}
