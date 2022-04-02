package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/leoadamek/iracing-data-go"
)

func main() {
	sessionId := flag.Uint64("s", 0, "Session ID")
	flag.Parse()

	client := iracing.New(iracing.StaticCredentials(os.Getenv("IRACING_USERNAME"), os.Getenv("IRACING_PASSWORD")))

	if err := client.Login(); err != nil {
		log.Fatalln("Unable to log in: ", err)
	}

	laps, err := client.GetLaps(*sessionId, 0)

	if err != nil {
		log.Fatalln("Unable to get laps: ", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 4, 2, 2, ' ', 0)

	fmt.Println("Lap #\tCar #\tPosition\tTime (Raw)\tTime\n")

	for _, lap := range laps {
		fmt.Fprintf(
			tw, "%3d\t%3s\t%2d\t%d\t%s\n",
			lap.LapNumber, lap.CarNumber,
			lap.LapPosition, lap.LapTime,
			lap.Time(),
		)
	}

	tw.Flush()
}
