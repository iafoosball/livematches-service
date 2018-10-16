package main

type Match struct {

	// Registered clients.
	clients map[*Client]bool

	// Outbound messages for all users inside a match
	matchBroadcast chan []byte

	// The match ID
	matchID string

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	// Only admin can create new matches.
	// Usually only the tablet on the table is admin
	admin bool
}

var matches = make([]Match, 0)

// Either return already existing match or create new one
func match(id string) *Match {
	for _, match := range matches {
		if match.matchID == id {
			return &match
		}
	}
	match := Match{
		clients:        make(map[*Client]bool),
		register:       make(chan *Client),
		unregister:     make(chan *Client),
		matchBroadcast: make(chan []byte),
		matchID:        id,
	}
	matches = append(matches, match)
	go match.runMatch()
	return &match
}

func joinMatch(c *Client, id string) *Match {
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

func (m *Match) runMatch() {
	for {
		select {
		case client := <-m.register:
			m.clients[client] = true
		case client := <-m.unregister:
			if _, ok := m.clients[client]; ok {
				delete(m.clients, client)
				close(client.send)
			}
		case message := <-m.matchBroadcast:
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
