package handlers

import (
	"../config"
	"../utils"
	"errors"
	"fmt"
	dgo "github.com/bwmarrin/discordgo"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Handler func(s *dgo.Session, m *dgo.MessageCreate)

var (
	handlers map[string]Handler
)

// Log error
func isErr(err error, args ...string) {
	if err != nil {
		log.Println(err, args)
	}
}

func isAllowedChannelType(channelType *dgo.ChannelType) bool {
	switch *channelType {
	case dgo.ChannelTypeDM:
	case dgo.ChannelTypeGroupDM:
		return false
	}
	return true
}

func channelType(s *dgo.Session, m *dgo.MessageCreate) (*dgo.ChannelType, error) {
	ch, err := s.State.Channel(m.ChannelID)
	if err != nil {
		if ch, err = s.Channel(m.ChannelID); err != nil {
			return nil, err
		}
	}
	return &ch.Type, nil
}

func commandAlias(m *dgo.MessageCreate) (alias string) {
	alias = ""
	r := regexp.MustCompile("^>(\\w+)")
	matches := r.FindAllStringSubmatch(m.Content, -1)
	if matches != nil && matches[0] != nil {
		alias = matches[0][1]
	}
	return
}

func arguments(m *dgo.MessageCreate) []string {
	return strings.Split(m.Content, " ")
}

func getIntArg(index int, args []string) (int, error) {
	if len(args)-1 < index {
		return 0, errors.New("index out of range")
	}
	limit, err := strconv.Atoi(args[1])
	if err != nil {
		return 0, err
	}
	return limit, nil
}

func MessageCreated(s *dgo.Session, m *dgo.MessageCreate) {
	if !strings.HasPrefix(m.Content, config.Data.Prefix) {
		return
	}
	channelType, err := channelType(s, m)
	if err != nil || !isAllowedChannelType(channelType) {
		return
	}

	alias := commandAlias(m)
	if alias != "" {
		handlers[alias](s, m)
	}
}

func pong(s *dgo.Session, m *dgo.MessageCreate) {
	e := utils.CreateEmbed("Pong", "")
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, e.MessageEmbed)
	isErr(err)
}

func del(s *dgo.Session, m *dgo.MessageCreate) {
	limit, err := getIntArg(1, arguments(m))
	if err != nil {
		return
	}
	messages, _ := s.ChannelMessages(m.ChannelID, limit, "", "", "")
	var messagesBuffer []string
	for _, mess := range messages {
		messagesBuffer = append(messagesBuffer, mess.ID)
	}
	err = s.ChannelMessagesBulkDelete(m.ChannelID, messagesBuffer)
	isErr(err)
}

func init() {
	handlers = make(map[string]Handler)

	handlers["ping"] = Handler(pong)
	handlers["del"] = Handler(del)

	fmt.Println("Loaded", len(handlers), "commands")
}
