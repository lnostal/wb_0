package nats

import (
	"L0/src/cache"
	"L0/src/utils"
	stan "github.com/nats-io/stan.go"
	"sync"
)

func Connect() {
	sc, _ := stan.Connect("test-cluster", "sub")

	receiveMessages(sc)

	w := sync.WaitGroup{}
	w.Add(1)
	w.Wait()

	sc.Close()
}

func receiveMessages(sc stan.Conn) {
	_, err := sc.Subscribe("wb", func(m *stan.Msg) {
		err := validateData(m.Data)
		if err == nil {
			cache.Add(m.Data)
			//db.PutData(m.Data)
		}
	})

	if err != nil {
		panic(err)
	}
}

func validateData(data []byte) error {
	return utils.ValidateJSON(data)
}
