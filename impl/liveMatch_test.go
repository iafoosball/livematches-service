package impl

import (
	"bytes"
	"encoding/json"
	"github.com/iafoosball/livematches-service/models"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestEndMatch(t *testing.T) {
	match := LiveMatch{
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
				MaxTime:         10,
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
	js, _ := json.Marshal(match.M)
	resp, err := http.Post("http://localhost:8000/"+"matches/", "application/json", bytes.NewReader(js))
	handleErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	handleErr(err)
	m := models.DocumentMeta{}
	err = json.NewDecoder(strings.NewReader(string(body))).Decode(&m)
	handleErr(err)
	log.Println(m)
	goals := match.Goals
	for _, g := range goals {
		g.From = m.ID
		g.To = m.ID
		js, _ = json.Marshal(g)
		resp, err = http.Post("http://localhost:8000/"+"goals/", "application/json", bytes.NewReader(js))
	}
}
