package table

import (
	"github.com/gorilla/websocket"
	"log"
	"testing"
)

func TestStartWssTable(t *testing.T) {
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("start ws client")
	//var addr = flag.String("addr", "iafoosball.aau.dk:9003", "http service address")
	//addr := "iafoosball.me:9003"
	addr := "localhost:8003"
	tableID := "table-1"

	u := "ws://" + addr + "/tables/?tableID=" + tableID
	//u := "wss://" + addr + "/tables/?tableID=" + tableID
	log.Println(u)
	log.Printf("connecting to %s", u)
	d := websocket.Dialer{}
	//d.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	_, _, err := d.Dial(u, nil)
	handleErr(err, "making websocket connection")
	//defer c.Close()

	//_ := &client{
	//	send: make(chan []byte, 256),
	//}

}
