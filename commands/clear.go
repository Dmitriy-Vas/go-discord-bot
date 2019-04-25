package commands

import (
	"../config"
	. "../models"
	"fmt"
	"regexp"
	"strconv"
)

func Clear(m *Message) {
	r := regexp.MustCompile("clear (\\d+)")
	matches := r.FindAllStringSubmatch(m.MessageCreate.Content, 1)
	if matches != nil && matches[0] != nil {
		limit, err := strconv.Atoi(matches[0][1])
		if err != nil {
			fmt.Println(err)
			return
		}

		id := ""
		if !config.Data.User {
			if len(m.Mentions) != 0 {
				id = m.Mentions[0].ID
			}
		} else {
			id = m.Author.ID
		}

		// Fetch messages up to the specified limit
		messages, err := m.Session.ChannelMessages(m.MessageCreate.ChannelID, limit, id, id, "")
		if err != nil {
			fmt.Println(err)
			return
		}

		// Save messages ID
		var buffer []string
		for _, message := range messages {
			// If program enabled in user-mode
			// Delete messages one-by-one
			if config.Data.User {
				err = m.Session.ChannelMessageDelete(m.MessageCreate.ChannelID, message.ID)
				if err != nil {
					fmt.Println(err)
					return
				}
			} else {
				buffer = append(buffer, message.ID)
			}
		}

		// Delete saved messages
		if !config.Data.User {
			err = m.Session.ChannelMessagesBulkDelete(m.MessageCreate.ChannelID, buffer)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
