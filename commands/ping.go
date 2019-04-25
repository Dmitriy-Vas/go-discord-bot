package commands

import (
	. "../models"
	"../utils"
	"fmt"
	"time"
)

func Ping(m *Message) {
	timestamp, err := m.MessageCreate.Timestamp.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}

	ping := time.Since(timestamp)
	embed := utils.CreateEmbed("Pong!", fmt.Sprintf("Your ping: %v", ping))

	_, err = m.Session.ChannelMessageSendEmbed(m.MessageCreate.ChannelID, embed.MessageEmbed)
	if err != nil {
		fmt.Println(err)
	}
}
