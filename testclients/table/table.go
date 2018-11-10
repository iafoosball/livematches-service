package table

import (
	"bufio"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"time"
)

var (
	serverMsg = ""
)

func Start(tableID string, scenario string, addr string, end chan bool) {
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("start ws client")
	//var addr = flag.String("addr", "iafoosball.aau.dk:9003", "http service address")
	if addr == "" {
		addr = "0.0.0.0:9003"
	}
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: addr, Path: "/tables/" + tableID}
	log.Printf("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	handleErr(err, "making websocket connection")
	//defer c.Close()

	client := &client{
		send: make(chan []byte, 256),
	}
	//handle input here. Is send to the server
	go func() {
		defer c.Close()
		path := "../testclients/" + scenario + "/" + tableID
		file, err := os.Open(path)
		defer file.Close()
		// Read first line, either description or quit
		scanner := bufio.NewScanner(file)
		for {
			for scanner.Scan() {
				line := scanner.Text()
				if line == "quit" {
					end <- true
					return
				}
				// Read and send command
				scanner.Scan()
				line = scanner.Text()
				log.Println(line)
				client.send <- []byte(line)
				// Read line and check if return from server contains line
				scanner.Scan()
				line = scanner.Text()
				// Check if response is as expected
				line = scanner.Text()
				time.Sleep(1 * time.Second)
				checkResponse(line)
			}
			handleErr(err, "Read commands from file")
		}
	}()

	go func() {
		defer c.Close()
		for {
			_, message, err := c.ReadMessage()
			serverMsg = string(message)
			log.Println(serverMsg)
			if err != nil {
				log.Println(tableID+" ReadPump :", err)
				end <- true
				return
			}
		}
	}()

	for {
		select {
		case message, ok := <-client.send:
			if !ok {
				c.WriteMessage(websocket.CloseMessage, []byte{})
			}
			c.WriteMessage(websocket.TextMessage, message)
		case _ = <-end:
			defer c.Close()
			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
			}
			return
		}
	}
}

func checkResponse(msg string) {
	if !strings.Contains(serverMsg, msg) {
		log.Println("Server msg: " + serverMsg)
		log.Println("Client msg: " + msg)
		log.Fatalln("Expected string was not found in message from server")
	}
}

type client struct {
	// The websocket connection.
	conn *websocket.Conn
	// Buffered channel of outbound messages.
	send chan []byte
}

func handleErr(err error, customMsg string) {
	if err != nil {
		log.Println(customMsg)
		log.Fatalln(err)
	}
}
