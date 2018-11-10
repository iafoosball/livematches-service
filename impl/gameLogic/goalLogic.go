package gameLogic

import (
	"github.com/iafoosball/livematches-service/impl"
	"github.com/iafoosball/livematches-service/models"
)

func AddGoalLogic(lm *impl.LiveMatch) {
	m := lm.M
	if m.Settings.TournamentMode {
		if m.ScoreBlue > 2 {
			impl.SendMatch(lm)
			resetScore(m)
		}
	}
}

func resetScore(m *models.Match) {
	m.ScoreBlue = 0
	m.ScoreRed = 0
}
