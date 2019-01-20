package impl

import (
	"bytes"
	"encoding/json"
	"github.com/iafoosball/livematches-service/models"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	MatchesAddr string
)

func SendMatch(Match models.Match, goals []*models.Goal) {
	js, err := json.Marshal(Match)
	handleErr(err)
	resp, err := http.Post(MatchesAddr+"/matches/", "application/json", bytes.NewReader(js))
	handleErr(err)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	handleErr(err)
	m := models.DocumentMeta{}
	err = json.NewDecoder(strings.NewReader(string(body))).Decode(&m)
	handleErr(err)
	for _, g := range goals {
		g.From = m.ID
		g.To = m.ID
		js, _ = json.Marshal(g)
		resp, err = http.Post(MatchesAddr+"/goals/", "application/json", bytes.NewReader(js))
	}
}
