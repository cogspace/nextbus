package main

import (
	"github.com/cogspace/nextbus"
	"log"
)

type tr struct {
	Stop    []struct{ Content, Tag, EpochTime string }
	BlockId string
}

func f(tr tr) {
}

func main() {
	schedules, err := nextbus.GetSchedules("glendale", "12")
	if err != nil {
		log.Fatal(err)
	}

	tr := schedules[0].Tr[0]
	f(tr)
}
