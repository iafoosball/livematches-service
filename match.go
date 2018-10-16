package main

import "log"

type Match struct {
	// Registered clients.
	clients map[*Client]bool

	// The match ID
	matchID string

	// Only admin can create new matches.
	// Usually only the tablet on the table is admin
	admin bool
}

// Either return already existing match or create new one
func match(hub *Hub, id string) Match {
	for _, match := range hub.matches {
		if match.matchID == id {
			return match
		}
	}
	match := Match{
		clients: make(map[*Client]bool),
		matchID: id,
	}
	hub.matches = append(hub.matches, match)
	return match
}

func joinMatch(matches *[]Match, client *Client, id string) bool {
	for _, match := range *matches {
		if match.matchID == id {
			match.clients[client] = true
			log.Println("match has following clients: ", match.clients)
			return true
		}
	}
	return false
}

func leaveMatch() {

}
