package cmd

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/globalsign/mgo"
	"github.com/pizza61/pie/utilities"
)

func Set(s *discordgo.Session, m *discordgo.MessageCreate, c *mgo.Database, args []string) {
	ifek, err := utilities.HasAdministrator(s, m.ChannelID, m.Author.ID)
	if err != nil {
		return
	}
	if ifek {
		if len(args) == 4 {
			// komenda w formie *set message/work 10 100
			guildd, err := utilities.GetGuild(s, m.ChannelID)
			if err != nil {
				return
			}
			guild, err := utilities.FindGuild(c, guildd.ID)
			if err != nil {
				return
			}
			if args[1] == "message" {
				min, err := strconv.Atoi(args[2])
				if err != nil {
					e := utilities.NewEmbed().SetError().SetTitle("Podałeś nieprawidłową liczbę!")
					s.ChannelMessageSendEmbed(m.ChannelID, e.Generate())
					return
				}
				max, err := strconv.Atoi(args[3])
				if err != nil {
					e := utilities.NewEmbed().SetError().SetTitle("Podałeś nieprawidłową liczbę!")
					s.ChannelMessageSendEmbed(m.ChannelID, e.Generate())
					return
				}
				if (min > 0 && min < 1000000) && (max > 0 && max < 1000000) && (max > min) {
					guild.PMessage.Min = min
					guild.PMessage.Max = max
					utilities.SaveGuild(c, guild, guildd.ID)
					e := utilities.NewEmbed().SetInfo().SetTitle("Ustawiono pomyślnie!")
					s.ChannelMessageSendEmbed(m.ChannelID, e.Generate())
				} else {
					e := utilities.NewEmbed().SetError().SetTitle("Wystąpił problem z twoimi liczbami!").SetDescription("Minimalna i maksymalna liczba musi być większa od zera i mniejsza od miliona, a minimalna nie może być większa od maksymalnej!")
					s.ChannelMessageSendEmbed(m.ChannelID, e.Generate())
					return
				}
				//s.ChannelMessageSend(m.ChannelID, "message")
			} else if args[1] == "work" {
				//s.ChannelMessageSend(m.ChannelID, "work")
				min, err := strconv.Atoi(args[2])
				if err != nil {
					e := utilities.NewEmbed().SetError().SetTitle("Podałeś nieprawidłową liczbę!")
					s.ChannelMessageSendEmbed(m.ChannelID, e.Generate())
					return
				}
				max, err := strconv.Atoi(args[3])
				if err != nil {
					e := utilities.NewEmbed().SetError().SetTitle("Podałeś nieprawidłową liczbę!")
					s.ChannelMessageSendEmbed(m.ChannelID, e.Generate())
					return
				}
				if (min > 0 && min < 1000000) && (max > 0 && max < 1000000) && (max > min) {
					guild.PWork.Min = min
					guild.PWork.Max = max
					utilities.SaveGuild(c, guild, guildd.ID)
					e := utilities.NewEmbed().SetInfo().SetTitle("Ustawiono pomyślnie!")
					s.ChannelMessageSendEmbed(m.ChannelID, e.Generate())
				} else {
					e := utilities.NewEmbed().SetError().SetTitle("Wystąpił problem z twoimi liczbami!").SetDescription("Minimalna i maksymalna liczba musi być większa od zera i mniejsza od miliona, a minimalna nie może być większa od maksymalnej!")
					s.ChannelMessageSendEmbed(m.ChannelID, e.Generate())
					return
				}
			} else {
				e := utilities.NewEmbed().SetWarn().SetTitle("Użycie komendy set").AddField(false, "Przeznaczenie", "Komenda set pozwala na ustawienie minimalnej oraz maksymalnej ilości pieniędzy przyznawanych za wiadomość lub za pracę")
				e = e.AddField(false, "Użycie komendy", "*set message/work min max")
				e = e.AddField(false, "np.", "*set message 1 10")
				s.ChannelMessageSendEmbed(m.ChannelID, e.Generate())
			}
		} else {
			e := utilities.NewEmbed().SetWarn().SetTitle("Użycie komendy set").AddField(false, "Przeznaczenie", "Komenda set pozwala na ustawienie minimalnej oraz maksymalnej ilości pieniędzy przyznawanych za wiadomość lub za pracę")
			e = e.AddField(false, "Użycie komendy", "*set message/work min max")
			e = e.AddField(false, "np.", "*set message 1 10")
			e = e.SetFooter("Aby zobaczyć aktualne wartości, użyj komendy settings")
			s.ChannelMessageSendEmbed(m.ChannelID, e.Generate())
		}
	} else {
		e := utilities.NewEmbed().SetError().SetTitle("Nie masz uprawnień")
		s.ChannelMessageSendEmbed(m.ChannelID, e.Generate())
	}
}
