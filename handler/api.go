package handler

import (
	"encoding/json"
)

var (
	err error
	m   *message
	b   bool
)

const (
	// Start: For Users
	setPosition = "setPosition"
	// { \"command\": \"setPosition\", \"values\": { \"side\": \"red\", \"position\": \"attack\" }}
	setColor = "setColor"
	// { \"command\": \"setColor\", \"values\": { \"color\": \"green\" }}
	setUsername = "setUsername"
	// { \"command\": \"setUsername\", \"values\": { \"username\": \"joe\" }}
	setBet = "setBet"
	// { \"command\": \"setBet\", \"values\": { \"bet\": 123 }}
	ready = "ready"
	// { \"command\": \"ready\", \"values\": { \"ready\": true }}
	// Start: For Admin
	twoOnTwo = "twoOnTwo"
	// { \"command\": \"twoOnTwo\", \"values\": { \"twoOnTwo\": true }}
	twoOnOne = "twoOnOne"
	// { \"command\": \"oneOnTwo\", \"values\": { \"oneOnTwo\": true }}
	oneOnOne = "oneOnOne"
	// { \"command\": \"oneOnOne\", \"values\": { \"oneOnOne\": true }}
	switchPositions = "switchPositions"
	// { \"command\": \"switchPositions\", \"values\": { \"switchPositions\": true }}
	bet = "bet"
	// { \"command\": \"bet\", \"values\": { \"bet\": true }}
	maxGoals = "maxGoals"
	// { \"command\": \"maxGoals\", \"values\": { \"maxGoals\": 10 }}
	tournament = "tournament"
	// { \"command\": \"tournament\", \"values\": { \"tournament\": true }}
	startMatch = "startMatch"
	// { \"command\": \"startMatch\", \"values\": { }}
	drunk = "drunk"
	// { \"command\": \"drunk\", \"values\": { \"drunk\": true }}
	freeGame = "freeGame"
	// { \"command\": \"freeGame\", \"values\": { \"freeGame\": true }}
	payed = "payed"
	// { \"command\": \"payed\", \"values\": { \"payed\": true }}
	maxTime = "maxTime"
	// { \"command\": \"maxTime\", \"values\": { \"maxTime\": 600 }}
	rated = "rated"
	// { \"command\": \"rated\", \"values\": { \"rated\": true }}
	cancelMatch = "cancelMatch"
	// { \"command\": \"candelMatch\", \"values\": { }}
	kickUser = "kickUser"
	// { \"command\": \"kickUser\", \"values\": { \"kickUser\": "userID" }}
	// Start: For Table, possible by admin as well
	addGoal = "addGoal"
	// { \"command\": \"addGoal\", \"values\": { \"speed\": 12, \"side\": \"blue\", \"position\": \"attack\"  }}
	removeGoal = "removeGoal"
)

func handleCommunication(c *Client, message []byte) {
	if m, b = unmarshal(message); !b {
		return
	}
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
	case setPosition:
		position(c, stringFromMap(m.Values, "position"), stringFromMap(m.Values, "side"))
		//sendMatchData(c)
	case setColor:
		leaveMatch(c)
	case setUsername:
	case setBet:
	}
	//sendMatchData(c)
}

func handleAdmin(c *Client, m *message) {
	handleUsers(c, m)
	switch m.Command {
	case "startMatch":
	case "double":
	case "rated":
	case "maxTime":
	case "maxGoals":
	case "switchPositions":
	}
}

func handleTable(c *Client, m *message) {
	//log.Println("Hanlde table " + c.table.ID + " with cmd: " + m.Command)
	switch m.Command {
	case "closeMatch":
		closeMatch(c)
	case "addGoal":
		addgoal(c, stringFromMap(m.Values, "side"), stringFromMap(m.Values, "speed"))
		sendMatch(c, "")
	case "removeGoal":
	case "startLobby":
	default:

	}
}

//stringFromMap returns a string for a certain id from a map
func stringFromMap(m map[string]string, key string) string {
	for k, v := range m {
		if k == key {
			return v
		}
	}
	return ""
}

//boolFromMap returns a string for a certain id from a map
func boolFromMap(m map[string]bool, key string) bool {
	for k, v := range m {
		if k == key {
			return v
		}
	}
	return false
}

//numberFromMap returns a string for a certain id from a map
func numberFromMap(m map[string]float64, key string) float64 {
	for k, v := range m {
		if k == key {
			return v
		}
	}
	return 0
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
