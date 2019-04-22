package utils

import (
	dgo "github.com/bwmarrin/discordgo"
	"time"
)

type Embed struct {
	MessageEmbed *dgo.MessageEmbed
}

func CreateEmbed(title string, description string) *Embed {
	me := &dgo.MessageEmbed{
		Author:      &dgo.MessageEmbedAuthor{Name: "VasBot"},
		Color:       16098851, // todo add random color
		Timestamp:   time.Now().Format(time.RFC3339),
		Title:       title,
		Description: description,
	}
	e := &Embed{MessageEmbed: me}
	return e
}

func (e *Embed) SetFooter(text string) {
	e.MessageEmbed.Footer = &dgo.MessageEmbedFooter{Text: text}
}

func (e *Embed) SetColor(color int) {
	e.MessageEmbed.Color = color
}

func (e *Embed) SetThumbnail(url string) {
	e.MessageEmbed.Thumbnail = &dgo.MessageEmbedThumbnail{URL: url}
}

func (e *Embed) SetThumbnailProperly(url string, width int, height int) {
	e.MessageEmbed.Thumbnail = &dgo.MessageEmbedThumbnail{URL: url, Width: width, Height: height}
}

func (e *Embed) SetImage(url string) {
	e.MessageEmbed.Image = &dgo.MessageEmbedImage{URL: url}
}

func (e *Embed) SetImageProperly(url string, width int, height int) {
	e.MessageEmbed.Image = &dgo.MessageEmbedImage{URL: url, Width: width, Height: height}
}

func (e *Embed) SetVideo(url string) {
	e.MessageEmbed.Video = &dgo.MessageEmbedVideo{URL: url}
}

func (e *Embed) SetVideoProperly(url string, width int, height int) {
	e.MessageEmbed.Video = &dgo.MessageEmbedVideo{URL: url, Width: width, Height: height}
}

func (e *Embed) SetFields(fields map[string]string, inline bool) {
	for key, value := range fields {
		field := &dgo.MessageEmbedField{Name: key, Value: value, Inline: inline}
		e.MessageEmbed.Fields = append(e.MessageEmbed.Fields, field)
	}
}

func getRandomColor() int {
	return 0
}
