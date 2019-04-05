package main

import (
	"context"
	pb "github.com/iafoosball/livematches-service/proto"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	*DevMode = true
	go main()
	time.Sleep(time.Second * 1)

	// dail server
	//conn, err := grpc.Dial(*port, grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf("can not connect with server %v", err)
	//}
	//
	//// create stream
	//client := pb.NewLivematchesClient(conn)
	//stream, err := client.Send(context.Background())
	//if err != nil {
	//	log.Fatalf("openn stream error %v", err)
	//}
	//
	//log.Println(stream.Recv())

	//ctx := stream.Context()
	//done := make(chan bool)

	//ctx.

}

func TestStream(t *testing.T) {
	// dail server
	conn, err := grpc.Dial(*port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	// create stream
	client := pb.NewLivematchClient(conn)
	stream, err := client.Send(context.Background())
	if err != nil {
		log.Fatalf("openn stream error %v", err)
	}
	match := &pb.Match{}

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
	//
	//log.Println(stream.Recv())
}
