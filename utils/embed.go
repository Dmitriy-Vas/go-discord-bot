package utils

import (
	dgo "github.com/bwmarrin/discordgo"
	"time"
)

const (
	BOT_NAME    string = "VasBot"
	PROJECT_URL string = "https://github.com/Dmitriy-Vas/go-discord-bot"
	COLOR       int    = 16098851
)

type Embed struct {
	*dgo.MessageEmbed
}

func CreateEmbed(title string, description string) *Embed {
	me := &dgo.MessageEmbed{
		Author:      &dgo.MessageEmbedAuthor{Name: BOT_NAME, URL: PROJECT_URL},
		Color:       COLOR, // todo add random color
		Timestamp:   time.Now().Format(time.RFC3339),
		Title:       title,
		Description: description,
	}
	e := &Embed{me}
	return e
}

func (e *Embed) SetTitle(text string) {
	e.Title = text
}

func (e *Embed) SetDescription(text string) {
	e.Description = text
}

func (e *Embed) SetFooter(text string) {
	e.Footer = &dgo.MessageEmbedFooter{Text: text}
}

func (e *Embed) SetColor(color int) {
	e.Color = color
}

func (e *Embed) SetThumbnail(url string) {
	e.Thumbnail = &dgo.MessageEmbedThumbnail{URL: url}
}

func (e *Embed) SetThumbnailProperly(url string, width int, height int) {
	e.Thumbnail = &dgo.MessageEmbedThumbnail{URL: url, Width: width, Height: height}
}

func (e *Embed) SetImage(url string) {
	e.Image = &dgo.MessageEmbedImage{URL: url}
}

func (e *Embed) SetImageProperly(url string, width int, height int) {
	e.Image = &dgo.MessageEmbedImage{URL: url, Width: width, Height: height}
}

func (e *Embed) SetVideo(url string) {
	e.Video = &dgo.MessageEmbedVideo{URL: url}
}

func (e *Embed) SetVideoProperly(url string, width int, height int) {
	e.Video = &dgo.MessageEmbedVideo{URL: url, Width: width, Height: height}
}

func (e *Embed) SetFields(fields map[string]string, inline bool) {
	for key, value := range fields {
		field := &dgo.MessageEmbedField{Name: key, Value: value, Inline: inline}
		e.Fields = append(e.Fields, field)
	}
}

func getRandomColor() int {
	return 0
}
