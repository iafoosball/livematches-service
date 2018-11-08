package main

import (
	"github.com/iafoosball/livematches-service/testclients/table"
	"github.com/iafoosball/livematches-service/testclients/user"
	"log"
	"os"
	"os/signal"
	"sync"
	"testing"
	"time"
)

// used by all test classes in package matches
var (
	scenario string
	addr     = "0.0.0.0:9003"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	//f, err := os.OpenFile("livematches.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	//if err != nil {
	//	log.Fatalf("error opening file: %v", err)
	//}
	//log.SetOutput(f)
	//log.Print("hallo")
}

func TestRunTestClient(*testing.T) {
	end := make(chan string)
	go user.Start("user1", "table2", scenario, addr, end)
}

func TestIntegrationScenario1(*testing.T) {
	scenario = "scenario1"
	end := make(chan string)
	var wg sync.WaitGroup

	log.Print("123")
	wg.Add(3)
	go main()
	time.Sleep(1 * time.Second)
	go table.Start("table2", scenario, addr, end)
	time.Sleep(2 * time.Second)
	go user.Start("user1", "table2", scenario, addr, end)
	//go user.Start("user2", scenario, addr)

	exit()
	//wg.Wait()

	// Test logic inside here
	for table.Stop != true {

	}

	log.Println("Exit now!")
	os.Exit(3)
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
