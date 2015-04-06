package db

import "testing"

//test password hashing functionality
func TestHash(t *testing.T) {

	var p Player
	err := p.SetPassword("123456")
	if err != nil {
		t.Error(err)
	}

	err = p.Checkhash("123456")
	if err != nil {
		t.Error(err)
	}

}

func TestSaveGetPlayer(t *testing.T) {

	var p Player

	p.Email = "pero@zdero.com"
	p.Username = "pero perić"
	p.SetPassword("123456")

	RegisterPlayer(p)

	pn, err := GetPlayer("pero perić")
	if err != nil {
		t.Error(err)
	}

	if pn.Username != "pero perić" {
		t.Errorf("Wrong username %s. Should be %s", pn.Username, "pero perić")

	}

	if pn.Email != "pero@zdero.com" {
		t.Errorf("Wrong email %s. Should be %s", pn.Username, "pero perić")
	}

	if pn.CheckPassword("123456") != nil {
		t.Errorf("Wrong password %s. Should be %s", "123456", "123456")

	}

}
