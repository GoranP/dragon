package conn

import (
	"log"

	"code.google.com/p/go-uuid/uuid"

	"github.com/gorilla/websocket"
)

type hub struct {
	// Registered connections.
	connections map[*Connection]bool

	// Inbound messages from the connections.
	broadcast chan []byte

	// Register requests from the connections.
	register chan *Connection

	// Unregister requests from connections.
	unregister chan *Connection
}

func NewHub() *hub {
	return &hub{
		broadcast:   make(chan []byte),
		register:    make(chan *Connection),
		unregister:  make(chan *Connection),
		connections: make(map[*Connection]bool),
	}
}

func (h *hub) Run() {
	log.Println("runing new hub..")
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

type Connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
	Id   string
	// The hub.
	h *hub
}

func (c *Connection) Send(msg string) {
	log.Printf("sending msg %v", msg)
	c.send <- []byte(msg)
}

func (c *Connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
		c.h.broadcast <- message
	}
	c.ws.Close()
}

func (c *Connection) writer() {
	log.Println("writer before send")
	for message := range c.send {
		log.Println("before write")
		err := c.ws.WriteMessage(websocket.TextMessage, message)
		log.Println("after write")
		if err != nil {
			break
		}
	}
	c.ws.Close()
}

func NewRaw(ws *websocket.Conn) *Connection {
	return &Connection{send: make(chan []byte, 256), ws: ws, Id: uuid.New()}
}

func Register(c *Connection) {

	c.h.register <- c
	defer func() { c.h.unregister <- c }()
	go c.writer()
	c.reader()

}

func New(ws *websocket.Conn, h *hub) {

	c := &Connection{send: make(chan []byte, 256), ws: ws, h: h}
	c.h.register <- c
	defer func() { c.h.unregister <- c }()
	go c.writer()
	c.reader()
}
