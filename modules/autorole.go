package modules

import (
	"github.com/bwmarrin/discordgo"
	"github.com/globalsign/mgo"
	"github.com/pizza61/pie/utilities"
)

func Autorole(s *discordgo.Session, m *discordgo.GuildMemberAdd, c *mgo.Database) {
	// sprawdź czy serwer ma autorole
	guildD, err := utilities.FindGuild(c, m.GuildID)
	if err != nil {
		utilities.LogErr(err.Error())
		return
	}
	if len(guildD.AutoroleID) > 0 {
		// nadaj rolę użytkownikowi
		err = s.GuildMemberRoleAdd(m.GuildID, m.User.ID, guildD.AutoroleID)
		if err != nil {
			utilities.LogErr("Na serwerze " + m.GuildID + ": " + err.Error())
			return
		}
	}

}
