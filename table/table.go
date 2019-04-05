package table

import "github.com/iafoosball/livematches-service/proto"

type Table struct {
	Users chan *proto.Livematch_SendServer
	Match proto.Match
}
