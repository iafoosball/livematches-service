package main

import (
	"strings"
)

func handleCommunication(c *Client, input string) {
	inputs := strings.Split(input, ":")
	switch inputs[0] {

	case "id":
		c.userID = &inputs[1]

	case "name":
		c.username = &inputs[1]

	case "create":
		//	implement create new match on matches-service
		c.match = match(inputs[1])
		// Tell all users on livematches service
		sendAll(c, "Created match: "+inputs[1])

	case "join":
		c.match = joinMatch(c, inputs[1])

		if c.match != nil {
			sendMatch(c, "joined-"+*c.userID+"-"+*c.username)
		} else {
			sendPrivate(c, "you could not join")

		}
	}
}
