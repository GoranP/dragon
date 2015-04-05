/* EchoServer
 */
package main

import (
	"hub"
	"srv"
)

func main() {

	disp.RunDisp()
	srv.Listen()

	//ctrl.GenerateUsers()

	//var err error
	/*	err := auth.DeletePlayer("goran@kempo.hr")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("User deleted")
	*/
	/*	err = auth.Register("goran", "goran@kempo.hr", "perozderopassword")

		if err != nil {
			fmt.Println("registration failed")
			fmt.Println(err)
			return
		}

		err = auth.ActivatePlayer("goran@kempo.hr")
		if err != nil {
			fmt.Println(err)
			fmt.Println("activation failed")
			return
		}

		err = auth.Login("goran@kempo.hr", "perozderopassword")
		if err != nil {
			fmt.Println("login failed")
			fmt.Println(err)
			return
		} else {
			fmt.Println("login OK")
		}
	*/
	/*	err = auth.DeletePlayer("goran@kempo.hr")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("User deleted")
	*/
	//	session, err := mgo.Dial("5-mongo.supersport.local")
	//	if err != nil {
	//		panic(err)
	//	}
	//	defer session.Close()

	//	fmt.Println(session.DatabaseNames())
	//	db := session.DB("tecajna")

	//	fmt.Println(db.CollectionNames())
	//	c := db.C("dogadjaji")

	//var data map[string]string

	//	var data bson.M

	//	res := c.Find(nil)
	//	//fmt.Println(err)
	//	er := res.One(&data)

	//	if er != nil {
	//		fmt.Println(er)
	//	}

	// http://stackoverflow.com/questions/18340031/unstructured-mongodb-collections-with-mgo
	// http://www.slideshare.net/spf13/painless-data-storage-with-mongodb-go?ref=http://spf13.com/presentation/MongoDB-and-Go/

	// game architecture
	// http://www.mmorpg.com/blogs/FaceOfMankind/052013/25185_A-Journey-Into-MMO-Server-Architecture
	// http://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.134.1944&rep=rep1&type=pdf
	// http://www.eldergame.com/2009/04/smartfoxserver-the-mmo-engine-for-indies/

	//	fmt.Println(data["vrijeme"])
	//	fmt.Println(data["domacin"])
	//	fmt.Println(data["gost"])

	//	c.FindId(66667111).One(&data)

	//	fmt.Println(data["vrijeme"])
	//	fmt.Println(data["domacin"])
	//	fmt.Println(data["gost"])

	//	fmt.Println("SLICES")
	//	var d bson.D
	//	err = c.Find(nil).One(&d)

	//	for _, elem := range d {
	//		fmt.Println(elem.Name, elem.Value)
	//	}

	//	type Dogadjaj struct {
	//		ID      bson.ObjectId `bson:"_id,omitempty"`
	//		Domacin string
	//		Gost    string
	//		Vrijeme time.Time
	//		Extra   bson.M `bson:",inline"`
	//	}

	//	var dog Dogadjaj

	//	err = c.FindId(66667111).One(&dog)

	//	fmt.Println(dog.Domacin)
	//	fmt.Println(dog.Gost)
	//	fmt.Println(dog.Vrijeme)
	//	fmt.Println(dog.Extra["naziv"])

}
