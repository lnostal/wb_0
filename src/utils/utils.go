package utils

import (
	"L0/src/db/model"
	"encoding/json"
	"errors"
)

func ValidateJSON(data []byte) error {
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

func ConvertToModel(data []byte, t any) any {

	err := json.Unmarshal(data, &t)

	if err != nil {
		return err
	}
	return t
}
