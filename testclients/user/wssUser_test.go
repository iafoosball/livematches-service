package user

import (
	"crypto/tls"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestStartWssTable(t *testing.T) {
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("start ws client")
	//var addr = flag.String("addr", "iafoosball.aau.dk:9003", "http service address")
	//addr := "iafoosball.me:9003"
	addr := "localhost:8013"
	tableID := "table-1"
	userID := "user"

	for i := 0; i < 1000; i++ {
		if i%50 == 0 {
			time.Sleep(3 * time.Second)
		}
		u := "ws://" + addr + "/users/?tableID=" + tableID + "&userID=" + userID + strconv.Itoa(i)
		//u := "wss://" + addr + "/tables/?tableID=" + tableID
		log.Println(u)
		log.Printf("connecting to %s", u)
		d := websocket.Dialer{}
		d.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		_, _, err := d.Dial(u, nil)
		handleErr(err, "making websocket connection")
	}
	//defer c.Close()

	//_ := &client{
	//	send: make(chan []byte, 256),
	//}

}
