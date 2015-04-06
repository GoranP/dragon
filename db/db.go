// db project db.go
package db

import (
	"fmt"
	"player"
	"time"
	"utl"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	MongoDBHosts = "127.0.0.1"
	AuthDatabase = "players"
	AuthUserName = ""
	AuthPassword = ""
	TestDatabase = "players"
)

// database session for all go routines
var (
	dbSession *mgo.Session
)

//set global session variable
func connectToMongo() error {

	if dbSession != nil {
		return nil
	}

	session, err := mgo.Dial(MongoDBHosts)
	defer func() {
		session.Close()
		utl.TimeTrack(time.Now(), "connectToMongo")
	}()
	if err != nil {
		return err
	}

	dbSession = session.Copy()
	return nil

}

func GetDBSession() (*mgo.Session, error) {

	err := connectToMongo()
	if err != nil {
		return nil, err
	}

	return dbSession.Copy(), nil

}

// specification notes
// login with email
// need index in collection over activationid and email/deletedat (_id is intrinsic)

// mongodb OPS http://info.mongodb.com/rs/mongodb/images/10gen-MongoDB_Operations_Best_Practices.pdf
// https://github.com/apexskier/httpauth/blob/master/mongoBackend.go

func GetPlayerByEmail(email string) (p player.Player, e error) {

	session, err := GetDBSession()
	if err != nil {
		return p, err
	}

	defer func() {
		session.Close()
		utl.TimeTrack(time.Now(), "GetPlayerByEmail")
	}()

	//session.SetMode(mgo.Monotonic, true)
	players := session.DB(AuthDatabase).C("players")

	query := bson.M{"email": email, "del_at": nil}
	err = players.Find(query).One(&p)

	if err != nil {
		//fmt.Println(err)
		return p, err
	}

	return p, nil

}

func RegisterPlayer(player player.Player) error {

	session, err := GetDBSession()
	if err != nil {
		return err
	}

	defer func() {
		session.Close()
		utl.TimeTrack(time.Now(), "RegisterPlayer")
	}()

	//session.SetMode(mgo.Monotonic, true)

	players := session.DB(AuthDatabase).C("players")
	err = players.Insert(&player)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func UpdatePlayer(player player.Player) error {

	session, err := GetDBSession()
	if err != nil {
		return err
	}

	defer func() {
		utl.TimeTrack(time.Now(), "GetPlayer")
		session.Close()
	}()

	player.Upd_at = time.Now()
	players := session.DB(AuthDatabase).C("players")

	players.UpdateId(player.Id, &player)

	return nil
}
