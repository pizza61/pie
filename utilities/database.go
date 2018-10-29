package utilities

import (
	"errors"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Guild struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	GuildID     string
	GuildName   string
	AutoroleID  string
	NougatCount int
	MsgCount    int
	Items       []Item
	PWork       MinMax
	PMessage    MinMax
}

type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	UserID   string
	Username string
	LastWork int64
	Guilds   map[string]Member
}

type Member struct {
	GuildID string
	Points  int
	Last    int64
}

type Item struct {
	ItemID string
	Name   string
	Price  int
	Type   string
	Value  string
}

type MinMax struct {
	Min int
	Max int
}

func RegisterUser(c *mgo.Database, id string, username string) error {
	users := c.C("users")
	newUser := User{
		UserID:   id,
		Username: username,
	}
	newUser.Guilds = make(map[string]Member)
	err := users.Insert(&newUser)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func RegisterMember(c *mgo.Database, userid string, guildid string) error {
	users := c.C("users")
	result := User{}
	err := users.Find(bson.M{"userid": userid}).One(&result)
	if err != nil {
		return errors.New(err.Error())
	}

	newMember := Member{
		GuildID: guildid,
		Points:  0,
		Last:    0,
	}

	result.Guilds[guildid] = newMember
	users.Remove(bson.M{"userid": userid})
	err = users.Insert(&result)
	if err != nil {
		return errors.New(err.Error())
	}
	/*users.Update(bson.M{"userid": userid}, bson.M{"$set": bson.M{

	}})*/
	//
	return nil
}

// FindUser (readonly)
func FindUser(c *mgo.Database, id string) (*User, error) {
	users := c.C("users")
	result := User{}
	err := users.Find(bson.M{"userid": id}).One(&result)
	if err != nil {
		return &User{}, errors.New(err.Error())
	}

	return &result, nil
}

// FindGuild in database
func FindGuild(c *mgo.Database, id string) (Guild, error) {
	guilds := c.C("guilds")
	result := Guild{}
	err := guilds.Find(bson.M{"guildid": id}).One(&result)

	if err != nil {
		return Guild{}, errors.New(err.Error())
	}
	return result, nil
}

func SaveGuild(c *mgo.Database, guild Guild, guildID string) error {
	guilds := c.C("guilds")
	guilds.Remove(bson.M{"guildid": guildID})
	err := guilds.Insert(&guild)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func SaveUser(c *mgo.Database, user *User, userID string) error {
	users := c.C("users")
	users.Remove(bson.M{"userid": userID})
	err := users.Insert(&user)
	if err != nil {
		return err
	}
	return nil
}

// AddGuild to database
func AddGuild(c *mgo.Database, id string, name string) error {
	guilds := c.C("guilds")
	err := guilds.Insert(
		&Guild{
			GuildID:   id,
			GuildName: name,
			PWork:     MinMax{100, 500},
			PMessage:  MinMax{1, 10}},
	)

	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func AddPoints(c *mgo.Database, userID string, guildID string) error {
	datg, err := FindGuild(c, guildID)
	if err != nil {
		return err
	}
	datu, err := FindUser(c, userID)
	if err != nil {
		return err
	}

	datm := datu.Guilds[guildID]

	if time.Now().Unix() > datm.Last+60 {
		datm.Last = time.Now().Unix()
		datm.Points += RandomBetween(datg.PMessage.Min, datg.PMessage.Max)

		datu.Guilds[guildID] = datm

		users := c.C("users")
		users.Remove(bson.M{"userid": userID})
		err = users.Insert(&datu)
		if err != nil {
			return err
		}
	}

	return nil
}
