package impl

import (
	"github.com/iafoosball/livematches-service/models"
	"log"
	"time"
)

func SendMatch(liveMatch LiveMatch) {
	match := models.Match{
		Bet:        liveMatch.Bet,
		Completed:  liveMatch.Completed,
		EndTime:    time.Now().Format(time.RFC3339),
		FreeGame:   liveMatch.FreeGame,
		MaxTime:    liveMatch.MaxTime,
		MaxGoals:   &liveMatch.MaxGoals,
		OneOnOne:   liveMatch.OneOnOne,
		Payed:      liveMatch.Payed,
		Positions:  liveMatch.Positions,
		RatedMatch: liveMatch.RatedMatch,
		ScoreBlue:  liveMatch.ScoreBlue,
	}
	log.Println(match)

}

/*


// payed
Payed bool `json:"payed,omitempty"`

// positions
Positions *MatchPositions `json:"positions,omitempty"`

// A match can be rated, ie a ranked game with points, or without.
RatedMatch bool `json:"ratedMatch,omitempty"`

// score blue
ScoreBlue int64 `json:"scoreBlue,omitempty"`

// score red
ScoreRed int64 `json:"scoreRed,omitempty"`

// the datetime when the game ends
StartTime string `json:"startTime,omitempty"`

// started
Started bool `json:"started,omitempty"`

// Switch the position after 50% of the goal (time or goals) is reached.
SwitchPosition bool `json:"switchPosition,omitempty"`

// the id of table
TableID string `json:"tableID,omitempty"`

// tournament
Tournament bool `json:"tournament,omitempty"`

// two on one
TwoOnOne bool `json:"twoOnOne,omitempty"`

// two on two
TwoOnTwo bool `json:"twoOnTwo,omitempty"`

// users
Users *MatchUsers `json:"users,omitempty"`

// Can be either "red" or "blue"
Winner string `json:"winner,omitempty"`
*/
