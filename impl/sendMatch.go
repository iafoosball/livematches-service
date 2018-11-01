package impl

import (
	"bytes"
	"encoding/json"
	"github.com/iafoosball/livematches-service/models"
	"io/ioutil"
	"net/http"
	"strings"
)

func SendMatch(liveMatch LiveMatch) {
	js, _ := json.Marshal(liveMatch)
	resp, err := http.Post("http://localhost:8000/"+"matches/", "application/json", bytes.NewReader(js))
	handleErr(err)
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
		resp, err = http.Post("http://localhost:8000/"+"goals/", "application/json", bytes.NewReader(js))
	}

}
