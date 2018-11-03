package main

import (
	"github.com/iafoosball/livematches-service/testclients/table"
	"github.com/iafoosball/livematches-service/testclients/user"
	"log"
	"os"
	"testing"
	"time"
)

// used by all test classes in package matches
var (
	testHost string
	testPort int
	testUrl  string
)

func TestIntegration(*testing.T) {
	log.SetFlags(log.Ltime | log.Lshortfile)
	go main()
	time.Sleep(2 * time.Second)
	go table.Start()
	time.Sleep(2 * time.Second)
	go user.Start()

	time.Sleep(5 * time.Second)
	log.Println("Exit now!")
	os.Exit(3)
}
