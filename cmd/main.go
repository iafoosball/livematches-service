package main

import (
	"flag"
	pb "github.com/iafoosball/livematches-service/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	host        = flag.String("host", "0.0.0.0", "the host to listen for connections")
	port        = flag.String("port", ":8003", "the port to listen for new clients")
	matchesHost = flag.String("matchesHost", "0.0.0.0", "the host for sending match data to")
	matchesPort = flag.String("matchesPort", "8000", "the host port for sending match data to")
	DevMode     = flag.Bool("dev", false, "enable if used as Developer. Serves over http.")

	err error
)

type server struct{}

// main first parses flags, adds/mocks tables then listen on specified port.
func main() {
	flag.Parse()
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println(*port)
	AddTablesMock()
	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLivematchServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
