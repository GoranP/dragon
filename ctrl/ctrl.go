package ctrl

import (
	"conn"
	"log"

	"github.com/gorilla/websocket"
)

// Registered connections.
var connections = make(map[*conn.Connection]string) //map[*Connection] string

//hub manager - currently only one hub with many connection - broadcast
//var globalHub *hub

// chat-hub-connections logic
//////////////////////////////
func Connection(ws *websocket.Conn) {
	log.Println("new websocket conn")
	//stats ws++

	//add to map
	c := conn.NewRaw(ws)
	connections[c] = c.Id
	go conn.Register(c)

	for k, v := range connections {
		log.Println(v)
		k.Send(v)
	}

	//
	/*h := conn.NewHub()
	go h.Run()

	//new app connection
	go conn.New(ws, h)
	*/
}
