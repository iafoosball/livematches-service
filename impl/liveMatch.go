package impl

import (
	"errors"
	"github.com/iafoosball/livematches-service/models"
	"log"
)

// operations for live LiveMatches

// All active lobbies and matches are registered here.
var LiveMatches = []*LiveMatch{}

// Creates a new match. (Either return already existing LiveMatch or create new one)
// How to handle contradictions??? If there is an already open match what to do....
func createMatch(c *Client, tableID string) {
	c.Table.ID = tableID
	for i, match := range LiveMatches {
		if match.M.TableID == c.Table.ID {
			LiveMatches = append(LiveMatches[:i], LiveMatches[i+1:]...)
		}
	}
	match := newMatch(tableID)
	go match.runMatch()
	LiveMatches = append(LiveMatches, match)
	c.LiveMatch = match
	c.LiveMatch.Register <- c
}

func newMatch(tableID string) *LiveMatch {
	return &LiveMatch{
		Clients:    make(map[*Client]string),
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
func stopmatch(c *Client) {
	go SendMatch(c.LiveMatch)
	//for cl, _ := range c.LiveMatch.Clients {
	//	closeUser(cl)
	//}
	c.LiveMatch.M.Started = false
	c.LiveMatch.M.ScoreBlue = 0
	c.LiveMatch.M.ScoreRed = 0
	sendMatchData(c)
}

type LiveMatch struct {
	// Registered Clients.
	Clients map[*Client]string `json:"-"`

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
			m.Clients[client] = client.ID
		case client := <-m.Unregister:
			log.Println("unregister")
			if _, ok := m.Clients[client]; ok {
				log.Println("unregistered")
				delete(m.Clients, client)
			}
		case message := <-m.MatchCast:
			for client := range m.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(m.Clients, client)
				}
			}
		}
	}
}
