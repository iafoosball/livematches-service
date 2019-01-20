package impl

import (
	"bytes"
	"encoding/json"
	"github.com/iafoosball/livematches-service/models"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	MatchesAddr string
)

func SendMatch(liveMatch *LiveMatch) {
	js, err := json.Marshal(*liveMatch.M)
	handleErr(err)
	//MatchesAddr = "http://iafoosball.me:9000"
	resp, err := http.Post(MatchesAddr+"/matches/", "application/json", bytes.NewReader(js))
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	handleErr(err)
	m := models.DocumentMeta{}
	err = json.NewDecoder(strings.NewReader(string(body))).Decode(&m)
	handleErr(err)
	goals := liveMatch.Goals
	for _, g := range goals {
		g.From = m.ID
		g.To = m.ID
		js, _ = json.Marshal(g)
		resp, err = http.Post(MatchesAddr+"/goals/", "application/json", bytes.NewReader(js))
	}
}
