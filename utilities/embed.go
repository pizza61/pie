package utilities

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

type RichEmbed struct {
	AuthorName  string
	AuthorURL   string
	Title       string
	Color       int
	Description string
	FooterText  string
	FooterIcon  string
	ImageURL    string
	Fields      []*discordgo.MessageEmbedField
}

func NewEmbed() RichEmbed {
	newembd := RichEmbed{
		AuthorName: "Pie",
		AuthorURL:  "https://cdn.discordapp.com/avatars/486535712879935488/ee2eede9813957d99ce99ce06bbc591f.png?size=64",
	}
	return newembd
}

func (embd RichEmbed) SetTitle(title string) RichEmbed {
	embd.Title = title
	return embd
}

func (embd RichEmbed) SetInfo() RichEmbed {
	now := time.Now()
	s1 := rand.NewSource(now.UnixNano())

	r1 := rand.New(s1)
	infoColors := []int{0x0b5591, 0x6014b7, 0x2057a0, 0x951fa0}
	embd.Color = infoColors[r1.Intn(4)]
	return embd
}

func (embd RichEmbed) SetWarn() RichEmbed {
	embd.Color = 0xfff000
	return embd
}

func (embd RichEmbed) SetError() RichEmbed {
	embd.Color = 0xba2828
	return embd
}

func (embd RichEmbed) SetDescription(descriptio string) RichEmbed {
	embd.Description = descriptio
	return embd
}

func (embd RichEmbed) SetFooter(params ...string) RichEmbed {
	if len(params) == 1 {
		embd.FooterText = params[0]
		return embd
	} else if len(params) == 2 {
		embd.FooterText = params[0]
		embd.FooterIcon = params[1]
		return embd
	} else {
		return embd
	}
}

func (embd RichEmbed) SetImage(link string) RichEmbed {
	embd.ImageURL = link
	return embd
}

func (embd RichEmbed) AddField(inline bool, params ...string) RichEmbed {
	field := &discordgo.MessageEmbedField{
		Name:   params[0],
		Value:  params[1],
		Inline: inline,
	}
	embd.Fields = append(embd.Fields, field)
	return embd
}

func (embd RichEmbed) Generate() *discordgo.MessageEmbed {
	embed := discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    embd.AuthorName,
			IconURL: embd.AuthorURL,
		},
		Title:       embd.Title,
		Color:       embd.Color,
		Description: embd.Description,
		Footer: &discordgo.MessageEmbedFooter{
			Text: embd.FooterText, IconURL: embd.FooterIcon,
		},
		Fields: embd.Fields,
		Image: &discordgo.MessageEmbedImage{
			URL: embd.ImageURL,
		},
	}
	return &embed
}
