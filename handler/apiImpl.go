package handler

func setusername(c *Client, username string) {
	c.user.Username = username
	sendMatchData(c)
}

func setposition(c *Client, position string, side string) {
	if position == "attack" && side == "blue" {
		if c.liveMatch.Positions.BlueAttach == "" {
			c.liveMatch.Positions.BlueAttach = c.user.ID
			sendMatchData(c)
		}
	} else if position == "defense" && side == "blue" {
		if c.liveMatch.Positions.BlueDefense == "" {
			c.liveMatch.Positions.BlueDefense = c.user.ID
			sendMatchData(c)
		}
	} else if position == "attack" && side == "red" {
		if c.liveMatch.Positions.RedAttack == "" {
			c.liveMatch.Positions.RedAttack = c.user.ID
			sendMatchData(c)
		}
	} else if position == "defense" && side == "red" {
		if c.liveMatch.Positions.RedDefense == "" {
			c.liveMatch.Positions.RedDefense = c.user.ID
			sendMatchData(c)
		}
	}
}

// startMatch writes everything to the Match object.
// Before users etc. is stored on the livematch
func startmatch(c *Client) {
	c.liveMatch.Started = true
	sendMatchData(c)
}

func switchpositions(c *Client, b bool) {
	c.liveMatch.SwitchPosition = b
	sendMatchData(c)
}

func twoontwo(c *Client, b bool) {
	c.liveMatch.TwoOnOne = b
	sendMatchData(c)
}
func twoonone(c *Client, b bool) {
	c.liveMatch.TwoOnOne = b
	sendMatchData(c)
}
func oneonone(c *Client, b bool) {
	c.liveMatch.OneOnOne = b
	sendMatchData(c)
}

func isBet(c *Client, b bool) {
	c.liveMatch.Bet = b
	sendMatchData(c)
}
func isTournament(c *Client, b bool) {
	c.liveMatch.Tournament = b
	sendMatchData(c)
}
func isDrunk(c *Client, b bool) {
	c.liveMatch.Drunk = b
	sendMatchData(c)
}
func isPayed(c *Client, b bool) {
	c.liveMatch.Payed = b
	sendMatchData(c)
}

func freegame(c *Client, b bool) {
	c.liveMatch.FreeGame = b
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
	c.liveMatch.RatedMatch = rated
	sendMatchData(c)
}

func maxtime(c *Client, goals int64) {
	c.liveMatch.MaxTime = goals
	sendMatchData(c)
}

func maxgoals(c *Client, maxTime int64) {
	c.liveMatch.MaxGoals = maxTime
	sendMatchData(c)
}

func removegoal(c *Client) {
	//c.liveMatch.Goals[] = c.liveMatch.Goals[len(c.liveMatch.Goals)-1]
	c.liveMatch.Goals = c.liveMatch.Goals[:len(c.liveMatch.Goals)-1]
	sendMatchData(c)
}
