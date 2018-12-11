package impl

import (
	"github.com/iafoosball/livematches-service/models"
	"time"
)

func closeUser(c *Client) {
	// reset position and delete from users
	for i, u := range c.LiveMatch.M.Users {
		if u.ID == c.User.ID {
			resetPosition(c)
			c.LiveMatch.M.Users = append(c.LiveMatch.M.Users[:i], c.LiveMatch.M.Users[i+1:]...)
			break
		}
	}
	for u, _ := range c.LiveMatch.Clients {
		if !u.IsUser {
			sendMatchData(u)
		}

	}
	c.LiveMatch.Unregister <- c
	c.Hub.unregister <- c
	c.End <- true
}

func setusername(c *Client, username string) {
	c.User.Username = username
	sendMatchData(c)
}

func setposition(c *Client, position string, side string) {
	resetPosition(c)
	if position == "attack" && side == "blue" {
		if c.LiveMatch.M.Positions.BlueAttack == "" {
			c.LiveMatch.M.Positions.BlueAttack = c.User.ID
		}
	} else if position == "defense" && side == "blue" {
		if c.LiveMatch.M.Positions.BlueDefense == "" {
			c.LiveMatch.M.Positions.BlueDefense = c.User.ID
		}
	} else if position == "attack" && side == "red" {
		if c.LiveMatch.M.Positions.RedAttack == "" {
			c.LiveMatch.M.Positions.RedAttack = c.User.ID
		}
	} else if position == "defense" && side == "red" {
		if c.LiveMatch.M.Positions.RedDefense == "" {
			c.LiveMatch.M.Positions.RedDefense = c.User.ID
		}
	}
	sendMatchData(c)
}

func resetPosition(c *Client) {
	if c.LiveMatch.M.Positions.RedDefense == c.User.ID {
		c.LiveMatch.M.Positions.RedDefense = ""
	} else if c.LiveMatch.M.Positions.RedAttack == c.User.ID {
		c.LiveMatch.M.Positions.RedAttack = ""
	} else if c.LiveMatch.M.Positions.BlueDefense == c.User.ID {
		c.LiveMatch.M.Positions.BlueDefense = ""
	} else if c.LiveMatch.M.Positions.BlueAttack == c.User.ID {
		c.LiveMatch.M.Positions.BlueAttack = ""
	}
}

func joinMatch(c *Client, id string) {
	for _, match := range LiveMatches {
		if match.M.TableID == id {
			if len(match.M.Users) == 0 {
				c.User.Admin = true
			}
			c.LiveMatch = match
			c.LiveMatch.M.Users = append(c.LiveMatch.M.Users, c.User)
			c.LiveMatch.Register <- c
		}
		return
	}
	handleErr(err)
}

// Start: Admin settings
func start(c *Client) {
	c.LiveMatch.M.StartTime = time.Now().Format(time.RFC3339)
	c.LiveMatch.M.Started = true
	sendMatchData(c)
}

func switchpositions(c *Client, b bool) {
	c.LiveMatch.M.Settings.SwitchPositions = b
	sendMatchData(c)
}

func twoontwo(c *Client, b bool) {
	resetPlayOption(c)
	c.LiveMatch.M.Settings.TwoOnTwo = b
	sendMatchData(c)
}
func twoonone(c *Client, b bool) {
	resetPlayOption(c)
	c.LiveMatch.M.Settings.TwoOnOne = b
	sendMatchData(c)
}
func oneonone(c *Client, b bool) {
	resetPlayOption(c)
	c.LiveMatch.M.Settings.OneOnOne = b
	sendMatchData(c)
}

func resetPlayOption(c *Client) {
	c.LiveMatch.M.Settings.OneOnOne = false
	c.LiveMatch.M.Settings.TwoOnTwo = false
	c.LiveMatch.M.Settings.TwoOnOne = false
}

func setcolor(c *Client, color string) {
	c.User.Color = color
	sendMatchData(c)
}

