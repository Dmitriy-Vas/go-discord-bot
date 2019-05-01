package handlers

//#region Header
import (
	"../commands"
	. "../models"
	"../storage"
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	dgo "github.com/bwmarrin/discordgo"
	"os"
	"regexp"
	"strings"
)

type Command func(message *Message)

var (
	listeners map[string]Command
)
//#endregion

// Load all available commands to the map of listeners
func init() {
	listeners = make(map[string]Command)

	listeners["help"] = commands.Help
	listeners["ping"] = commands.Ping
	listeners["clear"] = commands.Clear
	listeners["role"] = commands.Role

	// GitHub commands
	listeners["projects"] = commands.Projects
	listeners["stars"] = commands.Stars

	fmt.Println("Loaded", len(listeners), "commands")
}

// Fire when someone create new message
func MessageCreated(s *dgo.Session, m *dgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	message := &Message{Session: s, MessageCreate: m}
	// Check for disallowed channels
	channelType, err := channelType(message)
	if err != nil || !isAllowedChannelType(channelType) {
		return
	}
	if err = increaseXP(message); err != nil {
		fmt.Println(err)
	}
	if !strings.HasPrefix(m.Content, os.Getenv("PREFIX")) {
		return
	}
	// Get alias of command and try to send to the listener
	alias := commandAlias(message)
	if len(alias) != 0 {
		listener := listeners[alias]
		if listener != nil {
			go listener(message)
		}
	}
}

func increaseXP(m *Message) error {
	ctx := context.Background()
	user, exists := storage.Users[m.Author.ID]
	if !exists {
		u := map[string]interface{}{
			"xp":   1,
			"rank": "newbie",
		}
		if _, err := storage.Client.Collection("users").Doc(m.Author.ID).Set(ctx, u); err != nil {
			return err
		}
		storage.Users[m.Author.ID] = &DiscordUser{
			XP:   1,
			Rank: "newbie",
		}
	} else {
		user.XP++
		u := map[string]interface{}{
			"xp": user.XP,
		}
		_, err := storage.Client.Collection("users").Doc(m.Author.ID).Set(ctx, u, firestore.MergeAll)
		if err != nil {
			return err
		}
	}
	return nil
}

// Check channel for inconsistency with DM and GroupDM types
func isAllowedChannelType(channelType *dgo.ChannelType) bool {
	switch *channelType {
	case dgo.ChannelTypeDM:
	case dgo.ChannelTypeGroupDM:
		return false
	}
	return true
}

// Returns ChannelType of channel
func channelType(m *Message) (*dgo.ChannelType, error) {
	ch, err := m.State.Channel(m.ChannelID)
	if err != nil {
		if ch, err = m.Channel(m.ChannelID); err != nil {
			return nil, err
		}
	}
	return &ch.Type, nil
}

// Returns alias of command
func commandAlias(m *Message) (out string) {
	content := m.MessageCreate.Content
	prefix := os.Getenv("PREFIX")
	r := regexp.MustCompile("^" + prefix + "(\\w+)")
	matches := r.FindAllStringSubmatch(content, 1)
	if matches != nil && matches[0] != nil {
		out = matches[0][1]
	}
	return
}
