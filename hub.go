package main

type Hub struct {
	// Register clients.
	clients map[*Client]bool

	// map clients to a match
	matches []Match

	// Inbound messages from the clients.
	message chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		message:    make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		matches:    make([]Match, 0),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.message:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
