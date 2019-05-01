package commands

import (
	. "../models"
	"../utils"
	"fmt"
	"regexp"
)

func Stars(m *Message) {
	r := regexp.MustCompile("([\\w\\d-]+)$")
	matches := r.FindAllStringSubmatch(m.MessageCreate.Content, 1)
	if matches == nil || matches[0] == nil {
		return
	}

	userName := matches[0][1]

	embed := utils.CreateEmbed("GitHub "+userName+" stars", "")

	stars, status := utils.Stars(userName)
	if status != 200 {
		embed.SetDescription("Projects not found!")
	} else {
		embed.SetDescription(fmt.Sprintf("User has %d stars", stars))
	}

	_, err := m.ChannelMessageSendEmbed(m.MessageCreate.ChannelID, embed.MessageEmbed)
	if err != nil {
		fmt.Println(err)
	}
}
