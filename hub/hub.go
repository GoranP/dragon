package disp

import (
	"encoding/json"
	"log"
	"sync"

	"code.google.com/p/go-uuid/uuid"

	"github.com/gorilla/websocket"
)

var mx = &sync.Mutex{}

// fan in send two chan to fanIn function and return fan channel.
// Internally launch 2 go routines taht sends messages to fanIn channel
/// in messages json diferrentiant themn with sessionid (uuid) and process

type invite struct {
	Invite string
}

type mesg struct {
	ID  string `json:"id"`
	Msg struct {
		Invite string `json:"invite"`
	} `json:"msg"`
}

// or just launch two go routines that
type connection struct {
	ws *websocket.Conn
	rc chan []byte
	wc chan []byte
	id string
}

func (c *connection) reader() {
	go func() {

		for {
			_, message, err := c.ws.ReadMessage()
			if err != nil {
				break
			}
			c.rc <- message
		}
		c.ws.Close()
	}()

}

func (c *connection) writer() {

	go func() {

		for message := range c.wc {
			log.Println("writer before send")
			log.Println(string(message))
			err := c.ws.WriteMessage(websocket.TextMessage, message)
			log.Println("writer after send")
			if err != nil {
				break
			}
		}
		c.ws.Close()
	}()
}

type Dispatcher struct {
	conns      map[string]*connection
	controlin  chan []byte
	controlout chan []byte
	register   chan []byte
	unregister chan []byte
}

func (d *Dispatcher) AddCon(c *connection) {

	/*      WARNING: DATA RACE
	        Read by goroutine 12:
	        runtime.growslice()
	        /usr/lib/go/src/pkg/runtime/slice.c:62 +0x0
	        hub.(*Dispatcher).AddCon()
	        /home/goran/projects/go/src/hub/hub.go:61 +0x7f
	*/

	mx.Lock()
	d.conns[c.id] = c
	mx.Unlock()

	log.Printf("Size of connections %v", len(d.conns))

	d.fanin(c)
	d.fanout(c)

	c.reader()
	c.writer()

	//send id to client
	c.rc <- []byte("{\"id\":\"" + c.id + "\"}")
}

//messages from dispatcher to all clients
//special case when only two clients/players
func (d *Dispatcher) fanout(c *connection) {
	go func() {
		for {
			c.wc <- <-d.controlout
		}
	}()
}

//messages from all clients to dispatcher
func (d *Dispatcher) fanin(c *connection) {
	go func() {
		for {
			d.controlin <- <-c.rc
		}
	}()
}

func (d *Dispatcher) run() {
	log.Println("run")
	defer log.Println("out ")

	for {

		select {
		case msg := <-d.controlin:
			log.Printf("got message on control %v", string(msg))
			log.Println("Sending message to master game logic")

			s := mesg{}

			err := json.Unmarshal(msg, &s)

			log.Printf("err msg: %v", err)
			log.Printf("id: %v invite: %v", s.ID, s.Msg.Invite)

			if s.Msg.Invite != "" {
				mx.Lock()
				c := d.conns[s.Msg.Invite]
				mx.Unlock()
				if c == nil {
					log.Println("no connection wiht id " + s.Msg.Invite)
					return
				}
				c.wc <- []byte("(" + s.ID + ")")
				c.wc <- []byte("Would you like to play?")
			}

			if s.Msg.Invite == "" {
				//fanout them all
				for i := 0; i < len(d.conns); i++ {
					d.controlout <- msg
				}

			}

			if string(msg) == "close" {
				panic("aaa....end!")
			}

		}

	}
}

var gd = Dispatcher{controlin: make(chan []byte), controlout: make(chan []byte), conns: make(map[string]*connection)}

func RunDisp() {
	go gd.run()
}

func WebConn(ws *websocket.Conn) {
	log.Println("new ws")
	c := connection{ws: ws, rc: make(chan []byte), wc: make(chan []byte), id: uuid.New()}
	//s := session{ID: uuid.New(), Playerid: uuid.New(), expire: time.Now()}

	gd.AddCon(&c)

	//d := Dispatcher{channel1: c.GetChannel(), conn: c}

	//ws.WriteMessage(websocket.TextMessage, m)

	//go d.run()

	/*	log.Println("before chan write")
			log.Printf("channel %v", *d.channel1)
		d.channel1 <- []byte(s.ID)
	*/

}

/*type MsgDisp struct {

	// Inbound/putbound messages from the connections.
	broadcast chan []byte

	// Register requests from the connections.
	register chan *connection

	// Unregister requests from connections.
	unregister chan *connection
}

func newHub() *hub {
	return &hub{
		broadcast:   make(chan []byte),
		register:    make(chan *connection),
		unregister:  make(chan *connection),
		connections: make(map[*connection]bool),
	}
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			h.connections[c] = true
		case c := <-h.unregister:
			if _, ok := h.connections[c]; ok {
				delete(h.connections, c)
				close(c.send)
			}
		case m := <-h.broadcast:
			for c := range h.connections {
				select {
				case c.send <- m:
				default:
					delete(h.connections, c)
					close(c.send)
				}
			}
		}
	}
}
*/
