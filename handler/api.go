package handler

import (
	"encoding/json"
	"log"
)

var (
	err error
	m   *message
	b   bool
)

const (
	// Start: For Users
	joinTable   = "joinTable"
	setPosition = "setPosition"
	// Start: For Raspberry
	newGoal
)

func handleCommunication(c *Client, message []byte) {
	if m, b = unmarshal(message); !b {
		return
	}
	log.Println(string(message))
	sendPrivate(c, "hallo")
	if *c.user.ID == "" {
		if !initUser(c, *m) {
			return
		}
	}
	switch m.Command {
	case "joinTable":
		log.Println("done!")
		sendPrivate(c, "done")
	}
}

// initUser sets the ID and uname. If raspberry, just set that as name.
func initUser(c *Client, m message) bool {
	if m.ID == "" || m.Username == "" {
		return false
	}
	c.user.ID = &m.ID
	c.user.Username = &m.Username
	return true
}

// unmarshal converts the byte into a message struct.
// If it fails it returns an empty struct and false.
func unmarshal(msg []byte) (*message, bool) {
	var m = &message{}
	err = json.Unmarshal(msg, m)
	if err != nil {
		// Not logging error because its most probably a ping/pong message.
		//log.Println(err)
		return m, false
	}
	return m, true
}

type message struct {
	// The user key
	ID string `json:"id"`

	// The username
	Username string `json:"username"`

	// The table id
	Table bool `json:"table"`

	// Command to execute
	Command string `json:"command"`

	// Insert all possible command structures
	Values map[string]string `json:"values"`
}
