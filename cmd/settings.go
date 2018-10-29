package cmd

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/globalsign/mgo"
	"github.com/pizza61/pie/utilities"
)

func Settings(s *discordgo.Session, m *discordgo.MessageCreate, c *mgo.Database, args []string) {
	guildd, err := utilities.GetGuild(s, m.ChannelID)
	if err != nil {
		return
	}
	guild, err := utilities.FindGuild(c, guildd.ID)
	e := utilities.NewEmbed().SetInfo().SetTitle("Aktualne wartości")
	e = e.AddField(true, "Za wiadomość co minutę", "Od "+strconv.Itoa(guild.PMessage.Min)+" do "+strconv.Itoa(guild.PMessage.Max))
	e = e.AddField(true, "Za pracę", "Od "+strconv.Itoa(guild.PWork.Min)+" do "+strconv.Itoa(guild.PWork.Max))
	e = e.SetFooter("Aby zmienić te wartości, użyj komendy set")
	s.ChannelMessageSendEmbed(m.ChannelID, e.Generate())
}
