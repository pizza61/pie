package cmd

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/globalsign/mgo"
	"github.com/pizza61/pie/utilities"
)

func Prank(s *discordgo.Session, m *discordgo.MessageCreate, c *mgo.Database, args []string) {
	fullmsg := strings.Join(args[1:], " ")
	s.ChannelMessageSend(m.ChannelID, ":joy: :ok_hand: | "+utilities.Reverse(fullmsg))
	err := s.ChannelMessageDelete(m.ChannelID, m.ID)
	utilities.CheckErr(err)
}
