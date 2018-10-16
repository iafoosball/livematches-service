package main

import (
	"log"
	"strings"
)

func handleCommunication(c *Client, input string) {
	inputs := strings.Split(input, "-")
	if inputs[0] == "2018" {
		return
	}
	inputs = strings.Split(input, ":")
	log.Println(input)
	if checkClient(c) {
		switch inputs[0] {
		case "create":
			//	implement create new match on matches-service
			c.match = match(inputs[1])
			// Tell all users on livematches service
			sendAll(c, "Created match: "+inputs[1])

		case "join":
			c.match = joinMatch(c, inputs[1])

			if c.match != nil {
				log.Println("joined-" + *c.userID + "-" + *c.username)
				sendMatch(c, "joined-"+*c.userID+"-"+*c.username)
			} else {
				sendPrivate(c, "you could not join")

			}
		}
	} else {
		switch inputs[0] {
		case "id":
			c.userID = &inputs[1]
		case "name":
			c.username = &inputs[1]
		default:
			sendPrivate(c, "provide your id and name")
		}
	}
}

func checkClient(c *Client) bool {
	if c.userID != nil && c.username != nil {
		return true
	}
	return false
}
