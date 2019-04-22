package handlers

// todo move few functions to util
// todo rewrite handlers (functions) / and maybe split them?

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

func guildRoles(s *dgo.Session, m *dgo.MessageCreate) []*dgo.Role {
	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		if guild, err = s.Guild(m.GuildID); err != nil {
			return nil
		}
	}
	return guild.Roles
}

func findRole(s *dgo.Session, roles []*dgo.Role, role string) *dgo.Role {
	for _, r := range roles {
		if r.Name != role {
			continue
		}
		return r
	}
	return nil
}

func hasPermissions(s *dgo.Session, m *dgo.MessageCreate, user *dgo.User, permissions int) (bool, error) {
	perm, err := s.UserChannelPermissions(user.ID, m.ChannelID)
	if err != nil {
		return false, err
	}
	if perm&permissions != 0 {
		return false, nil
	}
	return true, nil
}

func hasRole(s *dgo.Session, m *dgo.MessageCreate, user *dgo.User, role *dgo.Role) (bool, error) {
	member, err := s.State.Member(m.GuildID, user.ID)
	if err != nil {
		if member, err = s.GuildMember(m.GuildID, user.ID); err != nil {
			return false, err
		}
	}
	for _, r := range member.Roles {
		if r != role.ID {
			continue
		}
		return true, nil
	}
	return false, nil
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

func notice(s *dgo.Session, m *dgo.MessageCreate) {
	e := utils.CreateEmbed("Notice", "By using this bot you agree that\n"+
		"U ar gud boi\n"+
		"No warranty is provided for using this bot and I (Dmitriy-Vas) disclaim responsibility for any damages that are caused as a result of using this bot")
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, e.MessageEmbed)
	isErr(err)
}

func setrole(s *dgo.Session, m *dgo.MessageCreate) {
	if len(m.Mentions) == 0 {
		return
	}
	args := arguments(m)
	if len(args) < 3 {
		return
	}
	roles := guildRoles(s, m)
	if roles == nil {
		return
	}
	role := findRole(s, roles, args[len(args)-1])
	if role == nil {
		return
	}
	for _, user := range m.Mentions {
		has, err := hasRole(s, m, user, role)
		if err != nil {
			return
		}
		if has {
			err = s.GuildMemberRoleRemove(m.GuildID, user.ID, role.ID)
			isErr(err)
		} else {
			err = s.GuildMemberRoleAdd(m.GuildID, user.ID, role.ID)
			isErr(err)
		}
	}
}

func ban(s *dgo.Session, m *dgo.MessageCreate) {
	if len(m.Mentions) == 0 {
		return
	}
	args := arguments(m)
	reason := ""
	time := 0
	switch {
	case len(args) > 3:
		reason = strings.Join(args[3:], " ")
		fallthrough
	case len(args) > 2:
		time, err := strconv.Atoi(args[2])
		if err == nil {
			time = time
		}
	}
	err := s.GuildBanCreateWithReason(m.GuildID, m.Mentions[0].ID, reason, time)
	isErr(err)
}

func init() {
	handlers = make(map[string]Handler)

	handlers["ping"] = Handler(pong)
	handlers["del"] = Handler(del)
	handlers["notice"] = Handler(notice)
	handlers["role"] = Handler(setrole)
	handlers["ban"] = Handler(ban)

	fmt.Println("Loaded", len(handlers), "commands")
}
