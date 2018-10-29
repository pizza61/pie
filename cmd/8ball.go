package cmd

import (
	"math/rand"
	"strings"
	"time"

	"github.com/pizza61/pie/utilities"

	"github.com/bwmarrin/discordgo"
	"github.com/globalsign/mgo"
)

func B8ball(s *discordgo.Session, m *discordgo.MessageCreate, c *mgo.Database, args []string) {
	now := time.Now()
	s1 := rand.NewSource(now.UnixNano())

	r1 := rand.New(s1)
	answers := []string{
		"Jeszcze jak!",
		"Raczej nie",
		"Nie wiem",
		"Chyba tak",
		"Tak",
		"Nie",
	}
	fullmsg := strings.Join(args[1:], " ")
	b8embed := utilities.NewEmbed().SetTitle(fullmsg).SetDescription(answers[r1.Intn(6)]).SetInfo()
	s.ChannelMessageSendEmbed(m.ChannelID, b8embed.Generate())
}
