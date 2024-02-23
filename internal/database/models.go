// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"encoding/json"
	"time"
)

type Order struct {
	OrderUid          string
	TrackNumber       string
	Entry             string
	Delivery          json.RawMessage
	Payment           json.RawMessage
	Items             json.RawMessage
	Locale            string
	InternalSignature string
	CustomerID        string
	DeliveryService   string
	Shardkey          string
	SmID              int32
	DateCreated       time.Time
	OofShard          string
}
