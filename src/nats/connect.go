package nats

import (
	"L0/src/db/model"
	"encoding/json"
	"errors"
	"fmt"
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
		if err != nil {
			fmt.Println("это в бд не пойдет азазаз")
		}
	}, stan.StartWithLastReceived())

	if err != nil {
		panic(err)
	}
}

func validateData(data []byte) error {
	if json.Valid(data) {
		var order model.Order
		err := json.Unmarshal(data, &order)

		if err != nil {
			return err
		}

		var delivery model.Delivery
		err = json.Unmarshal(order.Delivery, &delivery)

		if err != nil {
			return err
		}

		var payment model.Payment
		err = json.Unmarshal(order.Payment, &payment)

		if err != nil {
			return err
		}

		var items []model.Item
		err = json.Unmarshal(order.Items, &items)
		if err != nil {
			return err
		}

	} else {
		return errors.New("not json")
	}

	return nil
}
