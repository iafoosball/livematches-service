package impl

import (
	"encoding/json"
	"github.com/iafoosball/livematches-service/models"
	"log"
	"time"
)

var (
	err error
	m   *message
	b   bool
)

const (
	// Start: For Users
	setPosition = "setPosition"
	// { \"command\": \"setPosition\", \"values\": { \"side\": \"red\", \"setposition\": \"attack\" }}
	setColor = "setColor"
	// { \"command\": \"setColor\", \"values\": { \"color\": \"green\" }}
	setUsername = "setUsername"
	// { \"command\": \"setUsername\", \"values\": { \"setusername\": \"joe\" }}
	setBet = "setBet"
	// { \"command\": \"setBet\", \"values\": { \"bet\": 123 }}
	ready = "ready"
	// { \"command\": \"ready\", \"values\": { \"ready\": true }}
	leaveMatch = "leaveMatch"
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
	// { \"command\": \"addGoal\", \"values\": { \"speed\": 12, \"side\": \"blue\", \"setposition\": \"attack\"  }}
	removeGoal = "removeGoal"
)

func handleCommunication(c *Client, message []byte) {
	if m, b = unmarshal(message); !b {
		return
	}
	log.Println(string(message))
	if c.isUser && c.user.Admin {
		handleAdmin(c, m)
	}
	if c.isUser {
		handleUsers(c, m)
	}
	if !c.isUser || c.user.Admin {
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
		leavematch(c)
	}
}

func handleAdmin(c *Client, m *message) {
	handleUsers(c, m)
	switch m.Command {
	case startMatch:
		startmatch(c)
	case rated:
		setRated(c, boolFromMap(m.Values, rated))
	case maxTime:
		maxtime(c, intFromMap(m.Values, maxTime))
	case maxGoals:
		maxgoals(c, intFromMap(m.Values, maxGoals))
	case switchPositions:
		switchpositions(c, boolFromMap(m.Values, switchPositions))
	case twoOnTwo:
		twoontwo(c, boolFromMap(m.Values, twoOnTwo))
	case twoOnOne:
		twoonone(c, boolFromMap(m.Values, twoOnOne))
	case oneOnOne:
		oneonone(c, boolFromMap(m.Values, oneOnOne))
	case bet:
		isBet(c, boolFromMap(m.Values, bet))
	case tournament:
		isTournament(c, boolFromMap(m.Values, tournament))
	case drunk:
		isDrunk(c, boolFromMap(m.Values, drunk))
	case payed:
		isPayed(c, boolFromMap(m.Values, drunk))
	case kickUser:
		kickuser(c, stringFromMap(m.Values, kickUser))
	}
}

func handleTable(c *Client, m *message) {
	//log.Println("Hanlde table " + c.table.ID + " with cmd: " + m.Command)
	switch m.Command {
	case cancelMatch:
		closeMatch(c)
	case addGoal:
		addgoal(c, stringFromMap(m.Values, "side"), numberFromMap(m.Values, "speed"))
	case removeGoal:
		removegoal(c)
	case freeGame:
		freegame(c, boolFromMap(m.Values, freeGame))
	}
}

func addgoal(c *Client, side string, speed float64) {
	c.liveMatch.Goals = append(c.liveMatch.Goals, &models.Goal{
		Side:     side,
		Speed:    speed,
		DateTime: time.Now().Format(time.RFC3339),
	})
	if side == "red" {
		c.liveMatch.ScoreRed++
	} else if side == "blue" {
		c.liveMatch.ScoreBlue++
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
			return v.(int64)
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
		//log.Println(err)
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
