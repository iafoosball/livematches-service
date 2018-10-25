package handler

import (
	"errors"
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
		if match.TableID == c.table.ID {
			LiveMatches[i] = LiveMatches[len(LiveMatches)-1]
			LiveMatches = LiveMatches[:len(LiveMatches)-1]
		}
	}
	match := newMatch()
	match.TableID = tableID
	go match.runMatch()
	LiveMatches = append(LiveMatches, match)
	c.liveMatch = match
	c.liveMatch.Register <- c
	return true
}

func newMatch() *LiveMatch {
	return &LiveMatch{
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		MatchCast:  make(chan []byte),
		MatchData:  &models.Match{},
		Positions:  &models.MatchPositions{},
		Users:      []*models.User{},
		Started:    false,
	}
}

func joinMatch(c *Client, id string) {
	for _, match := range LiveMatches {
		if match.TableID == id {
			log.Println(len(match.Users))
			if len(match.Users) > 3 {
				closeClient(c)
				return
			}
			c.liveMatch = match
			c.liveMatch.Users = append(c.liveMatch.Users, c.user)
			c.liveMatch.Register <- c
		}
		return
	}
	handleErr(err)
	closeClient(c)
}

// startMatch writes everything to the Match object.
// Before users etc. is stored on the livematch
func startmatch() {

}

// Used by pi to finish a Match
// If match is finished it is send to matches-service and stored their
func closeMatch(c *Client) {
	for cl, _ := range c.liveMatch.Clients {
		leaveMatch(cl)
	}
	id := c.liveMatch.TableID
	for i, l := range LiveMatches {
		if l.TableID == id {
			LiveMatches[i] = LiveMatches[len(LiveMatches)-1]
			LiveMatches = LiveMatches[:len(LiveMatches)-1]
		}
	}
}

// Used by users to leave a Match
func leaveMatch(c *Client) {
	c.liveMatch.Unregister <- c
	for i, p := range c.liveMatch.Users {
		if p.ID == c.user.ID {
			c.liveMatch.Users[i] = c.liveMatch.Users[len(c.liveMatch.Users)-1]
			c.liveMatch.Users = c.liveMatch.Users[:len(c.liveMatch.Users)-1]
		}
	}
}

func position(c *Client, position string, side string) {
	if position == "attack" && side == "blue" {
		if c.liveMatch.Positions.BlueAttach == "" {
			c.liveMatch.Positions.BlueAttach = c.user.ID
		}
	} else if position == "defense" && side == "blue" {
		if c.liveMatch.Positions.BlueDefense == "" {
			c.liveMatch.Positions.BlueDefense = c.user.ID
		}
	} else if position == "attack" && side == "red" {
		if c.liveMatch.Positions.RedAttack == "" {
			c.liveMatch.Positions.RedAttack = c.user.ID
		}
	} else if position == "defense" && side == "red" {
		if c.liveMatch.Positions.RedDefense == "" {
			c.liveMatch.Positions.RedDefense = c.user.ID
		}
	}
}

func addgoal(c *Client, side string, speed string) {
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

type LiveMatch struct {
	// Registered Clients.
	Clients map[*Client]bool `json:"-"`

	// Outbound messages for all users inside a LiveMatch
	MatchCast chan []byte `json:"-"`

	// Register requests from the Clients.
	Register chan *Client `json:"-"`

	// Unregister requests from Clients.
	Unregister chan *Client `json:"-"`

	// holds the data of the LiveMatch
	MatchData *models.Match `json:"-"`

	// holds the data of the Goals for a LiveMatch
	Goals []*models.Goal `json:"goals"`

	// list of all Players
	Users []*models.User `json:"users"`

	//Start auto generated stuff
	// Is this game with bets
	Bet bool `json:"bet,omitempty"`

	// Was the game completed.
	Completed bool `json:"completed,omitempty"`

	// the datetime when the match ends
	EndTime string `json:"endTime,omitempty"`

	// free game
	FreeGame bool `json:"freeGame,omitempty"`

	// The maximum number of goals for this game. If a time is specified the first criteria which is true will stop the match.
	MaxGoals *int64 `json:"maxGoals,omitempty"`

	// The maximum tim in sec for this game. If a max goals is specified the first criteria which is true will stop the match.
	MaxTime int64 `json:"maxTime,omitempty"`

	// one on one
	OneOnOne bool `json:"oneOnOne,omitempty"`

	// payed
	Payed bool `json:"payed,omitempty"`

	// positions
	Positions *models.MatchPositions `json:"positions,omitempty"`

	// A match can be rated, ie a ranked game with points, or without.
	RatedMatch bool `json:"ratedMatch,omitempty"`

	// score blue
	ScoreBlue int64 `json:"scoreBlue,omitempty"`

	// score red
	ScoreRed int64 `json:"scoreRed,omitempty"`

	// the datetime when the game ends
	StartTime string `json:"startTime,omitempty"`

	// started
	Started bool `json:"started,omitempty"`

	// Switch the position after 50% of the goal (time or goals) is reached.
	SwitchPosition bool `json:"switchPosition,omitempty"`

	// the id of table
	TableID string `json:"tableID,omitempty"`

	// tournament
	Tournament bool `json:"tournament,omitempty"`

	// two on one
	TwoOnOne bool `json:"twoOnOne,omitempty"`

	// two on two
	TwoOnTwo bool `json:"twoOnTwo,omitempty"`

	// Can be either "red" or "blue"
	Winner string `json:"winner,omitempty"`
}

func (m *LiveMatch) runMatch() {
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if recover() != nil {
			err = errors.New("Probably connection interrupt")
		}

	}()
	for {
		select {
		case client := <-m.Register:
			m.Clients[client] = true
		case client := <-m.Unregister:
			if _, ok := m.Clients[client]; ok {
				delete(m.Clients, client)
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
