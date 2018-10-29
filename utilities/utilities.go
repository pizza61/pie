package utilities

import (
	"errors"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func RandomBetween(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max+1-min) + min
}

// GetGuild by channel id
func GetGuild(s *discordgo.Session, channel string) (*discordgo.Guild, error) {
	ch, err := s.Channel(channel)
	if err != nil {
		return &discordgo.Guild{}, errors.New(err.Error())
	}
	gd, err := s.Guild(ch.GuildID)
	if err != nil {
		return &discordgo.Guild{}, errors.New(err.Error())
	}
	return gd, nil
}

func FindRole(guild *discordgo.Guild, role string) (*discordgo.Role, error) {
	for _, el := range guild.Roles {
		if el.ID == role {
			return el, nil
		}
	}
	return &discordgo.Role{}, errors.New("not found")
}
func HasAdministrator(s *discordgo.Session, channel string, userID string) (bool, error) {
	gd, err := GetGuild(s, channel)
	if err != nil {
		return false, err
	}
	mem, err := s.GuildMember(gd.ID, userID)
	if err != nil {
		return false, err
	}
	for _, el := range mem.Roles {
		rola, err := FindRole(gd, el)
		if err != nil {
			return false, err
		}
		if rola.Permissions&8 != 0 {
			return true, nil
		}
	}

	/*for _, el := range mem.Roles {
		role, _ := gd.Roles[]
		if role.Permissions&8 != 0 {
			return true
		}
	}*/
	return false, nil
}
