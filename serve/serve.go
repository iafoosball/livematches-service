package serve

import (
	pb "github.com/iafoosball/livematches-service/proto"
)

func Serve(c pb.Command) {
	if c.AddGoalBlue {

	} else if c.AddGoalRed {

	} else if c.Mode == pb.Mode_oneOnOne {

	} else if c.Mode == pb.Mode_twoOnOne {

	} else if c.Mode == pb.Mode_twoOneTwo {

	} else if c.Mode == pb.Mode_tournamentModeTwoRed {

	} else if c.Mode == pb.Mode_tournamentModeTwoBlue {

	}
}
