package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Joke struct {
	ID   uint32 `json:"id"`
	Joke string `json:"joke"`
}

type JokeResponse struct {
	Type  string `json:"type"`
	Value Joke   `json:"value"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// "Шина" событий, регистрация клиентов и рассылка сообщений идет отсюда
type Bus struct {
	register  chan *websocket.Conn
	broadcast chan []byte
	clients   map[*websocket.Conn]bool
}

func (b *Bus) Run() {
	for {
		select {
		case message := <-b.broadcast:
			log.Println("sending...")
			// каждому зарегистрированному клиенту шлем сообщение
			for client := range b.clients {
				w, err := client.NextWriter(websocket.TextMessage)
				if err != nil {
					// если достучаться до клиента не удалось, то удаляем его
					log.Println("client connection lost")
					delete(b.clients, client)
					log.Println("clients: ", len(b.clients))
					continue
				}

				w.Write(message)
				w.Close()
			}
		case client := <-b.register:
			// регистрируем клиентов в мапе клиентов
			log.Println("User registered")
			b.clients[client] = true
			log.Println("clients: ", len(b.clients))
		}
	}
}

func NewBus() *Bus {
	return &Bus{
		register:  make(chan *websocket.Conn, 1024),
		broadcast: make(chan []byte, 8),
		clients:   make(map[*websocket.Conn]bool),
	}
}

func runJoker(b *Bus) {
	for {
		// каждые 5 секунд ходим за шутками
		<-time.After(5 * time.Second)
		s := getJoke()
		log.Printf("Its joke time:%s\n", string(s))
		b.broadcast <- s
	}
}

func getJoke() []byte {
	c := http.Client{}
	resp, err := c.Get("http://api.icndb.com/jokes/random?limitTo=[nerdy]")
	if err != nil {
		return []byte("jokes API not responding")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	joke := JokeResponse{}

	err = json.Unmarshal(body, &joke)
	if err != nil {
		return []byte("Joke error")
	}

	return []byte(fmt.Sprintf("Joke #[%d]:%s", joke.Value.ID, joke.Value.Joke))
}

func main() {
	bus := NewBus()
	go bus.Run()
	go runJoker(bus)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// апгрейд соединения
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err)
		}
		bus.register <- ws
	})
	_ = http.ListenAndServe(":8081", nil)
}
