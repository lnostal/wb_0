package db

import (
	"L0/src/db/model"
	"L0/src/utils"
	"encoding/json"
)

func PutData(data []byte) {
	var order model.Order
	utils.ConvertToModel(data, &order)

	base := Connect()

	_, err := base.Exec("insert into orders(order_uid, track_number, entry, delivery,payment,items,locale,internal_signature,customer_id,delivery_service, shardkey, sm_id,date_created,oof_hard) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14);",
		order.OrderUID, order.TrackNumber, order.Entry, order.Delivery, order.Payment, order.Items, order.Locale, order.InternalSignature, order.CustomerID, order.DeliveryService, order.Shardkey, order.SmID, order.DateCreated, order.OofShard)
	if err != nil {
		defer base.Close()
		return
	}

	defer base.Close()
}

func GetData() ([][]byte, error) {
	var dt [][]byte

	base := Connect()
	rows, err := base.Query("select * from orders;")

	if err != nil {
		defer base.Close()
		return dt, err
	}

	for rows.Next() {
		var order model.Order
		err := rows.Scan(&order.OrderUID, &order.TrackNumber, &order.Entry, &order.Delivery, &order.Payment, &order.Items, &order.Locale, &order.InternalSignature, &order.CustomerID, &order.DeliveryService, &order.Shardkey, &order.SmID, &order.DateCreated, &order.OofShard)
		if err != nil {
			return dt, err
		} else {
			data, _ := json.Marshal(order)
			dt = append(dt, data)
		}
	}
	defer base.Close()
	return dt, nil
}
