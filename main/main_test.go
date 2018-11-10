package main

import (
	"github.com/iafoosball/livematches-service/testclients/table"
	"github.com/iafoosball/livematches-service/testclients/user"
	"log"
	"os"
	"os/signal"
	"testing"
	"time"
)

// used by all test classes in package matches
var (
	scenario string
	addr     = "0.0.0.0:9003"
	//addr     = "iafoosball.aau.dk:9003"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	go main()

	//f, err := os.OpenFile("livematches.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	//if err != nil {
	//	log.Fatalf("error opening file: %v", err)
	//}
	//log.SetOutput(f)
	//log.Print("hallo")
}

func TetRunTestClient(*testing.T) {
	end := make(chan bool)
	go user.Start("user1", "table1", scenario, addr, end)
}

func TestUser(*testing.T) {
	log.Println("TestUser")
	scenario = "testUser"
	end := make(chan bool)
	go table.Start("table1", scenario, addr, end)
	time.Sleep(1 * time.Second)
	go user.Start("user1", "table1", scenario, addr, end)
	for {

		select {
		case _ = <-end:
			return
		}
	}
}

func TestAdmin(t *testing.T) {
	log.Println("testAdmin")
	scenario = "testAdmin"
	end := make(chan bool)
	go table.Start("table1", scenario, addr, end)
	time.Sleep(1 * time.Second)
	go user.Start("user1", "table1", scenario, addr, end)
	time.Sleep(1 * time.Second)
	go user.Start("user2", "table1", scenario, addr, end)
	for {
		select {
		case _ = <-end:
			return
		}
	}
}

// exit the program if Ctrl-C is pressed
func exit() {
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan struct{})
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		<-signalChan
		os.Exit(0)
		close(cleanupDone)
	}()
	<-cleanupDone

}
