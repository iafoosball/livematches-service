package main

import (
	"context"
	pb "github.com/iafoosball/livematches-service/proto"
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"google.golang.org/grpc"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	*DevMode = true
	go main()
	time.Sleep(time.Second * 1)

}

// newUser creates a new user struct
func newUser(tableID string) *pb.User {
	return &pb.User{
		Id:           strconv.Itoa(rand.Int()),
		CurrentTable: tableID,
	}
}

// newClient opens a new client connection to the server
func newClient() pb.LivematchClient {
	conn, err := grpc.Dial(*port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}
	return pb.NewLivematchClient(conn)
}

func TestStream(t *testing.T) {
	client := newClient()
	stream, err := client.Send(context.Background())
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}
	match := &pb.Match{}
	user := newUser("1")

	stream.Send(&pb.Command{User: user})
	match, _ = stream.Recv()
	log.Println(match)

	stream.Send(&pb.Command{AddGoalBlue: true})
	match, _ = stream.Recv()
	log.Println(match)

	time.Sleep(time.Second * 1)
	stream.Send(&pb.Command{AddGoalBlue: true})
	match, _ = stream.Recv()
	log.Println(match)

	time.Sleep(time.Second * 1)
	s := &pb.Settings{Mode: pb.Mode_twoOnOne}
	stream.Send(&pb.Command{Settings: s})

	match, _ = stream.Recv()
	log.Println(match)

	stream.Send(&pb.Command{Join: "id"})

	match, _ = stream.Recv()
	log.Println(match)
	//
	//log.Println(stream.Recv())
}
