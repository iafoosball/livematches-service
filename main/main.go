package main

import (
	"flag"
	"github.com/go-openapi/swag"
	"github.com/iafoosball/livematches-service/impl"
	"github.com/iafoosball/livematches-service/models"
	"log"
	"net/http"
	"strings"
)

var (
	host        = flag.String("host", "0.0.0.0", "the host to listen for connections")
	port        = flag.String("port", "8003", "the port to listen for new clients")
	matchesHost = flag.String("matchesHost", "0.0.0.0", "the host for sending match data to")
	matchesPort = flag.String("matchesPort", "8000", "the host port for sending match data to")
	DevMode     = flag.Bool("dev", false, "enable if used as Developer. Serves over http.")

	err error
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
		var m []*models.Match
		for _, n := range impl.LiveMatches {
			m = append(m, n.M)
		}
		b, err := swag.WriteJSON(m)
		if err != nil {
			log.Println(err)
			return
		}
		w.Write(b)
	}
}

func main() {
	// init everything
	flag.Parse()
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("V1: Open for clients on: " + *host + ":" + *port)
	impl.MatchesAddr = "http://" + *matchesHost + ":" + *matchesPort
	log.Println("Database is on " + impl.MatchesAddr)
	hub := impl.NewHub()
	go hub.Run()
	mux := http.NewServeMux()
	mux.HandleFunc("/", serveHome)
	mux.HandleFunc("/matches", listMatches)
	mux.HandleFunc("/tables/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("connected")
		s := strings.Split(r.URL.Path, "/")
		// 2 and 3 are hardedcoded so it fails, if id is not specified.
		impl.ServeWs(hub, w, r, false, s[2], "")
	})
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		s := strings.Split(r.URL.Path, "/")
		impl.ServeWs(hub, w, r, true, s[2], s[3])
	})
	if *DevMode {
		http.ListenAndServe(":"+*port, nil)
	} else {
		go http.ListenAndServe(":9010", http.HandlerFunc(redirect))
		log.Println("http server running")
		err = http.ListenAndServeTLS(":"+*port, "/certs/cert.pem", "/certs/privkey.pem", mux)
	}
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func redirect(w http.ResponseWriter, req *http.Request) {
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, req, target,
		http.StatusTemporaryRedirect)
}
