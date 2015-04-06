package srv

import (
	"flag"
	"hub"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/gorilla/websocket"
)

var (
	addr      = flag.String("addr", ":8080", "http service address")
	assets    = flag.String("assets", defaultAssetPath(), "path to assets")
	homeTempl *template.Template
)

func Listen() {

	flag.Parse()
	homeTempl = template.Must(template.ParseFiles(filepath.Join(*assets, "home.html")))

	//this goes to controller
	/*	h := newHub()
		go h.run()
	*/http.HandleFunc("/", homeHandler)
	http.Handle("/ws", wsHandler{})
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func defaultAssetPath() string {
	/*	p, err := build.Default.Import("gary.burd.info/go-websocket-chat", "", build.FindOnly)
		if err != nil {
			return "."
		}
	*/
	return "."
}

func homeHandler(c http.ResponseWriter, req *http.Request) {
	homeTempl.Execute(c, req.Host)
}

var upgrader = &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

type wsHandler struct {
}

func (wsh wsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	//ctrl.Connection(ws)
	go disp.WebConn(ws)

}
