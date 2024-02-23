package order

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/2twin/L0/internal/database"
	"github.com/2twin/L0/internal/model"
	repoOrder "github.com/2twin/L0/internal/repository/order/model"
	"github.com/jellydator/ttlcache/v3"

	_ "github.com/lib/pq"
)

type repository struct {
	cache *ttlcache.Cache[string, repoOrder.Order]
	db    *database.Queries
}

// type PostgresClient interface {
// 	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
// 	PrepareContext(context.Context, string) (*sql.Stmt, error)
// 	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
// 	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
// }

func NewRepository(dbURL string) (*repository, error) {
	db, err := NewDB(dbURL)
	if err != nil {
		return nil, err
	}

	return &repository{
		cache: ttlcache.New[string, repoOrder.Order](),
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

func (r *repository) Create(_ context.Context, orderUUID string, order *model.Order) error {
	panic("unimplemented!")
}

func (r *repository) Get(ctx context.Context, orderUUID string) (*model.Order, error) {
	// order, err := r.db.GetOrder(ctx, orderUUID)
	return &model.Order{
		OrderUID: "abc",
	}, nil
}
