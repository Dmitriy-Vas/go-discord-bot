package models

import (
	dgo "github.com/bwmarrin/discordgo"
	"time"
)

type Firebase struct {
	Type                    string `json:"type,omitempty"`
	ProjectID               string `json:"project_id,omitempty"`
	PrivateKeyID            string `json:"private_key_id,omitempty"`
	PrivateKey              string `json:"private_key,omitempty"`
	ClientEmail             string `json:"client_email,omitempty"`
	ClientID                string `json:"client_id,omitempty"`
	AuthURI                 string `json:"auth_uri,omitempty"`
	TokenURI                string `json:"token_uri,omitempty"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url,omitempty"`
	ClientX509CertURL       string `json:"client_x509_cert_url,omitempty"`
}

type DiscordUser struct {
	XP   int64  `json:"xp"`
	Rank string `json:"rank"`
}

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
