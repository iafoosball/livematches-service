package handler

import (
	"github.com/iafoosball/livematches-service/models"
)

var matches = []liveMatch{}

// Either return already existing liveMatch or create new one
func match(id string) *liveMatch {
	for _, match := range matches {
		if match.matchID == id {
			return &match
		}
	}
	match := liveMatch{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		matchCast:  make(chan []byte),
		matchData:  models.Match{},
		matchID:    id,
	}
	matches = append(matches, match)
	go match.runMatch()
	return &match
}

func joinMatch(c *Client, id string) *liveMatch {
	for _, match := range matches {
		if match.matchID == id {
			match.clients[c] = true
			return &match
		}
	}
	return nil
}

func leaveMatch() {

}

func (m *liveMatch) runMatch() {
	for {
		select {
		case client := <-m.register:
			m.clients[client] = true
		case client := <-m.unregister:
			if _, ok := m.clients[client]; ok {
				delete(m.clients, client)
				close(client.send)
			}
		case message := <-m.matchCast:
			for client := range m.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(m.clients, client)
				}
			}
		}
	}
}

type liveMatch struct {
	// Registered clients.
	clients map[*Client]bool

	// Outbound messages for all users inside a liveMatch
	matchCast chan []byte

	// The liveMatch ID
	matchID string

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	// holds the data of the liveMatch
	matchData models.Match
}
