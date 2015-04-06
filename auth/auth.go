// auth project auth.go
package auth

import (
	"db"
	"fmt"
	"log"
	"player"
	"time"
	"utl"

	"code.google.com/p/go-uuid/uuid"
)

func Register(nick string, email string, password string) error {

	defer utl.TimeTrack(time.Now(), "Register")

	// already exists?
	p, err := db.GetPlayerByEmail(email)

	if err != nil && err.Error() != "not found" {
		return err
	}

	if p.Email == email {
		//		log.Printf("User %v already exists", nick)
		return fmt.Errorf("User %v already exists", nick)
	}

	newplayer := player.Player{
		Email:  email,
		Nick:   nick,
		Act_id: uuid.New(),
		Cre_at: time.Now(),
		Upd_at: time.Now(),
	}
	newplayer.SetPassword(password)

	err = db.RegisterPlayer(newplayer)
	return err

}

func Login(email string, password string) error {

	defer utl.TimeTrack(time.Now(), "Login")

	player, err := db.GetPlayerByEmail(email)
	if err != nil {
		return err
	}

	if player.Act_at == nil {
		return fmt.Errorf("not activated")
	}

	err = player.CheckPassword(password)
	if err != nil {
		return err
	}

	// IsActivated?

	log.Printf("player activated at %v", player.Act_at)
	//if player.Activatedat > time.NewTimer(d)
	// TODO
	// connect conncetion and session in local array

	return nil

}

func DeletePlayer(email string) error {

	defer utl.TimeTrack(time.Now(), "DeletePlayer")

	p, err := db.GetPlayerByEmail(email)
	if err != nil {
		return err
	}

	log.Println(p.Id)
	t := time.Now()
	p.Del_at = &t

	err = db.UpdatePlayer(p)
	if err != nil {
		return err
	}
	return nil

}

func ActivatePlayer(email string) error {

	defer utl.TimeTrack(time.Now(), "ActivatePlayer")

	p, err := db.GetPlayerByEmail(email)
	if err != nil {
		return err
	}

	if p.Act_at != nil {
		log.Printf("%v already activated", email)
		return nil
	}

	log.Println(p.Id)
	t := time.Now()
	p.Act_at = &t

	err = db.UpdatePlayer(p)
	if err != nil {
		return err
	}

	return nil

}
