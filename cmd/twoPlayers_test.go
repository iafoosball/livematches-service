package main

import (
	"context"
	"log"
	"testing"
	"time"

	pb "github.com/iafoosball/livematches-service/proto"
)

func TestTwoPlayers(t *testing.T) {
	//wg := sync.WaitGroup{}
	//wg.Add(1)
	go doStuff()
	time.Sleep(time.Millisecond * 1000)
	client := newClient()
	stream, _ := client.Send(context.Background())
	stream.Send(&pb.Command{User: newUser("1")})
	log.Println(stream.Recv())
	log.Println(stream.Recv())
	log.Println(stream.Recv())
	//log.Println(stream.Recv())
	//log.Println(stream.Recv())
	//
	//log.Println(stream.Recv())
	//wg.Wait()
}

func doStuff() {
	client := newClient()
	stream, err := client.Send(context.Background())
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}
	match := &pb.Match{}
	user := newUser("1")

	stream.Send(&pb.Command{User: user})

	match, _ = stream.Recv()
	time.Sleep(time.Second * 3)
	log.Println(match)
	log.Println(user.Id)

	stream.Send(&pb.Command{AddGoalBlue: true})
	match, _ = stream.Recv()
	log.Println(match)
	log.Println(user.Id)

	time.Sleep(time.Second * 3)
	stream.Send(&pb.Command{AddGoalBlue: true})
	match, _ = stream.Recv()
	log.Println(match)
	log.Println(user.Id)

	time.Sleep(time.Second * 3)
	s := &pb.Settings{Mode: pb.Mode_twoOnOne}
	stream.Send(&pb.Command{Settings: s})

	match, _ = stream.Recv()
	log.Println(match)
	log.Println(user.Id)

	stream.Send(&pb.Command{Join: "id"})

	match, _ = stream.Recv()
	log.Println(match)
}
