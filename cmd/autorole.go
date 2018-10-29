package cmd

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/globalsign/mgo"
	"github.com/pizza61/pie/utilities"
)

func Autorole(s *discordgo.Session, m *discordgo.MessageCreate, c *mgo.Database, args []string) {
	ifek, err := utilities.HasAdministrator(s, m.ChannelID, m.Author.ID)
	if err != nil {
		return
	}
	if ifek {
		gd, err := utilities.GetGuild(s, m.ChannelID)
		utilities.CheckErr(err)
		gildia, err := utilities.FindGuild(c, gd.ID)
		utilities.CheckErr(err)
		if len(args) > 1 {
			if args[1] == "off" {
				gildia.AutoroleID = ""
				utilities.SaveGuild(c, gildia, gd.ID)
				e := utilities.NewEmbed().SetInfo().SetTitle("Usunięto autorole")
				s.ChannelMessageSendEmbed(m.ChannelID, e.Generate())
			} else {
				ullmsg := strings.Join(args[1:], " ")
				for _, el := range gd.Roles {
					if strings.ToLower(el.Name) == strings.ToLower(ullmsg) {
						// mamy juz role teraz ja sprawdzmy
						err = s.GuildMemberRoleAdd(gd.ID, s.State.User.ID, el.ID)
						if err != nil {
							e := utilities.NewEmbed().SetError().SetTitle("Nie możesz ustawić tej roli").SetDescription("Rola musi być pod rolą bota, bot musi posiadać również uprawnienie Zarządzanie rolami")
							s.ChannelMessageSendEmbed(m.ChannelID, e.Generate())
							return
						} else {
							s.GuildMemberRoleRemove(gd.ID, s.State.User.ID, el.ID)
							// widac bez problemow wiec ustawiamy
							gildia.AutoroleID = el.ID
							err = utilities.SaveGuild(c, gildia, gd.ID)
							if err == nil {
								e := utilities.NewEmbed().SetInfo().SetTitle("Pomyślnie ustawiono autorole!")
								s.ChannelMessageSendEmbed(m.ChannelID, e.Generate())
								return
							}
						}
					}
				}
				e := utilities.NewEmbed().SetError().SetTitle("Nie odnaleziono roli!")
				s.ChannelMessageSendEmbed(m.ChannelID, e.Generate())
				// nie znaleziono nie ustawiono
			}
		} else {
			e := utilities.NewEmbed().SetInfo().SetTitle("Autorole").AddField(true, "Aby ustawić nową wpisz", "*autorole nazwa-roli").AddField(true, "Aby wyłączyć wpisz", "*autorole off")
			if len(gildia.AutoroleID) > 0 {
				role, err := utilities.FindRole(gd, gildia.AutoroleID)
				if err == nil {
					e = e.AddField(true, "Aktualna", role.Name)
					s.ChannelMessageSendEmbed(m.ChannelID, e.Generate())
				}
			} else {
				s.ChannelMessageSendEmbed(m.ChannelID, e.Generate())
			}

		}
	} else {
		e := utilities.NewEmbed().SetError().SetTitle("Nie masz uprawnień")
		s.ChannelMessageSendEmbed(m.ChannelID, e.Generate())
	}
}
