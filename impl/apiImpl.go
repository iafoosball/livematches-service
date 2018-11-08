package impl

import (
	"time"
)

func setusername(c *Client, username string) {
	c.user.Username = username
	sendMatchData(c)
}

func setposition(c *Client, position string, side string) {
	resetPosition(c)
	if position == "attack" && side == "blue" {
		if c.liveMatch.Positions.BlueAttack == "" {
			c.liveMatch.Positions.BlueAttack = c.user.ID
		}
	} else if position == "defense" && side == "blue" {
		if c.liveMatch.Positions.BlueDefense == "" {
			c.liveMatch.Positions.BlueDefense = c.user.ID
		}
	} else if position == "attack" && side == "red" {
		if c.liveMatch.Positions.RedAttack == "" {
			c.liveMatch.Positions.RedAttack = c.user.ID
		}
	} else if position == "defense" && side == "red" {
		if c.liveMatch.Positions.RedDefense == "" {
			c.liveMatch.Positions.RedDefense = c.user.ID
		}
	}
	sendMatchData(c)
}

func resetPosition(c *Client) {
	if c.liveMatch.Positions.RedDefense == c.user.ID {
		c.liveMatch.Positions.RedDefense = ""
	} else if c.liveMatch.Positions.RedAttack == c.user.ID {
		c.liveMatch.Positions.RedAttack = ""
	} else if c.liveMatch.Positions.BlueDefense == c.user.ID {
		c.liveMatch.Positions.BlueDefense = ""
	} else if c.liveMatch.Positions.BlueAttack == c.user.ID {
		c.liveMatch.Positions.BlueAttack = ""
	}
}

// startMatch writes everything to the Match object.
// Before users etc. is stored on the livematch
func startmatch(c *Client) {
	c.liveMatch.StartTime = time.Now().Format(time.RFC3339)
	c.liveMatch.Started = true
	sendMatchData(c)
}

func switchpositions(c *Client, b bool) {
	c.liveMatch.Settings.SwitchPosition = b
	sendMatchData(c)
}

func twoontwo(c *Client, b bool) {
	resetPlayOption(c)
	c.liveMatch.Settings.TwoOnOne = b
	sendMatchData(c)
}
func twoonone(c *Client, b bool) {
	resetPlayOption(c)
	c.liveMatch.Settings.TwoOnOne = b
	sendMatchData(c)
}
func oneonone(c *Client, b bool) {
	resetPlayOption(c)
	c.liveMatch.Settings.OneOnOne = b
	sendMatchData(c)
}

func resetPlayOption(c *Client) {
	c.liveMatch.Settings.OneOnOne = false
	c.liveMatch.Settings.TwoOnTwo = false
	c.liveMatch.Settings.TwoOnOne = false
}

func setcolor(c *Client, color string) {
	c.user.Color = color
	sendMatchData(c)
}

func isBet(c *Client, b bool) {
	c.liveMatch.Settings.Bet = b
	sendMatchData(c)
}
func isTournament(c *Client, b bool) {
	c.liveMatch.Settings.Tournament = b
	sendMatchData(c)
}
func isDrunk(c *Client, b bool) {
	c.liveMatch.Settings.Drunk = b
	sendMatchData(c)
}
func isPayed(c *Client, b bool) {
	c.liveMatch.Settings.Payed = b
	sendMatchData(c)
}

func freegame(c *Client, b bool) {
	c.liveMatch.Settings.FreeGame = b
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
	c.liveMatch.Settings.RatedMatch = rated
	sendMatchData(c)
}

func maxtime(c *Client, goals int64) {
	c.liveMatch.Settings.MaxTime = goals
	sendMatchData(c)
}

func maxgoals(c *Client, maxTime int64) {
	c.liveMatch.Settings.MaxGoals = maxTime
	sendMatchData(c)
}

func removegoal(c *Client, side string) {
	if side == "blue" {
		c.liveMatch.ScoreBlue--
	} else {
		c.liveMatch.ScoreRed--
	}
	sendMatchData(c)
}
