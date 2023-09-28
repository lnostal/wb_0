package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"os"
	"sync"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: <channel>")
		return
	}

	channel := os.Args[1]
	sc, _ := stan.Connect("test-cluster", "sub")
	defer sc.Close()

	_, err := sc.Subscribe(channel, func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	}, stan.DeliverAllAvailable())

	if err != nil {
		return
	}

	w := sync.WaitGroup{}
	w.Add(1)
	w.Wait()
}
