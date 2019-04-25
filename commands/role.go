package commands

import (
	. "../models"
	"../utils"
	"fmt"
	"regexp"
)

const (
	ROLE_ADDED_MSG   string = "Got new role!"
	ROLE_REMOVED_MSG string = "Lost his role!"
	MANAGE_ROLE_PERM int    = 1 << 26
)

func Role(m *Message) {
	has, err := utils.HasPermissions(m, m.Author, MANAGE_ROLE_PERM)
	if err != nil || !has {
		fmt.Println("Don't have rights to manage roles")
		return
	}

	mentions := m.MessageCreate.Mentions
	if len(mentions) == 0 {
		return
	}

	r := regexp.MustCompile("(\\w+)$")
	matches := r.FindAllStringSubmatch(m.MessageCreate.Content, 1)
	if matches == nil || matches[0] == nil {
		return
	}

	// Get role from name
	role, err := utils.RoleFromName(m, matches[0][1])
	if err != nil {
		fmt.Println(err)
		return
	}

	embed := utils.CreateEmbed("Success!", "Role: "+role.Name)

	gID := m.MessageCreate.GuildID

	stat := make(map[string]string)

	for _, user := range mentions {
		// Check user for role
		has, err = utils.HasRole(m, user, role)
		if err != nil {
			fmt.Println(err)
			return
		}
		if has {
			// Remove role from the user if he have this role
			err = m.GuildMemberRoleRemove(gID, user.ID, role.ID)
			if err != nil {
				fmt.Println(err)
				return
			}
			stat[user.Username] = ROLE_REMOVED_MSG
		} else {
			// Add role to the user if he don't have this role
			err = m.GuildMemberRoleAdd(gID, user.ID, role.ID)
			if err != nil {
				fmt.Println(err)
				return
			}
			stat[user.Username] = ROLE_ADDED_MSG
		}
	}

	// Check for oversize.
	if len(mentions) > 25 {
		embed.SetDescription("Over 25 users got (or lost) role " + role.Name + "!")
	} else {
		embed.SetFields(stat, true)
	}

	_, err = m.ChannelMessageSendEmbed(m.MessageCreate.ChannelID, embed.MessageEmbed)
	if err != nil {
		fmt.Println(err)
	}
}
