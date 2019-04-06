package main

import (
	pb "github.com/iafoosball/livematches-service/proto"
	"log"
)

func selectCommand(c *pb.Command, t *Table, srv *pb.Livematch_SendServer) {
	log.Println("add CMD")
	if c.AddGoalBlue {
		t.Match.ScoreBlue++
	} else if c.AddGoalRed {
		t.Match.ScoreRed++
	}
	t.Data <- t.Match
}