func isBet(c *Client, b bool) {
	c.LiveMatch.M.Settings.Bet = b
	sendMatchData(c)
}
func isTournament(c *Client, b bool) {
	c.LiveMatch.M.Settings.Tournament = b
	sendMatchData(c)
}
func tournamentmode(c *Client, b bool) {
	c.LiveMatch.M.Settings.TournamentMode = b
	resetPlayOption(c)
	c.LiveMatch.M.Settings.TwoOnOne = true
	setMaxGoals(c, 3)
	sendMatchData(c)
}
func isDrunk(c *Client, b bool) {
	c.LiveMatch.M.Settings.Drunk = b
	sendMatchData(c)
}
func isPayed(c *Client, b bool) {
	c.LiveMatch.M.Settings.Payed = b
	sendMatchData(c)
}

func freegame(c *Client, b bool) {
	c.LiveMatch.M.Settings.FreeGame = b
	sendMatchData(c)
}

func setReady(c *Client, r bool) {
	c.User.Ready = r
	sendMatchData(c)
}

func setbet(c *Client, bet int64) {
	c.User.Bet = bet
	sendMatchData(c)
}

func setRated(c *Client, rated bool) {
	c.LiveMatch.M.Settings.Rated = rated
	sendMatchData(c)
}

func maxtime(c *Client, time int64) {
	c.LiveMatch.M.Settings.MaxTime = time
	sendMatchData(c)
}

func kickuser(c *Client, id string) {
	for client, _ := range c.LiveMatch.Clients {
		if client.IsUser && client.User.ID == id {
			closeUser(client)
			return
		}
	}
}

// End: Admin M.Settings

// For Admin and Table
func maxgoals(c *Client, maxTime int64) {
	c.LiveMatch.M.Settings.MaxGoals = maxTime
	sendMatchData(c)
}

func removegoal(c *Client, side string) {
	if side == "blue" {
		c.LiveMatch.M.ScoreBlue--
	} else {
		c.LiveMatch.M.ScoreRed--
	}
	sendMatchData(c)
}

func addgoal(c *Client, side string, speed float64) {
	c.LiveMatch.Goals = append(c.LiveMatch.Goals, &models.Goal{
		Side:     side,
		Speed:    speed,
		DateTime: time.Now().Format(time.RFC3339),
	})
	if side == "red" {
		c.LiveMatch.M.ScoreRed++
	} else if side == "blue" {
		c.LiveMatch.M.ScoreBlue++
	}
	if checkGameCompleted(c) {
		c.LiveMatch.M.Started = false
		// Implement reset game function (score, kick players)
		c.LiveMatch.M.ScoreRed = 0
		c.LiveMatch.M.ScoreBlue = 0

		defer func() {
			SendMatch(c.LiveMatch)
		}()
		if c.LiveMatch.M.Settings.TournamentMode {
			rotatePeople(c.LiveMatch.M.Positions)
		}
		sendMatchData(c)
	} else {
		sendMatchData(c)
	}
}

func checkGameCompleted(c *Client) bool {
	m := c.LiveMatch.M
	if m.Settings.MaxGoals > 0 {
		if m.Settings.MaxGoals <= m.ScoreRed || m.Settings.MaxGoals <= m.ScoreBlue {
			return true
		}
	} else {
		// TODO: Implement in GoRoutine where it keeps track of time
		//if m.Settings.MaxTime <= m.ScoreRed ||m.Settings.MaxGoals <= m.ScoreBlue {
		//	return true
		//}
		return false
	}
	return false
}

func setMaxGoals(c *Client, i int64) {
	c.LiveMatch.M.Settings.MaxGoals = i
	c.LiveMatch.M.Settings.MaxTime = 0
}

func rotatePeople(positions *models.MatchPositions) {
	var p1 string
	if positions.RedAttack != "" && positions.RedDefense != "" {
		if positions.BlueAttack != "" {
			p1 = positions.BlueAttack
			positions.BlueAttack = positions.RedAttack
		} else {
			p1 = positions.BlueDefense
			positions.BlueDefense = positions.RedAttack
		}
		positions.RedAttack = positions.RedDefense
		positions.RedDefense = p1
	} else {
		if positions.RedAttack != "" {
			p1 = positions.RedAttack
			positions.RedAttack = positions.BlueAttack
		} else {
			p1 = positions.RedDefense
			positions.RedDefense = positions.BlueAttack
		}
		positions.BlueAttack = positions.BlueDefense
		positions.BlueDefense = p1
	}
}
