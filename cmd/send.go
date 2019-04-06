package main

import (
	"io"
	"log"

	pb "github.com/iafoosball/livematches-service/proto"
)

// Send implements SendServer.
// The server waits for input from one client of a table and then broadcasts to all on same table.
func (s server) Send(srv pb.Livematch_SendServer) error {
	log.Println("start new server")
	ctx := srv.Context()
	tableID := ""

	for {
		// exit if context is done or continue
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		cmd, err := srv.Recv()
		if cmd.GetUser() != nil {
			tableID = cmd.User.CurrentTable
			t := Tables[tableID]
			t.Connections = append(t.Connections, srv)
			t.Match.Users = append(t.Match.Users, cmd.User)
			t.Data <- t.Match
		} else if tableID != "" {
			selectCommand(cmd, Tables[tableID], &srv)
		}

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
