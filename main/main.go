package main

import (
	"flag"
	"github.com/iafoosball/livematches-service/handler"
	"log"
	"net/http"
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

func main() {
	flag.Parse()
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("Open for clients on: " + *host + ":" + *port)
	hub := handler.NewHub()
	go hub.Run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handler.ServeWs(hub, w, r)
	})
	err := http.ListenAndServe(*host+":"+*port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
