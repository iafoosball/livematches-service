package impl

import (
	"flag"
	"log"
)

var (
	testHost string
	testPort string
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	flag.StringVar(&testHost, "testhost", "0.0.0.0", "the test host")
	flag.StringVar(&testPort, "testport", "8000", "the port of the matches service where the test should connect")

	flag.Parse()
	MatchesAddr = "http://" + testHost + ":" + testPort + "/"
	log.Println("The test addr is " + MatchesAddr)
}
