package model

import (
	"time"

	"github.com/google/uuid"
)

func GenerageOrder() *Order {
	orderUid := uuid.New().String()
	trackNumber := uuid.New().String()

	delivery := Delivery{
		Name:    "Test Testov",
		Phone:   "+9720000000",
		Zip:     "2639809",
		City:    "Kiryat Mozkin",
		Address: "Ploshad Mira 15",
		Region:  "Kraiot",
		Email:   "test@gmail.com",
	}

	payment := Payment{
		Transaction: orderUid,
		RequestID: "",
		Currency: "USD",
		Provider: "wbpay",
		Amount: 1817,
		PaymentDt: 1637907727,
		Bank: "alpha",
		DeliveryCost: 1500,
		GoodsTotal: 317,
		CustomFee: 0,
	}

	item1 := Item{
		ChrtID: 9934930,
		TrackNumber: trackNumber,
		Price: 453,
		Rid: "ab4219087a764ae0btest",
		Name: "Mascaras",
		Sale: 30,
		Size: "0",
		TotalPrice: 317,
		NmID: 2389212,
		Brand: "Vivienne Sabo",
		Status: 202,
	}

	item2 := Item{
		ChrtID: 12313,
		TrackNumber: trackNumber,
		Price: 231,
		Rid: "afaef13rfaesfaftest",
		Name: "Name",
		Sale: 10,
		Size: "02",
		TotalPrice: 266,
		NmID: 213123,
		Brand: "Brand",
		Status: 201,
	}

	return &Order{
		OrderUID: orderUid,
		TrackNumber: trackNumber,
		Entry: "WBIL",
		Delivery: delivery,
		Payment: payment,
		Items: []Item{item1, item2},
		Locale: "en",
		InternalSignature: "",
		CustomerID: "test",
		DeliveryService: "meest",
		Shardkey: "9",
		SmID: 99,
		DateCreated: time.Now(),
		OofShard: "1",
	}
}