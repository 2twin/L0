package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/nats-io/stan.go"

	_ "github.com/lib/pq"
)

// Order represents the structure of order data
type Order struct {
	OrderUID          string    `json:"order_uid"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Delivery          Delivery  `json:"delivery"`
	Payment           Payment   `json:"payment"`
	Items             []Item    `json:"items"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerID        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmID              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}

// Delivery represents delivery information
type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

// Payment represents payment information
type Payment struct {
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int64  `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

// Item represents an item in the order
type Item struct {
	ChrtID      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmID        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

type NatsConfig struct {
	ClusterID string
	ClientID  string
	Subject   string
}

func NewNatsConfig() (*NatsConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("couldn't find .env file")
	}

	clusterId := os.Getenv("NATS_CLUSTER_ID")
	if clusterId == "" {
		return nil, errors.New("NATS_CLUSTER_ID environment variable is not set")
	}

	clientId := os.Getenv("NATS_CLIENT_ID")
	if clientId == "" {
		return nil, errors.New("NATS_CLIENT_ID environment variable is not set")
	}

	subject := os.Getenv("NATS_SUBJECT")
	if subject == "" {
		return nil, errors.New("NATS_SUBJECT environment variable is not set")
	}

	return &NatsConfig{
		ClusterID: clusterId,
		ClientID:  clientId,
		Subject:   subject,
	}, nil
}

func main() {
	db, err := sql.Open("postgres", "host=postgres port=5432 user=postgres password=postgres dbname=l0 sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	natsConf, err := NewNatsConfig()
	if err != nil {
		log.Fatal(err)
	}

	sc, err := stan.Connect(natsConf.ClusterID, natsConf.ClientID)
	if err != nil {
		log.Fatal(err)
	}

	sub, err := sc.Subscribe(natsConf.Subject, func(msg *stan.Msg) {
		var order Order
		if err := json.Unmarshal(msg.Data, &order); err != nil {
			log.Printf("Error decoding JSON: %v", err)
		}
	})

	defer sub.Unsubscribe()
}
