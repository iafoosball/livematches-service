package impl

import (
	"errors"
	"github.com/iafoosball/livematches-service/models"
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
		Settings:   &models.MatchSettings{},
		Users:      []*models.MatchUsersItems0{},
	}
}

func joinMatch(c *Client, id string) {
	for _, match := range LiveMatches {
		if match.TableID == id {
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

// Used by pi to finish a Match
// If match is finished it is send to matches-service and stored their
func closeMatch(c *Client) {
	for cl, _ := range c.liveMatch.Clients {
		leavematch(cl)
	}
	id := c.liveMatch.TableID
	for i, l := range LiveMatches {
		if l.TableID == id {
			SendMatch(*l)
			LiveMatches[i] = LiveMatches[len(LiveMatches)-1]
			LiveMatches = LiveMatches[:len(LiveMatches)-1]
		}
	}
}

func kickuser(c *Client, id string) {
	for _, u := range c.liveMatch.Users {
		if u.ID == id {
			leavematch(c)
			return
		}
	}
}

// leavematch is used by users to leave a Match
func leavematch(c *Client) {
	c.liveMatch.Unregister <- c
	for i, p := range c.liveMatch.Users {
		if p.ID == c.user.ID {
			setposition(c, "", "")
			c.liveMatch.Users[i] = c.liveMatch.Users[len(c.liveMatch.Users)-1]
			c.liveMatch.Users = c.liveMatch.Users[:len(c.liveMatch.Users)-1]
		}
	}
	sendMatchData(c)
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

	// started indicates if the match started
	Started bool `json:"started,omitempty"`

	// holds the data of the Goals for a LiveMatch
	Goals []*models.Goal `json:"-"`

	//Start auto generated stuff
	// Was the game completed.
	Completed bool `json:"completed,omitempty"`

	// the datetime when the match ends
	EndTime string `json:"endTime,omitempty"`

	// positions
	Positions *models.MatchPositions `json:"positions,omitempty"`

	// score blue
	ScoreBlue int64 `json:"scoreBlue,omitempty"`

	// score red
	ScoreRed int64 `json:"scoreRed,omitempty"`

	// settings
	Settings *models.MatchSettings `json:"settings,omitempty"`

	// the datetime when the game ends
	StartTime string `json:"startTime,omitempty"`

	// the id of table
	TableID string `json:"tableID,omitempty"`

	// users
	Users []*models.MatchUsersItems0 `json:"users"`

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
