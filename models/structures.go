package models

import (
	dgo "github.com/bwmarrin/discordgo"
	"time"
)

type Message struct {
	*dgo.Session
	*dgo.MessageCreate
}

type URL struct {
	Link    string
	Queries map[string]string
}

type Project struct {
	Name        string    `json:"name"`
	FullName    string    `json:"full_name"`
	Link        string    `json:"html_url"`
	Description string    `json:"description"`
	Fork        bool      `json:"fork"`
	Created     time.Time `json:"created_at"`
	Updated     time.Time `json:"updated_at"`
	Stars       int       `json:"stargazers_count"`
	Watchers    int       `json:"watchers_count"`
	Language    string    `json:"language"`
}

type User struct {
	Login        string    `json:"login"`
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Company      string    `json:"company"`
	Blog         string    `json:"blog"`
	Location     string    `json:"location"`
	Description  string    `json:"bio"`
	Repositories int       `jon:"public_repos"`
	Gists        int       `json:"public_gists"`
	Followers    int       `json:"Followers"`
	Following    int       `json:"Following"`
	Created      time.Time `json:"created_at"`
	Updated      time.Time `json:"updated_at"`
}
