-- name: CreateOrder :one
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
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders WHERE order_uid = $1;