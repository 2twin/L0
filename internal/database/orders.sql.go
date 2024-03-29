// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: orders.sql

package database

import (
	"context"
	"encoding/json"
	"time"
)

const createOrder = `-- name: CreateOrder :one
INSERT INTO orders (
    order_uid,
    track_number,
    entry,
    delivery,
    payment,
    items,
    locale,
    internal_signature,
    customer_id,
    delivery_service,
    shardkey,
    sm_id,
    date_created,
    oof_shard
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING order_uid, track_number, entry, delivery, payment, items, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard
`

type CreateOrderParams struct {
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

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, createOrder,
		arg.OrderUid,
		arg.TrackNumber,
		arg.Entry,
		arg.Delivery,
		arg.Payment,
		arg.Items,
		arg.Locale,
		arg.InternalSignature,
		arg.CustomerID,
		arg.DeliveryService,
		arg.Shardkey,
		arg.SmID,
		arg.DateCreated,
		arg.OofShard,
	)
	var i Order
	err := row.Scan(
		&i.OrderUid,
		&i.TrackNumber,
		&i.Entry,
		&i.Delivery,
		&i.Payment,
		&i.Items,
		&i.Locale,
		&i.InternalSignature,
		&i.CustomerID,
		&i.DeliveryService,
		&i.Shardkey,
		&i.SmID,
		&i.DateCreated,
		&i.OofShard,
	)
	return i, err
}

const getAllOrders = `-- name: GetAllOrders :many
SELECT order_uid, track_number, entry, delivery, payment, items, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard FROM orders
`

func (q *Queries) GetAllOrders(ctx context.Context) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, getAllOrders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.OrderUid,
			&i.TrackNumber,
			&i.Entry,
			&i.Delivery,
			&i.Payment,
			&i.Items,
			&i.Locale,
			&i.InternalSignature,
			&i.CustomerID,
			&i.DeliveryService,
			&i.Shardkey,
			&i.SmID,
			&i.DateCreated,
			&i.OofShard,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOrder = `-- name: GetOrder :one
SELECT order_uid, track_number, entry, delivery, payment, items, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard FROM orders WHERE order_uid = $1
`

func (q *Queries) GetOrder(ctx context.Context, orderUid string) (Order, error) {
	row := q.db.QueryRowContext(ctx, getOrder, orderUid)
	var i Order
	err := row.Scan(
		&i.OrderUid,
		&i.TrackNumber,
		&i.Entry,
		&i.Delivery,
		&i.Payment,
		&i.Items,
		&i.Locale,
		&i.InternalSignature,
		&i.CustomerID,
		&i.DeliveryService,
		&i.Shardkey,
		&i.SmID,
		&i.DateCreated,
		&i.OofShard,
	)
	return i, err
}
