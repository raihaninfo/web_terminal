package main

import (
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/ws", Ws)
	http.HandleFunc("/", Home)

	// static file server
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/index.gohtml")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)

}

func Ws(w http.ResponseWriter, r *http.Request) {
	// allow origin all
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// fmt.Println(messageType)
		command, err := ExecuteCommand(string(p))
		if err != nil {
			log.Println(err)
			return
		}

		if err := conn.WriteMessage(messageType, []byte(command)); err != nil {
			log.Println(err)
			return
		}
	}

}

func ExecuteCommand(cmd string) (string, error) {
	s := strings.Split(cmd, " ")
	c, _ := exec.Command(s[0], s[1:]...).Output()
	// if err != nil {
	// 	return "", err
	// }

	cString := string(c)
	cString = strings.ReplaceAll(cString, "\n", "<br/>")

	return cString, nil

}
