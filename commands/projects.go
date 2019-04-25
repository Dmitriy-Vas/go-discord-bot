package commands

import (
	. "../models"
	"../utils"
	"fmt"
	"regexp"
)

func Projects(m *Message) {
	r := regexp.MustCompile("([\\w\\d-]+)$")
	matches := r.FindAllStringSubmatch(m.MessageCreate.Content, 1)
	if matches == nil || matches[0] == nil {
		return
	}

	userName := matches[0][1]

	embed := utils.CreateEmbed("GitHub "+userName+" projects", "")

	projects, _ := utils.Projects(userName, 1)
	if len(projects) == 0 {
		embed.SetDescription("Projects not found!")
	} else {
		if len(projects) > 10 {
			embed.SetDescription("Shows last 5 projects")
		}
		projectsMap := make(map[string]string)
		for index, project := range projects {
			if index > 5 {
				break
			}
			projectsMap[project.Name] = getProjectField(project)
		}
		embed.SetFields(projectsMap, false)
	}

	_, err := m.Session.ChannelMessageSendEmbed(m.MessageCreate.ChannelID, embed.MessageEmbed)
	if err != nil {
		fmt.Println(err)
	}
}

func getProjectField(project *Project) string {
	out := "Full name: " + project.FullName + "\n" +
		"Link to the project: " + project.Link + "\n"
	if project.Language != "" {
		out += "Language: " + project.Language + "\n"
	}
	if project.Description != "" {
		out += "Description: `" + project.Description + "`\n"
	}
	if project.Stars > 0 {
		out += fmt.Sprintf("Stars: %d\n", project.Stars)
	}
	if project.Watchers > 0 {
		out += fmt.Sprintf("Watchers: %d\n", project.Watchers)
	}
	out += fmt.Sprintf("Forked: %v\nCreated: %v\nUpdated: %v\n", project.Fork, project.Created, project.Updated)
	return out
}
