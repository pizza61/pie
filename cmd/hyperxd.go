package cmd

import (
	"github.com/bwmarrin/discordgo"
	"github.com/globalsign/mgo"
)

func Hyperxd(s *discordgo.Session, m *discordgo.MessageCreate, c *mgo.Database, args []string) {
	s.ChannelMessageSend(m.ChannelID, "<:hyperxd:501054879742033950> n")
}
