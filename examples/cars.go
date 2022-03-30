package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/leoadamek/iracing-data-go"
)

func main() {
	client := iracing.New(iracing.StaticCredentials(os.Getenv("IRACING_USERNAME"), os.Getenv("IRACING_PASSWORD")))

	client.Verbose = true

	if err := client.Login(); err != nil {
		log.Fatalln("Unable to log in: ", err)
	}

	cars, err := client.GetCars()

	if err != nil {
		log.Fatalln("Unable to list cars: ", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 4, 2, 2, ' ', 0)

	fmt.Fprintln(tw, "#\tID\tName\tWeight\tPower\tCreated")

	for i, car := range cars {
		fmt.Fprintf(tw, "%2d\t%4d\t%s\t%d\t%d\t%s\n", i, car.ID, car.Name, car.Weight, car.Power, car.Created.Format("2006-01-02"))
	}

	tw.Flush()
}
