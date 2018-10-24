package main

import (
	"encoding/json"
	"flag"
	"github.com/iafoosball/livematches-service/handler"
	"log"
	"net/http"
	"strings"
)

var (
	host = flag.String("host", "0.0.0.0", "the host to listen for connections")
	port = flag.String("port", "9003", "the port to listen for new clients")
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func listMatches(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		for _, m := range handler.LiveMatches {
			log.Printf("%+v\n", m)
		}
		b, err := json.Marshal(handler.LiveMatches)
		if err != nil {
			log.Println(err)
			return
		}
		w.Write(b)
	}
}

func main() {
	flag.Parse()
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("Open for clients on: " + *host + ":" + *port)
	hub := handler.NewHub()
	go hub.Run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/matches", listMatches)
	// if request goes to table, bool isUser is set to false
	http.HandleFunc("/tables/", func(w http.ResponseWriter, r *http.Request) {
		s := strings.Split(r.URL.Path, "/")
		// 3 is hardedcoded so it fails, if id is not specified.
		handler.ServeWs(hub, w, r, false, s[2], "")
	})
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		s := strings.Split(r.URL.Path, "/")
		handler.ServeWs(hub, w, r, true, s[2], s[3])
	})
	err := http.ListenAndServe(*host+":"+*port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
