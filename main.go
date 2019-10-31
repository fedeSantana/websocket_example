package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)
//192.168.0.4:9036
var addr = flag.String("addr", "192.168.0.6:3001", "http service address")
var upgrader = websocket.Upgrader{}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, "")
}

func echo(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade: ", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read: ", err)
			break
		}
		log.Printf("recv: %s", message)
		message = []byte(fmt.Sprintf("respondiendo msg %s", message))
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write: ", err)
			break
		}
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	router := mux.NewRouter()
	router.HandleFunc("/", rootHandler).Methods("GET")
	router.HandleFunc("/ws", echo)

	log.Fatal(http.ListenAndServe(*addr, router))
}
