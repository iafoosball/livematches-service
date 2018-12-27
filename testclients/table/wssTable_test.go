package table

import (
	"crypto/tls"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"testing"
)

func TestStartWssTable(t *testing.T) {
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("start ws client")
	//var addr = flag.String("addr", "iafoosball.aau.dk:9003", "http service address")
	addr := "iafoosball.me:9003"
	tableID := "table-1"

	u := url.URL{Scheme: "wss", Host: addr, Path: "/tables/" + tableID}
	log.Println(u.Path)
	log.Printf("connecting to %s", u.String())
	d := websocket.Dialer{}
	d.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	_, _, err := d.Dial(u.String(), nil)
	handleErr(err, "making websocket connection")
	//defer c.Close()

	//_ := &client{
	//	send: make(chan []byte, 256),
	//}

}
