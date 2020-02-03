package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sync"
)

var chat chatroom

type chatroom struct {
	sync.Mutex
	clients []chan string
}

func (c *chatroom) Register() (<-chan string, int) {
	log.Infof("Registering #%d client...", len(c.clients)+1)
	ch := make(chan string)

	c.Lock()
	defer c.Unlock()
	closer := len(c.clients)
	c.clients = append(c.clients, ch)

	return ch, closer
}

func (c *chatroom) Broadcast(msg string) {
	log.Infof("Broadcast msg to all clients: %s", msg)
	for _, ch := range c.clients {
		ch <- msg
	}
}

func (c *chatroom) Unregister(i int) {
	log.Infof("Unregistering #%d client...", i)
	c.Lock()
	defer c.Unlock()

	close(c.clients[i])

	copy(c.clients[i:], c.clients[i+1:])
	c.clients[len(c.clients)-1] = nil
	c.clients = c.clients[:len(c.clients)-1]
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func echo(w http.ResponseWriter, r *http.Request) {
	log.Info(r)

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		http.Error(w, "unable to upgrade connection", http.StatusInternalServerError)
	}
	defer c.Close()

	ch, closer := chat.Register()
	defer chat.Unregister(closer)

	go func() {
		for m := range ch {
			err := c.WriteMessage(1, []byte(m))
			if err != nil {
				log.Error(err)
			}
		}
	}()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Errorln("read:", err)
			break
		}
		log.Infof("MessageType: %d Message: %s", mt, string(message))

		chat.Broadcast(string(message))
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", echo)
	log.Fatal(http.ListenAndServe("localhost:8088", router))
}
