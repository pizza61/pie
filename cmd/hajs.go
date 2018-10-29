package cmd

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/globalsign/mgo"
	"github.com/pizza61/pie/utilities"
)

func Hajs(s *discordgo.Session, m *discordgo.MessageCreate, c *mgo.Database, args []string) {
	if len(m.Mentions) > 0 {
		user, err := utilities.FindUser(c, m.Mentions[0].ID)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Wystąpił błąd!")
			return
		}
		guild, err := utilities.GetGuild(s, m.ChannelID)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Wystąpił błąd!")
			return
		}
		embd := utilities.NewEmbed()
		embd = embd.SetInfo().SetTitle("Stan konta użytkownika " + m.Mentions[0].Username + " na serwerze " + guild.Name).SetDescription(fmt.Sprintf("%d %s", user.Guilds[guild.ID].Points, "BTC"))
		s.ChannelMessageSendEmbed(m.ChannelID, embd.Generate())
	} else {
		user, err := utilities.FindUser(c, m.Author.ID)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Wystąpił błąd!")
			return
		}
		guild, err := utilities.GetGuild(s, m.ChannelID)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Wystąpił błąd!")
			return
		}

		embd := utilities.NewEmbed()
		embd = embd.SetInfo().SetTitle("Stan konta użytkownika " + m.Author.Username + " na serwerze " + guild.Name).SetDescription(fmt.Sprintf("%d %s", user.Guilds[guild.ID].Points, "BTC"))
		s.ChannelMessageSendEmbed(m.ChannelID, embd.Generate())
	}
}
