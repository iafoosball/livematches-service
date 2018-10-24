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
	joinLobby   = "joinLobby"
	leaveLobby  = "leaveLobby"
	joinTable   = "joinTable"
	leaveTable  = "leaveTable"
	setPosition = "setPosition"
	// Start: For Admin

	newGoal
)

func handleCommunication(c *Client, message []byte) {
	if m, b = unmarshal(message); !b {
		return
	}
	log.Println(m)
	if c.isUser && c.user.Admin {
		handleAdmin(c, m)
	} else if c.isUser {
		handleUsers(c, m)
	} else {
		handleTable(c, m)
	}

}

func handleUsers(c *Client, m *message) {
	switch m.Command {
	case "setPosition":
		sendMatchData(c)
	case "leaveMatch":
		// Tested for normal user.
		leaveMatch(c)
	}
}

func handleAdmin(c *Client, m *message) {
	handleUsers(c, m)
	switch m.Command {
	case "startMatch":
	case "double":
	case "rated":
	case "max_time":
	case "max_goals":
	case "switch_positions":
	}
}

func handleTable(c *Client, m *message) {
	log.Println(c.isUser)
	log.Println("Hanlde table " + c.table.ID + " with cmd: " + m.Command)
	switch m.Command {
	case "closeMatch":
		closeMatch(c)
	case "addGoal":
		addGoal(c, stringFromMap(m.Values, "side"), stringFromMap(m.Values, "speed"))
		sendMatch(c, "")
	case "removeGoal":
	case "startLobby":
	default:

	}
}

func stringFromMap(m map[string]string, key string) string {
	for k, v := range m {
		if k == key {
			return v
		}
	}
	return ""
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
	// Command to execute
	Command string `json:"command"`

	// Insert all possible command structures
	Values map[string]string `json:"values"`
}
