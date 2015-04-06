package player

import (
	"fmt"
	"time"
	"utl"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type Player struct {
	Id     bson.ObjectId `bson:"_id,omitempty"`
	Nick   string
	Email  string
	Pass   string
	Act_id string

	Cre_at time.Time
	Upd_at time.Time
	Del_at *time.Time `bson:"del_at,omitempty"`
	Act_at *time.Time `bson:"act_at,omitempty"`

	//	Extra bson.M `bson:",inline"`
}

// format of bcrypt hash
// $2a$[Cost]$[Base64Salt][Base64Hash]
// example
// $2a$10$TwentytwocharactersaltThirtyonecharacterspasswordhash

func (p *Player) SetPassword(password string) error {

	defer utl.TimeTrack(time.Now(), "SetPassword")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	p.Pass = string(hashedPassword)
	return nil

}

func (p *Player) CheckPassword(password string) error {

	defer utl.TimeTrack(time.Now(), "CheckPassword")

	err := bcrypt.CompareHashAndPassword([]byte(p.Pass), []byte(password))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
