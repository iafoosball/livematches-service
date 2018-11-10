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
func createMatch(c *Client, tableID string) {
	c.table.ID = tableID
	for i, match := range LiveMatches {
		if match.M.TableID == c.table.ID {
			LiveMatches = append(LiveMatches[:i], LiveMatches[i+1:]...)
		}
	}
	match := newMatch(tableID)
	go match.runMatch()
	LiveMatches = append(LiveMatches, match)
	c.liveMatch = match
	c.liveMatch.Register <- c
}

func newMatch(tableID string) *LiveMatch {
	return &LiveMatch{
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		MatchCast:  make(chan []byte),
		M: &models.Match{
			Positions: &models.MatchPositions{},
			Settings: &models.MatchSettings{
				Bet:             false,
				Drunk:           false,
				FreeGame:        false,
				MaxGoals:        10,
				MaxTime:         0,
				OneOnOne:        true,
				Payed:           false,
				Rated:           false,
				SwitchPositions: false,
				Tournament:      false,
				TwoOnOne:        false,
				TwoOnTwo:        false,
			},
			Users:     []*models.MatchUsersItems0{},
			Started:   false,
			TableID:   tableID,
			ScoreBlue: 0,
			ScoreRed:  0,
		},
	}
}

// If match is finished it is send to matches-service and stored their
// sending data still needs implementation
func closeMatch(c *Client) {
	SendMatch(*c.liveMatch)
	for cl, _ := range c.liveMatch.Clients {
		closeUser(cl)
	}
	id := c.liveMatch.M.TableID
	for i, l := range LiveMatches {
		if l.M.TableID == id {
			LiveMatches = append(LiveMatches[:i], LiveMatches[i+1:]...)
		}
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
	M *models.Match `json:"-"`

	// started indicates if the match started
	Started bool `json:"started,omitempty"`

	// holds the data of the Goals for a LiveMatch
	Goals []*models.Goal `json:"-"`
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
