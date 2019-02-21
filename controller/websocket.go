package controller

import (
	"github.com/gorilla/websocket"
	"net/http"
	"traefik/log"
)

var upgrader = websocket.Upgrader{} // use default options
func Home(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w,"Hello:%s",r.Host)
	http.ServeFile(w, r, "./webchart/html/home.html")
}

func Echo(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal("create websocket shack failed,", err)
		return
	}
	defer conn.Close()

	for {
		msgTyp, msg, err := conn.ReadMessage()

		if err != nil {
			log.Fatal("read message failed:", err)
			continue
		}

		err = conn.WriteMessage(msgTyp, msg)

		if err != nil {
			log.Fatal("write message failed:", err)
			continue
		}

	}

}
