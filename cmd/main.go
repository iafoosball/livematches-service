package main

import (
	"flag"
	pb "github.com/iafoosball/livematches-service/proto"
	"github.com/iafoosball/livematches-service/table"
	"io"
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

	err    error
	tables []table.Table
)

type server struct{}

// SayHello implements helloworld.GreeterServer
func (s server) Send(srv pb.Livematch_SendServer) error {
	log.Println("start new server")
	ctx := srv.Context()
	match := &pb.Match{Settings: &pb.Settings{}}

	for {

		// exit if context is done
		// or continue
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		cmd, err := srv.Recv()
		log.Println(cmd)
		if cmd.AddGoalBlue == true {
			match.ScoreBlue++
		} else if cmd.Settings.Mode == pb.Mode_twoOnOne {
			match.Settings.Mode = cmd.Settings.Mode
		}
		srv.Send(match)

		// receive data from stream
		if err == io.EOF {
			// return will close stream from server side
			log.Println("exit")
			return nil
		}
		if err != nil {
			log.Printf("receive error %v", err)
			continue
		}
	}
}

func main() {
	flag.Parse()
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println(*port)
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
