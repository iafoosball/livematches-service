package main

import (
	"testing"
)

// used by all test classes in package matches
var (
	testHost string
	testPort int
	testUrl  string
)

func init() {
	//log.SetFlags(log.Ltime | log.Lshortfile)
	//flag.StringVar(&testHost, "testHost", "0.0.0.0", "the test host")
	//testPort = *flag.Int("testPort", 8000, "the port of the matches service where the test should connect")
	//flag.Parse()
	//testUrl = "http://" + testHost + ":" + strconv.Itoa(testPort) + "/"
	//log.Println(testUrl)
}

func TestWS(*testing.T) {
	//go client.Client()
	//client.Client()

}
