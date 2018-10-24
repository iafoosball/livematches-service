package handler

import (
	"github.com/iafoosball/livematches-service/models"
	"log"
	"time"
)

// operations for live LiveMatches

// All active lobbies and matches are registered here.
var LiveMatches = []*LiveMatch{}

// Creates a new match. (Either return already existing LiveMatch or create new one)
// How to handle contradictions??? If there is an already open match what to do....
func createMatch(c *Client, tableID string) bool {
	c.table.ID = tableID
	for i, match := range LiveMatches {
		if match.MatchID == c.table.ID {
			LiveMatches[i] = LiveMatches[len(LiveMatches)-1]
			LiveMatches = LiveMatches[:len(LiveMatches)-1]
		}
	}
	match := &LiveMatch{
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		MatchCast:  make(chan []byte),
		MatchData:  &models.Match{},
		MatchID:    c.table.ID,
		Started:    false,
	}
	go match.initMatch()
	LiveMatches = append(LiveMatches, match)
	c.liveMatch = match
	c.liveMatch.Register <- c
	return true
}

func joinMatch(c *Client, id string) bool {
	log.Println(id)
	for _, match := range LiveMatches {
		log.Println(match)
		if match.MatchID == id {
			c.liveMatch = match
			c.liveMatch.Register <- c
			c.liveMatch.Players = append(c.liveMatch.Players, c.user)
			log.Println(c.liveMatch)
		}
		return true
	}
	handleErr(err)
	return false
}

// startMatch writes everything to the Match object.
// Before users etc. is stored on the livematch
func startMatch() {

}

// Used by pi to finish a Match
// If match is finished it is send to matches-service and stored their
func closeMatch(c *Client) {
	for cl, _ := range c.liveMatch.Clients {
		leaveMatch(cl)
	}
	id := c.liveMatch.MatchID
	for i, l := range LiveMatches {
		if l.MatchID == id {
			LiveMatches[i] = LiveMatches[len(LiveMatches)-1]
			LiveMatches = LiveMatches[:len(LiveMatches)-1]
		}
	}
}

// Used by users to leave a Match
func leaveMatch(c *Client) {
	c.liveMatch.Unregister <- c
	for i, p := range c.liveMatch.Players {
		log.Println(c.user.ID)
		log.Println(p.ID)

		if p.ID == c.user.ID {
			log.Println("delete user from match")
			c.liveMatch.Players[i] = c.liveMatch.Players[len(c.liveMatch.Players)-1]
			c.liveMatch.Players = c.liveMatch.Players[:len(c.liveMatch.Players)-1]
		}
	}
}

func addGoal(c *Client, side string, speed string) {
	c.liveMatch.Goals = append(c.liveMatch.Goals, &models.Goal{
		Side:     side,
		Speed:    speed,
		Datetime: time.Now().Format(time.RFC3339),
	})
	if side == "red" {
		c.liveMatch.ScoreRed++
	} else if side == "blue" {
		c.liveMatch.ScoreBlue++
	}
}

type LiveMatch struct {
	// Registered Clients.
	Clients map[*Client]bool `json:"-"`

	// Outbound messages for all users inside a LiveMatch
	MatchCast chan []byte `json:"-"`

	// The LiveMatch ID
	MatchID string `json:"matchID"`

	// Started is True if match is active
	Started bool `json:"started"`

	// Register requests from the Clients.
	Register chan *Client `json:"-"`

	// Unregister requests from Clients.
	Unregister chan *Client `json:"-"`

	// holds the data of the LiveMatch
	MatchData *models.Match `json:"-"`

	// holds the data of the Goals for a LiveMatch
	Goals []*models.Goal `json:"-"`

	// list of all Players
	Players []*models.User `json:"users"`

	// list of all Visitors
	Visitors []*models.User `json:"Visitors"`

	ScoreRed int `json:"score_red"`

	ScoreBlue int `json:"score_blue"`
}

func (m *LiveMatch) initMatch() {
	for {
		select {
		case client := <-m.Register:
			m.Clients[client] = true
		case client := <-m.Unregister:
			if _, ok := m.Clients[client]; ok {

			}
		case message := <-m.MatchCast:
			for client := range m.Clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(m.Clients, client)
				}
			}
		}
	}
}
