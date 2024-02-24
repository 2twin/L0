package order

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/2twin/L0/internal/database"
	"github.com/2twin/L0/internal/model"
	"github.com/jellydator/ttlcache/v3"

	_ "github.com/lib/pq"
)

type repository struct {
	cache *ttlcache.Cache[string, model.Order]
	db    *database.Queries
}

func NewRepository(dbURL string) (*repository, error) {
	db, err := NewDB(dbURL)
	if err != nil {
		return nil, err
	}

	return &repository{
		cache: ttlcache.New[string, model.Order](),
		db:    database.New(db),
	}, nil
}

func NewDB(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %s", err)
	}

	return db, nil
}

func (r *repository) Create(ctx context.Context, orderUUID string, order *model.Order) error {
	r.cache.Set(orderUUID, *order, ttlcache.NoTTL)

	delivery, err := json.Marshal(order.Delivery)
	if err != nil {
		return err
	}

	payment, err := json.Marshal(order.Payment)
	if err != nil {
		return err
	}

	items, err := json.Marshal(order.Items)
	if err != nil {
		return err
	}

	r.db.CreateOrder(ctx, database.CreateOrderParams{
		OrderUid:          order.OrderUID,
		TrackNumber:       order.TrackNumber,
		Entry:             order.Entry,
		Delivery:          delivery,
		Payment:           payment,
		Items:             items,
		Locale:            order.Locale,
		InternalSignature: order.InternalSignature,
		CustomerID:        order.CustomerID,
		DeliveryService:   order.DeliveryService,
		Shardkey:          order.Shardkey,
		SmID:              int32(order.SmID),
		DateCreated:       order.DateCreated,
		OofShard:          order.OofShard,
	})

	return nil
}

func (r *repository) Get(ctx context.Context, orderUUID string) (*model.Order, error) {
	cacheOrder := r.cache.Get(orderUUID)
	if cacheOrder != nil {
		val := cacheOrder.Value()
		return &val, nil
	}

	dbOrder, err := r.db.GetOrder(ctx, orderUUID)
	if err != nil {
		return nil, err
	}

	order, err := fromDBOrderToOrder(dbOrder)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func fromDBOrderToOrder(dbOrder database.Order) (*model.Order, error) {
	delivery, err := fromDBDeliveryToDelivery(dbOrder.Delivery)
	if err != nil {
		return nil, err
	}

	payment, err := fromDBPaymentToPayment(dbOrder.Payment)
	if err != nil {
		return nil, err
	}

	items, err := fromDBItemsToItems(dbOrder.Items)
	if err != nil {
		return nil, err
	}

	return &model.Order{
		OrderUID:          dbOrder.OrderUid,
		TrackNumber:       dbOrder.TrackNumber,
		Entry:             dbOrder.Entry,
		Delivery:          *delivery,
		Payment:           *payment,
		Items:             *items,
		Locale:            dbOrder.Locale,
		InternalSignature: dbOrder.InternalSignature,
		CustomerID:        dbOrder.CustomerID,
		DeliveryService:   dbOrder.DeliveryService,
		Shardkey:          dbOrder.Shardkey,
		SmID:              int(dbOrder.SmID),
		DateCreated:       dbOrder.DateCreated,
		OofShard:          dbOrder.OofShard,
	}, nil
}

func fromDBDeliveryToDelivery(dbDelivery []byte) (*model.Delivery, error) {
	var delivery model.Delivery
	err := json.Unmarshal(dbDelivery, &delivery)
	if err != nil {
		return nil, err
	}
	return &delivery, nil
}

func fromDBPaymentToPayment(dbPayment []byte) (*model.Payment, error) {
	var payment model.Payment
	err := json.Unmarshal(dbPayment, &payment)
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func fromDBItemsToItems(dbItems []byte) (*[]model.Item, error) {
	var items []model.Item
	err := json.Unmarshal(dbItems, &items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}
