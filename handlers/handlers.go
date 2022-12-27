package handlers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/websocket"
	"github.com/raihaninfo/web_terminal/helpers"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/index.gohtml")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)

}

func Ws(w http.ResponseWriter, r *http.Request) {
	// allow only localhost:8080 to connect
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return r.Host == "localhost:8080"
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	go func() {
		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			command, err := helpers.ExecuteCommand(string(p))
			if err != nil {
				log.Println(err)
				if err := conn.WriteMessage(messageType, []byte(err.Error()+"<br/>")); err != nil {
					log.Println(err)
					return
				}
			} else {
				if err := conn.WriteMessage(messageType, []byte(command)); err != nil {
					log.Println(err)
					return
				}
			}
		}
	}()

}
