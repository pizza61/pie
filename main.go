package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/globalsign/mgo"
	"github.com/pizza61/pie/cmd"
	"github.com/pizza61/pie/modules"
	"github.com/pizza61/pie/utilities"
)

type Settings struct {
	Prefix string
	Token  string
}

func main() {
	jsondata, err := ioutil.ReadFile("settings.json")
	utilities.CheckErr(err)

	var settings Settings
	json.Unmarshal(jsondata, &settings)
	utilities.Log("Loading...")
	dg, err := discordgo.New("Bot " + settings.Token)
	utilities.CheckErr(err)

	// Database cośtamy
	dbsession, err := mgo.Dial("localhost")
	utilities.CheckErr(err)

	defer dbsession.Close()
	dbsession.SetMode(mgo.Monotonic, true)
	c := dbsession.DB("pie")

	// guilds
	guilds := c.C("guilds")

	indexGuilds := mgo.Index{
		Key:        []string{"guildid", "guildname", "autoroleid", "nougatcount", "msgcount", "items", "pwork", "pmessage"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = guilds.EnsureIndex(indexGuilds)
	utilities.CheckErr(err)

	// users
	users := c.C("users")

	indexUsers := mgo.Index{
		Key:        []string{"userid", "username", "lastwork", "guilds"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = users.EnsureIndex(indexUsers)
	utilities.CheckErr(err)

	//utilities.TestUser(c)

	dg.AddHandler(ready)
	dg.AddHandler(func(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
		modules.Autorole(s, m, c)
	})
	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) { // MessageCreate
		if m.Author.Bot {
			return
		}
		handleCmd(s, m, c, settings)

		modules.HandlePoints(s, m, c)
	})
	dg.AddHandler(func(s *discordgo.Session, g *discordgo.GuildCreate) { // GuildCreate
		_, err := utilities.FindGuild(c, g.ID)
		if err != nil { // serwer nie istnieje trzeba dodac
			if err.Error() == "not found" {
				utilities.AddGuild(c, g.ID, g.Name)
				utilities.Log("New guild: " + g.ID)
			} else {
				utilities.CheckErr(err)
			}
		}
	})

	err = dg.Open()
	utilities.CheckErr(err)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	fmt.Print("\n")
	utilities.Log("Closing...")
	dg.Close()
}

func ready(s *discordgo.Session, r *discordgo.Ready) {
	utilities.Log("Ready!")
}

func handleCmd(s *discordgo.Session, m *discordgo.MessageCreate, c *mgo.Database, settings Settings) {
	msg := m.Content
	if strings.HasPrefix(msg, settings.Prefix) {
		msg = strings.Trim(msg, settings.Prefix)
		args := strings.Split(msg, " ")
		//msg = strings.ToLower(msg)

		switch args[0] {

		case "8ball":
			cmd.B8ball(s, m, c, args)
		case "odwroc":
			cmd.Odwroc(s, m, c, args)
		case "prank":
			cmd.Prank(s, m, c, args)
		case "hyperxd":
			cmd.Hyperxd(s, m, c, args)
		case "hajs":
			cmd.Hajs(s, m, c, args)
		case "autorole":
			cmd.Autorole(s, m, c, args)
		case "set":
			cmd.Set(s, m, c, args)
		case "settings":
			cmd.Settings(s, m, c, args)
		case "work":
			cmd.Work(s, m, c, args)
		}
	}
}
