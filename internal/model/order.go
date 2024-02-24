package model

import "time"

// Order represents the structure of order data
type Order struct {
	OrderUID          string    `json:"order_uid" db:"order_uid" validate:"required"`
	TrackNumber       string    `json:"track_number" db:"track_number" validate:"required"`
	Entry             string    `json:"entry" db:"entry" validate:"required"`
	Delivery          Delivery  `json:"delivery" db:"delivery" validate:"required"`
	Payment           Payment   `json:"payment" db:"payment" validate:"required"`
	Items             []Item    `json:"items" db:"items" validate:"required"`
	Locale            string    `json:"locale" db:"locale" validate:"required"`
	InternalSignature string    `json:"internal_signature" db:"internal_signature"`
	CustomerID        string    `json:"customer_id" db:"customer_id" validate:"required"`
	DeliveryService   string    `json:"delivery_service" db:"delivery_service" validate:"required"`
	Shardkey          string    `json:"shardkey" db:"shardkey" validate:"required"`
	SmID              int       `json:"sm_id" db:"sm_id" validate:"required"`
	DateCreated       time.Time `json:"date_created" db:"date_created" validate:"required"`
	OofShard          string    `json:"oof_shard" db:"oof_shard" validate:"required"`
}

// Delivery represents delivery information
type Delivery struct {
	Name    string `json:"name" db:"name" validate:"required"`
	Phone   string `json:"phone" db:"phone" validate:"required"`
	Zip     string `json:"zip" db:"zip" validate:"required"`
	City    string `json:"city" db:"city" validate:"required"`
	Address string `json:"address" db:"address" validate:"required"`
	Region  string `json:"region" db:"region" validate:"required"`
	Email   string `json:"email" db:"email" validate:"required"`
}

// Payment represents payment information
type Payment struct {
	Transaction  string `json:"transaction" db:"transaction" validate:"required"`
	RequestID    string `json:"request_id" db:"request_id"`
	Currency     string `json:"currency" db:"currency" validate:"required"`
	Provider     string `json:"provider" db:"provider" validate:"required"`
	Amount       int    `json:"amount" db:"amount" validate:"gt=0"`
	PaymentDt    int64  `json:"payment_dt" db:"payment_dt" validate:"required"`
	Bank         string `json:"bank" db:"bank" validate:"required"`
	DeliveryCost int    `json:"delivery_cost" db:"delivery_cost" validate:"gt=0"`
	GoodsTotal   int    `json:"goods_total" db:"goods_total" validate:"gt=0"`
	CustomFee    int    `json:"custom_fee" db:"custom_fee" validate:"gte=0"`
}

// Item represents an item in the order
type Item struct {
	ChrtID      int    `json:"chrt_id" db:"chrt_id" validate:"required"`
	TrackNumber string `json:"track_number" db:"track_number" validate:"required"`
	Price       int    `json:"price" db:"price" validate:"gt=0"`
	Rid         string `json:"rid" db:"rid" validate:"required"`
	Name        string `json:"name" db:"name" validate:"required"`
	Sale        int    `json:"sale" db:"sale" validate:"gte=0"`
	Size        string `json:"size" db:"size" validate:"required"`
	TotalPrice  int    `json:"total_price" db:"total_price" validate:"gt=0"`
	NmID        int    `json:"nm_id" db:"nm_id" validate:"required"`
	Brand       string `json:"brand" db:"brand" validate:"required"`
	Status      int    `json:"status" db:"status" validate:"required"`
}
