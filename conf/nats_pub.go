package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"os"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: <channel> <message>")
		return
	}
	channel := os.Args[1]
	message := os.Args[2]
	sc, _ := stan.Connect("test-cluster", "pub")

	err := sc.Publish(channel, []byte(message))
	if err != nil {
		return
	}
}
