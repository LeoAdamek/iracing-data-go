package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/leoadamek/iracing-data-go"
)

func main() {
	sessionId := flag.Uint64("s", 0, "Session ID")
	flag.Parse()

	client := iracing.New(iracing.StaticCredentials(os.Getenv("IRACING_USERNAME"), os.Getenv("IRACING_PASSWORD")))

	client.Verbose = true

	if err := client.Login(); err != nil {
		log.Fatalln("Unable to log in: ", err)
	}

	session, err := client.GetSession(*sessionId)

	if err != nil {
		log.Fatalln("Unable to get session: ", err)
	}

	fmt.Printf("Session: %#+v\n", session)
}
