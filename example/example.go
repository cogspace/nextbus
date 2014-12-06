package main

import (
	"github.com/cogspace/nextbus"
	"github.com/kr/pretty"
	"log"
)

func main() {
	predictions, err := nextbus.GetPredictionsMulti(
		"glendale",
		[]nextbus.Stop{
			{"12", "sanchev"},
			{"12", "gtc_d"},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	pretty.Println(predictions)
}
