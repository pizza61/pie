package modules

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/globalsign/mgo"
	"github.com/pizza61/pie/utilities"
)

func HandlePoints(s *discordgo.Session, m *discordgo.MessageCreate, c *mgo.Database) {
	user, err := utilities.FindUser(c, m.Author.ID)
	if err != nil {
		if err.Error() == "not found" {
			// uzytkownik nie istnieje, dodaj uzytkownika i serwer
			errA := utilities.RegisterUser(c, m.Author.ID, m.Author.Username)
			if errA != nil {

			} else {
				guild, err := utilities.GetGuild(s, m.ChannelID)
				utilities.CheckErr(err)

				err = utilities.RegisterMember(c, m.Author.ID, guild.ID)
				utilities.CheckErr(err)
			}
		} else {
			// wystapil powazny blad
			fmt.Println(err.Error())
		}
	} else {
		guild, err := utilities.GetGuild(s, m.ChannelID)
		utilities.CheckErr(err)

		if _, ok := user.Guilds[guild.ID]; !ok { // jezeli u uzytkwnika nie ma tego serwera
			err = utilities.RegisterMember(c, m.Author.ID, guild.ID)
			utilities.CheckErr(err)
		} else { // serwer u użytkownika istnieje, dodaj mu punkty
			err = utilities.AddPoints(c, m.Author.ID, guild.ID)
			utilities.CheckErr(err)
		}

	}
}
