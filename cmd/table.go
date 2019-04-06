package main

import (
	"github.com/iafoosball/livematches-service/proto"
	pb "github.com/iafoosball/livematches-service/proto"
	"log"
)

var (
	Tables = make(map[string]*Table)
)

type Table struct {
	Data        chan *pb.Match
	Connections []pb.Livematch_SendServer
	Match       *proto.Match
}

// AddTables should call table service for current tables to list. Atm function is mocked
func AddTablesMock() {
	Tables["1"] = &Table{
		Data:        make(chan *pb.Match),
		Connections: []pb.Livematch_SendServer{},
		Match:       &pb.Match{Settings: &pb.Settings{}, Users: []*pb.User{}}}
	go Tables["1"].broadcast()
}

func GetTables() {
}

// broadcast listens on channel for change and then sends to all clients
func (t *Table) broadcast() {
	log.Println("started broadcast channel")
	for {
		select {

		case <-t.Data:
			log.Println("braodcasting")
			for _, c := range t.Connections {
				c.Send(t.Match)
			}
		default:
		}
	}
}
