package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
)

var (
	app = kingpin.New("intruder-alert", "Detect bad things happening in our pg2pubsub stream").Version("0.0.0")
)

func main() {
	if _, err := app.Parse(os.Args[1:]); err != nil {
		kingpin.Fatalf("%s, try --help", err)
	}

	fmt.Println("intruder detected!")
}
