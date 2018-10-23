package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
)

var (
	serverMsg      = ""
	stop      bool = false
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("start ws client")
	var addr = flag.String("addr", "localhost:9003", "http service address")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws/tables/table-2"}
	log.Printf("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})
	client := &client{
		send: make(chan []byte, 256),
	}
	//handle input here. Is send to the server
	go func() {
		defer close(done)
		for {
			if !stop {
				switch serverMsg {
				case "":
					msg := "{ \"command\": \"createMatch\", \"values\": { \"match\": \"123\", \"side\": \"blue\", \"attack\": \"true\" } }"
					client.send <- []byte(msg)
					serverMsg = "closeMatch"
				case "closeMatch":
					msg := "{ \"command\": \"closeMatch\", \"values\": { } }"
					client.send <- []byte(msg)
					stop = true

					//case "":
					//case "":
					//case "":
				}
				time.Sleep(30 * time.Second)
			}
		}
	}()

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			// Prints the server message
			serverMsg = string(message)
			log.Println(serverMsg)
			if err != nil {
				log.Println("read:", err)
				return
			}
			c.WriteMessage(websocket.TextMessage, message)

		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case message, ok := <-client.send:
			if !ok {
				c.WriteMessage(websocket.CloseMessage, []byte{})
			}
			log.Println(string(message))
			c.WriteMessage(websocket.TextMessage, message)
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}

}

type client struct {
	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}