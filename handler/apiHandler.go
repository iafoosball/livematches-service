package handler

import (
	"encoding/json"
	"fmt"
	"github.com/iafoosball/livematches-service/handler"
	"github.com/iafoosball/livematches-service/models"
	"log"
	"strings"
)

var (
	err error
)

func handleCommunication(c *Client, body string) {
	if !isJSON(body) {
		return
	}
	var m message
	json.Unmarshal([]byte(body), &m)
	if !checkClient(c, m) {
		unmarshalUser(c, m)
		return
	}
	if !c.user.Raspberry && !c.user.Visitor {
		for k, v := range c.hub.clients {
			fmt.Println("k:", k, "v:", v)
		}
	}

	//if checkClient()
	//if checkClient(c) {
	//
	//	switch inputs[0] {
	//	case "create":
	//		//	implement create new liveMatch on matches-service
	//		c.liveMatch = match(inputs[1])
	//		// Tell all users on livematches service
	//		sendAll(c, "Created liveMatch: "+inputs[1])
	//
	//	case "join":
	//		c.liveMatch = joinMatch(c, inputs[1])
	//
	//		if c.liveMatch != nil {
	//			log.Println("joined-" + *c.user.Username + "-" + *c.username)
	//			sendMatch(c, "joined-"+ *c.user.ID+"-"+*c.username)
	//		} else {
	//			sendPrivate(c, "you could not join")
	//
	//		}
	//	}
	//} else {
	//	switch inputs[0] {
	//	case "id":
	//		c.userID = &inputs[1]
	//	case "name":
	//		c.username = &inputs[1]
	//	default:
	//		sendPrivate(c, "provide your id and name")
	//	}
	//}
}

func unmarshalUser(c *handler.Client, body string) {
	user := &models.User{}
	err = json.NewDecoder(strings.NewReader(string(body))).Decode(&user)
	if err != nil {
		log.Println(err)
		return
	}
	c.user = user
}

func checkClient(c *Client, m message) bool {
	if c.user.ID != nil && c.user.Username != nil {
		return true
	}
	return false
}

func isJSON(str string) bool {
	s := strings.Split(str, "-")
	if s[0] == "2018" {
		return false
	}
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

type message struct {

	// The user key
	Key string `json:"_key"`

	// The username
	Username string `json:"username"`

	// The table id
	Table bool `json:"table"`

	// Command to execute
	Command string `json:"command"`

	// Insert all possible command structures
	Values map[string]string `json:"values"`
}
