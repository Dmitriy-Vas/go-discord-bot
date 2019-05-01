package commands

import (
	. "../models"
	"../utils"
	"fmt"
	"os"
)

func Help(m *Message) {
	prefix := os.Getenv("PREFIX")
	list := "Here is a list of commands below.\n" +
		prefix + "help - shows list of commands.\n" +
		prefix + "ping - send message ping.\n" +
		prefix + "clear [0-100] [user] - removes specified amount of last messages (or from mentioned user) in the current channel.\n" +
		prefix + "role [user]+ [role] - add or remove users role.\n" +
		prefix + "projects [word] - shows list of projects from specified GitHub username.\n" +
		prefix + "stars [word] - shows how many stars developer have in its GitHub projects"

	embed := utils.CreateEmbed("Help", list)

	_, err := m.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
	if err != nil {
		fmt.Println(err)
	}
}
