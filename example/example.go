package main

import (
	"fmt"
	"log"

	"github.com/cogspace/nextbus"
)

type tr struct {
	Stop    []struct{ Content, Tag, EpochTime string }
	BlockId string
}

func f(tr tr) {
	fmt.Printf("%#v\n", tr)
}

func main() {
	schedules, err := nextbus.GetSchedules("glendale", "12")
	if err != nil {
		log.Fatal(err)
	}

	tr := schedules[0].Tr[0]
	f(tr)
}
