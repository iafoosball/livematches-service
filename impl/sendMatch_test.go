package impl

import (
	"github.com/iafoosball/livematches-service/models"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestEndMatch(t *testing.T) {
	match := &LiveMatch{
		M: &models.Match{
			Users: []*models.MatchUsersItems0{
				&models.MatchUsersItems0{
					Admin:    true,
					Bet:      123,
					Color:    "gree",
					ID:       "3333",
					Position: "attack",
					Username: "kickAss",
				}, &models.MatchUsersItems0{
					Admin:    false,
					Bet:      2,
					Color:    "blue",
					ID:       "4444",
					Position: "defense",
					Username: "kickme",
				},
			},
			ScoreBlue: 2,
			ScoreRed:  10,
			StartTime: time.Now().Format(time.RFC3339),
			EndTime:   time.Now().Format(time.RFC3339),
			TableID:   "123",
			Winner:    "red",
			Settings: &models.MatchSettings{
				Bet:             true,
				Drunk:           false,
				FreeGame:        true,
				MaxGoals:        10,
				MaxTime:         1000,
				OneOnOne:        true,
				Payed:           false,
				Rated:           true,
				SwitchPositions: false,
				Tournament:      false,
				TwoOnOne:        false,
				TwoOnTwo:        false,
			},
		},
		Goals: []*models.Goal{
			&models.Goal{
				DateTime: time.Now().Format(time.RFC3339),
				Position: false,
				Side:     "attack",
			},
			&models.Goal{
				DateTime: time.Now().Format(time.RFC3339),
				Position: false,
				Side:     "attack",
			},
			&models.Goal{
				DateTime: time.Now().Format(time.RFC3339),
				Position: false,
				Side:     "attack",
			},
		},
	}
	SendMatch(match)
	query := MatchesAddr + "/matches/?filter=settings.maxTime==1000"
	resp, err := http.Get(query)
	handleTesterr(t, err)
	if http.StatusOK != resp.StatusCode {
		t.Error("wrong http code")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if !strings.Contains(string(body), "\"maxTime\":1000") {
		log.Fatal("\"maxTime\":1000 not in string!")
	}
}

func handleTesterr(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
