package handlers

import (
	"../utils"
	"fmt"
	dgo "github.com/bwmarrin/discordgo"
)

func MemberAdded(s *dgo.Session, m *dgo.GuildMemberAdd) {
	e := utils.CreateEmbed("Welcome!", "Hello! My name is VasBot.\n"+
		"Since I'm still in development, I may have some bugs and typo errors.\n"+
		"If you want to contribute, feel free to send pull request with any changes.\n +"+
		"Link to the repository located in the footer.")
	e.SetFooter("https://github.com/Dmitriy-Vas/go-discord-bot")

	// Get DM channel for specified user
	ch, err := s.UserChannelCreate(m.User.ID)
	if err != nil {
		fmt.Println(err)
	}

	// Send embed message to specified DM channel
	_, err = s.ChannelMessageSendEmbed(ch.ID, e.MessageEmbed)
	if err != nil {
		fmt.Println(err)
	}
}

func MemberRemoved(s *dgo.Session, m *dgo.GuildMemberRemove) {
	e := utils.CreateEmbed("Bye-Bye!", "I'm so sad that you leave our server.\n"+
		"I hope you'll return in someday.")

	// Get DM channel for specified user
	ch, err := s.UserChannelCreate(m.User.ID)
	if err != nil {
		fmt.Println(err)
	}

	// Send embed message to specified DM channel
	_, err = s.ChannelMessageSendEmbed(ch.ID, e.MessageEmbed)
	if err != nil {
		fmt.Println(err)
	}
}
