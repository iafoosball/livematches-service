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
	scenario string
)

func TestIntegrationScenario1(*testing.T) {
	scenario = "scenario1"
	log.SetFlags(log.Ltime | log.Lshortfile)
	go main()
	time.Sleep(1 * time.Second)
	go table.Start("table-1", scenario)
	time.Sleep(1 * time.Second)
	go user.Start("user-1", scenario)
	go user.Start("user-2", scenario)

	time.Sleep(5 * time.Second)
	log.Println("Exit now!")
	os.Exit(3)
}
