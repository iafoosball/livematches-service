package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/iafoosball/livematches-service/models"
	"log"
	"net/http"
	"time"
)

const (
	// Time allowed to write a broadcast to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong broadcast from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum broadcast size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		handleCommunication(c, message)
	}
}

func sendAll(c *Client, msg string) {
	c.hub.broadcast <- []byte(msg)
}

func sendMatch(c *Client, msg string) {
	c.liveMatch.MatchCast <- []byte(msg)
}

func sendMatchData(c *Client) {
	b, err := json.Marshal(*c.liveMatch)
	handleErr(err)
	c.liveMatch.MatchCast <- b
}

func sendPrivate(c *Client, msg string) {
	c.send <- []byte(msg)
}

func closeUser(c *Client) {
	log.Println("cloooose")
	for i, u := range c.liveMatch.Users {
		if u.ID == c.user.ID {
			c.liveMatch.Users[i] = c.liveMatch.Users[len(c.liveMatch.Users)-1]
			c.liveMatch.Users = c.liveMatch.Users[:len(c.liveMatch.Users)-1]
		}
	}
}

func closeTable(c *Client) {
	c.liveMatch = nil
	//for i, u := range LiveMatches {
	//	if  {
	//		c.liveMatch.Users[i] = c.liveMatch.Users[len(c.liveMatch.Users)-1]
	//		c.liveMatch.Users = c.liveMatch.Users[:len(c.liveMatch.Users)-1]
	//	}
	//}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is Started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		if c.isUser {
			closeUser(c)
		} else {
			closeTable(c)
		}
		//log.Printf("%+v\n", *c.user)
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket broadcast.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request, isUser bool, tableID string, userID string) {
	log.Println("new Client connected with  tableID: " + tableID + " userID: " + userID)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	var client *Client
	client = &Client{hub: hub, conn: conn, send: make(chan []byte, 256), isUser: false}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()

	if isUser {
		client.isUser = true
		client.user = &models.User{ID: userID}
		joinMatch(client, tableID)
		sendMatchData(client)
	} else {
		client.table = &models.Table{ID: tableID}
		createMatch(client, tableID)
	}
	sendMatchData(client)

}

// User is a middleman between the websocket connection and the LiveMatch.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	// The LiveMatch which the user joins
	liveMatch *LiveMatch

	// isUser is true for a user. False for a table.
	isUser bool

	//The client data. Going to be nil for a table (or empty pointer?)
	user *models.User

	// The table data. Nil for a user.
	table *models.Table
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
