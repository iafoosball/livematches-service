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
	scenario string
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

func TestIntegrationScenario1(*testing.T) {
	scenario = "scenario1"
	addr := "0.0.0.0:9003"

	go main()
	time.Sleep(1 * time.Second)
	go table.Start("table1", scenario, addr)
	time.Sleep(1 * time.Second)
	go user.Start("user1", scenario, addr)
	//go user.Start("user2", scenario, addr)

	// Test logic inside here
	for table.Stop != true {

	}

	log.Println("Exit now!")
	os.Exit(3)
}
