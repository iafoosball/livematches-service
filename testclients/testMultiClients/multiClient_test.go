package testMultiClients

import (
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"testing"
	"time"
)

var (
	i int
)

func TestMultiClientHttp(t *testing.T) {
	for i := 1; i < 1000000; i++ {
		log.Println(i)
		time.Sleep(1 * time.Millisecond)
		go httpRequ()
	}
}

func httpRequ() {
	r, _ := http.Get("http://iafoosball.me:9000/matches")
	_, _ = ioutil.ReadAll(r.Body)
}

func TestMultiClient(t *testing.T) {
	c := 0
	for i := 1; i < 100; i++ {
		for c = 0; c < 100; c++ {
			go connect()
			time.Sleep(50 * time.Millisecond)
		}
		//time.Sleep(500 * time.Millisecond)
		log.Println(time.Now())
		log.Println("Number of running clients: " + strconv.Itoa(i*100))
	}
	t.Error(i*100 + c)

}

func connect() {
	i++
	userID := strconv.Itoa(i) + "-_-t"
	tableID := "table-1"
	addr := "iafoosball.me:9003"

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: addr, Path: "/users/" + tableID + "/" + userID}
	c, _, _ := websocket.DefaultDialer.Dial(u.String(), nil)

	go func() {
		defer c.Close()
		for {
			_, _, err := c.ReadMessage()
			if err != nil {
				//log.Println(userID+" ReadPump :", err)
				return
			}
		}
	}()
}
