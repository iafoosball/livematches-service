package impl

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
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
}

func sendAll(c *Client, msg string) {
	c.Hub.broadcast <- []byte(msg)
}

func sendMatch(c *Client, msg string) {
	c.LiveMatch.MatchCast <- []byte(msg)
}

func sendMatchData(c *Client) {
	b, err := json.Marshal(*c.LiveMatch.M)
	handleErr(err)
	c.LiveMatch.MatchCast <- b
}

func sendPrivate(c *Client, msg string) {
	c.Send <- []byte(msg)
}

// Opens: If Table closes, kick all clients, if full kick
func closeTable(c *Client) {
	log.Println("called close Table")
	for u, _ := range c.LiveMatch.Clients {
		if u.IsUser {
			log.Println(u.User.ID)
			closeUser(u)
		}
	}
	c.LiveMatch = nil
	c.Conn.Close()
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		//closeUser(c)
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		handleCommunication(c, message)
	}
	select {
	case _ = <-c.End:
		return
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is Started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case ok := <-c.End:
			log.Println("close write pump")
			if ok {
				if c.IsUser {
					log.Println(c.User.ID)
				}
				c.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				time.Sleep(1 * time.Second)
				c.Conn.Close()
				return
			}
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket broadcast.
			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request, isUser bool, tableID string, userID string) {
	log.Println("new Client connected with  tableID: " + tableID + " userID: " + userID)
	defer r.Body.Close()
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	conn.SetWriteDeadline(time.Now().Add(30 * time.Second))
	handleErr(err)
	if isUser {
		if !tableExists(tableID, hub) {
			conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		} else {
			newU := true
			for c := range hub.clients {
				log.Println(c.ID)
				if c.ID == userID {
					log.Println("match of ids")
					leavematch(c)
					c.End <- true
					close(c.Send)

					c.Conn = conn
					c.Send = make(chan []byte, 256)
					go c.writePump()
					go c.readPump()
					joinMatch(c, userID)
					log.Println("old User")
					newU = false
				}
			}
			if newU {
				client := newClient(userID, hub, conn, true)
				client.User = &models.MatchUsersItems0{ID: userID}
				joinMatch(client, tableID)
				sendMatchData(client)
			}
		}
	} else {
		client := newClient(userID, hub, conn, false)
		client.Table = &models.Table{ID: tableID}
		createMatch(client, tableID)
	}
}

func newClient(id string, hub *Hub, conn *websocket.Conn, isUser bool) *Client {
	var client *Client
	client = &Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256), IsUser: false}
	client.ID = id
	client.Hub.register <- client
	// Allow collection of memory referenced by the caller by doing all work in
	go client.writePump()
	go client.readPump()
	return client
}

func tableExists(tableID string, hub *Hub) bool {
	for c := range hub.clients {
		log.Println(c.ID)
		if !c.IsUser && c.ID == tableID {
			log.Println("table exists")
			return true
		}
	}
	log.Println("table no table tlikea")
	return false
}

// MatchUsersItems0 is a middleman between the websocket connection and the LiveMatch.
type Client struct {
	Hub *Hub

	// The websocket connection.
	Conn *websocket.Conn

	// Buffered channel of outbound messages.
	Send chan []byte

	// dicsonnects the client gracefully
	End chan bool

	// The LiveMatch which the User joins
	LiveMatch *LiveMatch

	// Is the client id -- > username
	ID string

	// IsUser is true for a User. False for a Table.
	IsUser bool

	//The client data. Going to be nil for a Table (or empty pointer?)
	User *models.MatchUsersItems0

	// The Table data. Nil for a User.
	Table *models.Table
}

func handleErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
