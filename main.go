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

func (d *Bot) createBot() {
	// Create new session
	s, err := dgo.New(config.Data.Token)
	isErr(err)
	d.Session = s

	// Fetch information about bot
	u, err := s.User("@me")
	isErr(err)
	d.User = u

	fmt.Println("Logged as", u.Username)

	// Open connection to discord
	err = s.Open()
	isErr(err)
}

func (d *Bot) addHandlers() {
	d.Session.AddHandler(handlers.MessageCreated)
	d.Session.AddHandler(handlers.MemberAdded)
	d.Session.AddHandler(handlers.MemberRemoved)
}

/**
Entry point
*/
func main() {
	config.LoadConfig()

	// Initialize bot
	bot.createBot()
	bot.addHandlers()

	// Close connection to discord at the end
	defer bot.Session.Close()

	// Catch interrupt signal and save config
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	config.SaveConfig()
	os.Exit(0)
}
