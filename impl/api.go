package impl

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
	setPosition = "setPosition"
	// { "command": "setPosition", "values": { "side": "red", "position": "attack" }}
	setColor = "setColor"
	// { "command": "setColor", "values": { "color": "green" }}
	setUsername = "setUsername"
	// { "command": "setUsername", "values": { "username": "joe" }}
	setBet = "setBet"
	// { "command": "setBet", "values": { "bet": 123 }}
	ready = "ready"
	// { "command": "ready", "values": { "ready": true }}
	leaveMatch = "leaveMatch"
	// { \"command\": \"leaveMatch\", \"values\": { }}

	// Start: For Admin
	twoOnTwo = "twoOnTwo"
	// { "command": "settings", "values": { "twoOnTwo": true }}
	twoOnOne = "twoOnOne"
	// { "command": "settings", "values": { "oneOnTwo": true }}
	oneOnOne = "oneOnOne"
	// { "command": "settings", "values": { "oneOnOne": true }}
	switchPositions = "switchPositions"
	// { "command": "settings", "values": { "switchPositions": true }}
	bet = "bet"
	// { "command": "settings", "values": { "bet": true }}
	maxGoals = "maxGoals"
	// { "command": "settings", "values": { "maxGoals": 10 }}
	tournament = "tournament"
	// { "command": "settings", "values": { "tournament": true }}
	tournamentMode = "tournamentMode"
	// { "command": "settings", "values": { "tournamentMode": true }}
	started = "started"
	// { "command": "started", "values": { }}
	drunk = "drunk"
	// { "command": "settings", "values": { "drunk": true }}
	freeGame = "freeGame"
	// { "command": "settings", "values": { "freeGame": true }}
	payed = "payed"
	// { "command": "settings", "values": { "payed": true }}
	maxTime = "maxTime"
	// { "command": "settings", "values": { "maxTime": 600 }}
	rated = "rated"
	// { "command": "settings", "values": { "rated": true }}
	kickUser = "kickUser"
	// { "command": "settings", "values": { "kickUser": "userID" }}

	// Start: For Table, possible by admin as well
	stopMatch = "cancelMatch"
	// { "command": "cancelMatch", "values": { }}
	addGoal = "addGoal"
	// { \"command\": \"addGoal\", \"values\": { \"speed\": 12, \"side\": \"blue\", \"setposition\": \"attack\"  }}
	removeGoal = "removeGoal"
)

func handleCommunication(c *Client, message []byte) {
	if m, b = unmarshal(message); !b {
		return
	}
	log.Println(string(message))
	if c.IsUser && c.User.Admin {
		handleAdmin(c, m)
	} else if c.IsUser {
		handleUsers(c, m)
	} else {
		handleTable(c, m)
	}
}

func handleUsers(c *Client, m *message) {
	switch m.Command {
	case setPosition:
		setposition(c, stringFromMap(m.Values, "position"), stringFromMap(m.Values, "side"))
	case setColor:
		setcolor(c, stringFromMap(m.Values, "color"))
	case setUsername:
		setusername(c, stringFromMap(m.Values, "setusername"))
	case setBet:
		setbet(c, intFromMap(m.Values, "bet"))
	case ready:
		setReady(c, boolFromMap(m.Values, ready))
	case leaveMatch:
		//closeUser(c)
	}
}

func handleAdmin(c *Client, m *message) {
	handleTable(c, m)
	handleUsers(c, m)
	if m.Command == "settings" {
		for k, v := range m.Values {
			switch k {
			case rated:
				setRated(c, v.(bool))
			case maxTime:
				maxtime(c, int64(v.(float64)))
			case maxGoals:
				maxgoals(c, int64(v.(float64)))
			case switchPositions:
				switchpositions(c, v.(bool))
			case twoOnTwo:
				twoontwo(c, v.(bool))
			case twoOnOne:
				twoonone(c, v.(bool))
			case oneOnOne:
				oneonone(c, v.(bool))
			case bet:
				isBet(c, v.(bool))
			case tournament:
				isTournament(c, v.(bool))
			case tournamentMode:
				tournamentmode(c, v.(bool))
			case drunk:
				isDrunk(c, v.(bool))
			case payed:
				isPayed(c, v.(bool))
			case kickUser:
				kickuser(c, v.(string))
			}
			break
		}
	}

}

func handleTable(c *Client, m *message) {
	switch m.Command {
	case started:
		start(c)
	case stopMatch:
		stopmatch(c)
	case addGoal:
		addgoal(c, stringFromMap(m.Values, "side"), numberFromMap(m.Values, "speed"))
	case removeGoal:
		removegoal(c, stringFromMap(m.Values, "side"))
	case "settings":
		if stringFromMap(m.Values, "freeGame") == freeGame {
			freegame(c, boolFromMap(m.Values, "freeGame"))
		}
	}
}

//stringFromMap returns a string for a certain id from a map
func stringFromMap(m map[string]interface{}, key string) string {
	for k, v := range m {
		if k == key {
			return v.(string)
		}
	}
	return ""
}

//boolFromMap returns a string for a certain id from a map
func boolFromMap(m map[string]interface{}, key string) bool {
	for k, v := range m {
		if k == key {
			return v.(bool)
		}
	}
	return false
}

//intFromMap returns a string for a certain id from a map
func intFromMap(m map[string]interface{}, key string) int64 {
	for k, v := range m {
		if k == key {
			return int64(v.(float64))
		}
	}
	return 0
}

//numberFromMap returns a string for a certain id from a map
func numberFromMap(m map[string]interface{}, key string) float64 {
	for k, v := range m {
		if k == key {
			return v.(float64)
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
		log.Println(string(msg))
		log.Println(err)
		return m, false
	}
	return m, true
}

type message struct {
	// Command to execute
	Command string `json:"command"`

	// Insert all possible command structures
	Values map[string]interface{} `json:"values"`
}
