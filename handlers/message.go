package handlers

import (
	"../config"
	dgo "github.com/bwmarrin/discordgo"
	"log"
	"regexp"
	"strings"
)

// Log error
func isErr(err error, args ...string) {
	if err != nil {
		log.Println(err, args)
	}
}

func getMatch(r string, m string) (match string, found bool) {
	reg, err := regexp.Compile(r)
	isErr(err)
	match = reg.FindString(m)
	found = reg.MatchString(m)
	return
}

// todo found best solution
func MessageCreated(s *dgo.Session, mc *dgo.MessageCreate) {
	if !strings.HasPrefix(mc.Content, config.Data.Prefix) {
		return
	}

	var f bool

	_, f = getMatch("ping", mc.Content)
	if f {
		pong(s, mc)
	}
}

func pong(s *dgo.Session, m *dgo.MessageCreate) {
	_, err := s.ChannelMessageSend(m.ChannelID, "Pong")
	isErr(err)
}
