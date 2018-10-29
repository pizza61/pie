package cmd

import (
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/globalsign/mgo"
	"github.com/pizza61/pie/utilities"
)

func Work(s *discordgo.Session, m *discordgo.MessageCreate, c *mgo.Database, args []string) {
	guildD, err := utilities.GetGuild(s, m.ChannelID)
	guildM, err := utilities.FindGuild(c, guildD.ID)
	user, err := utilities.FindUser(c, m.Author.ID)
	if err != nil {
		return
	}
	if user.LastWork+86400 < time.Now().Unix() {
		if _, ok := user.Guilds[guildD.ID]; ok {
			randomix := utilities.RandomBetween(guildM.PWork.Min, guildM.PWork.Max)
			mapka := user.Guilds[guildD.ID]
			mapka.Points += randomix
			user.Guilds[guildD.ID] = mapka
			user.LastWork = time.Now().Unix()
			utilities.SaveUser(c, user, m.Author.ID)
			embd := utilities.NewEmbed().SetInfo().SetDescription("Pracowałeś w Lidlu i zarobiłeś " + strconv.Itoa(randomix)).SetFooter("Pracować możesz na jednym, wybranym serwerze co 24 godziny. Sprawdź, ile możesz zarobić komendą `settings`")
			s.ChannelMessageSendEmbed(m.ChannelID, embd.Generate())
		} else {
			embd := utilities.NewEmbed().SetError().SetTitle("Wystapił nieznany błąd")
			s.ChannelMessageSendEmbed(m.ChannelID, embd.Generate())
		}
	} else {
		poczekac := (user.LastWork + 86400 - time.Now().Unix()) / 3600
		embd := utilities.NewEmbed().SetError().SetDescription("Już pracowałeś! Musisz poczekać jeszcze około " + strconv.FormatInt(poczekac, 10) + " godzin")
		s.ChannelMessageSendEmbed(m.ChannelID, embd.Generate())
	}
}
