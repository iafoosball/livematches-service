package impl

import (
	"time"
)

func closeUser(c *Client) {
	// reset position and delete from users
	for i, u := range c.liveMatch.M.Users {
		if u.ID == c.user.ID {
			resetPosition(c)
			c.liveMatch.M.Users = append(c.liveMatch.M.Users[:i], c.liveMatch.M.Users[i+1:]...)
			break
		}
	}
	for u, _ := range c.liveMatch.Clients {
		if !u.isUser {
			sendMatchData(u)
		}

	}
	c.liveMatch.Unregister <- c
	c.hub.unregister <- c
	c.end <- true
}

func setusername(c *Client, username string) {
	c.user.Username = username
	sendMatchData(c)
}

func setposition(c *Client, position string, side string) {
	resetPosition(c)
	if position == "attack" && side == "blue" {
		if c.liveMatch.M.Positions.BlueAttack == "" {
			c.liveMatch.M.Positions.BlueAttack = c.user.ID
		}
	} else if position == "defense" && side == "blue" {
		if c.liveMatch.M.Positions.BlueDefense == "" {
			c.liveMatch.M.Positions.BlueDefense = c.user.ID
		}
	} else if position == "attack" && side == "red" {
		if c.liveMatch.M.Positions.RedAttack == "" {
			c.liveMatch.M.Positions.RedAttack = c.user.ID
		}
	} else if position == "defense" && side == "red" {
		if c.liveMatch.M.Positions.RedDefense == "" {
			c.liveMatch.M.Positions.RedDefense = c.user.ID
		}
	}
	sendMatchData(c)
}

func resetPosition(c *Client) {
	if c.liveMatch.M.Positions.RedDefense == c.user.ID {
		c.liveMatch.M.Positions.RedDefense = ""
	} else if c.liveMatch.M.Positions.RedAttack == c.user.ID {
		c.liveMatch.M.Positions.RedAttack = ""
	} else if c.liveMatch.M.Positions.BlueDefense == c.user.ID {
		c.liveMatch.M.Positions.BlueDefense = ""
	} else if c.liveMatch.M.Positions.BlueAttack == c.user.ID {
		c.liveMatch.M.Positions.BlueAttack = ""
	}
}

func joinMatch(c *Client, id string) {
	for _, match := range LiveMatches {
		if match.M.TableID == id {
			if len(match.M.Users) == 0 {
				c.user.Admin = true
			}
			c.liveMatch = match
			c.liveMatch.M.Users = append(c.liveMatch.M.Users, c.user)
			c.liveMatch.Register <- c
		}
		return
	}
	handleErr(err)
}

// Start: Admin settings
func startmatch(c *Client) {
	c.liveMatch.M.StartTime = time.Now().Format(time.RFC3339)
	c.liveMatch.M.Settings.StartMatch = true
	sendMatchData(c)
}

func switchpositions(c *Client, b bool) {
	c.liveMatch.M.Settings.SwitchPositions = b
	sendMatchData(c)
}

func twoontwo(c *Client, b bool) {
	resetPlayOption(c)
	c.liveMatch.M.Settings.TwoOnTwo = b
	sendMatchData(c)
}
func twoonone(c *Client, b bool) {
	resetPlayOption(c)
	c.liveMatch.M.Settings.TwoOnOne = b
	sendMatchData(c)
}
func oneonone(c *Client, b bool) {
	resetPlayOption(c)
	c.liveMatch.M.Settings.OneOnOne = b
	sendMatchData(c)
}

func resetPlayOption(c *Client) {
	c.liveMatch.M.Settings.OneOnOne = false
	c.liveMatch.M.Settings.TwoOnTwo = false
	c.liveMatch.M.Settings.TwoOnOne = false
}

func setcolor(c *Client, color string) {
	c.user.Color = color
	sendMatchData(c)
}

func isBet(c *Client, b bool) {
	c.liveMatch.M.Settings.Bet = b
	sendMatchData(c)
}
func isTournament(c *Client, b bool) {
	c.liveMatch.M.Settings.Tournament = b
	sendMatchData(c)
}
func isDrunk(c *Client, b bool) {
	c.liveMatch.M.Settings.Drunk = b
	sendMatchData(c)
}
func isPayed(c *Client, b bool) {
	c.liveMatch.M.Settings.Payed = b
	sendMatchData(c)
}

func freegame(c *Client, b bool) {
	c.liveMatch.M.Settings.FreeGame = b
	sendMatchData(c)
}

func setReady(c *Client, r bool) {
	c.user.Ready = r
	sendMatchData(c)
}

func setbet(c *Client, bet int64) {
	c.user.Bet = bet
	sendMatchData(c)
}

func setRated(c *Client, rated bool) {
	c.liveMatch.M.Settings.Rated = rated
	sendMatchData(c)
}

func maxtime(c *Client, time int64) {
	c.liveMatch.M.Settings.MaxTime = time
	sendMatchData(c)
}

func kickuser(c *Client, id string) {
	for client, _ := range c.liveMatch.Clients {
		if client.isUser && client.user.ID == id {
			closeUser(client)
			return
		}
	}
}

// End: Admin M.Settings

// For Admin and table
func maxgoals(c *Client, maxTime int64) {
	c.liveMatch.M.Settings.MaxGoals = maxTime
	sendMatchData(c)
}

func removegoal(c *Client, side string) {
	if side == "blue" {
		c.liveMatch.M.ScoreBlue--
	} else {
		c.liveMatch.M.ScoreRed--
	}
	sendMatchData(c)
}
