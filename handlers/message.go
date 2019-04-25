package handlers

//#region Header
import (
	"../commands"
	. "../models"
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
	if !strings.HasPrefix(m.Content, os.Getenv("PREFIX")) {
		return
	}
	// Check for disallowed channels
	channelType, err := channelType(s, m)
	if err != nil || !isAllowedChannelType(channelType) {
		return
	}
	// Get alias of command and try to send to the listener
	message := &Message{Session: s, MessageCreate: m}
	alias := commandAlias(message)
	if len(alias) != 0 {
		listener := listeners[alias]
		if listener != nil {
			go listener(message)
		}
	}
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
func channelType(s *dgo.Session, m *dgo.MessageCreate) (*dgo.ChannelType, error) {
	ch, err := s.State.Channel(m.ChannelID)
	if err != nil {
		if ch, err = s.Channel(m.ChannelID); err != nil {
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
