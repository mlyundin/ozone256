package respository

import (
	"context"
	"errors"

	"route256/loms/internal/repository/schema"
	"route256/loms/pkg/model"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

var (
	ordersColumns = []string{"status", "user_id"}
)

const (
	ordersTable = "orders"
)

var (
	ErrUnknownOrderId   = errors.New("unknow order id")
	ErrStatusUpdateFail = errors.New("status update fails")
	ErrOrderNotFound    = errors.New("order not found")
)

func (r *LomsRepo) NewOrder(ctx context.Context, user int64) (int64, error) {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)

	sql, args, err := sq.Insert(ordersTable).
		Columns(ordersColumns...).
		Values(model.StatusNew, user).
		Suffix("RETURNING order_id").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return -1, err
	}

	rows, err := db.Query(ctx, sql, args...)
	if err != nil {
		return -1, err
	}
	defer rows.Close()

	if rows.Next() {
		var orderId int64
		err = rows.Scan(&orderId)
		if err != nil {
			return -1, err
		}

		return orderId, nil
	}

	return -1, ErrUnknownOrderId
}

func (r *LomsRepo) UpdateStatus(ctx context.Context, orderId int64, newStatus model.OrderStatus, currStatus model.OrderStatus) error {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)

	sql, args, err := sq.Update(ordersTable).
		Where(sq.Eq{"order_id": orderId}).
		Where(sq.Eq{"status": currStatus}).
		Set("status", newStatus).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING status").
		ToSql()

	if err != nil {
		return err
	}

	rows, err := db.Query(ctx, sql, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		return nil
	}

	return ErrStatusUpdateFail
}

func (r *LomsRepo) GetOrder(ctx context.Context, orderId int64) (*model.Order, error) {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)

	sql, args, err := sq.Select("order_id", "status", "user_id").
		From(stockTable).
		Where(sq.Eq{"order_id": orderId}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	var items []schema.Order
	if err := pgxscan.Select(ctx, db, &items, sql, args...); err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, ErrOrderNotFound
	}

	return &model.Order{Status: model.OrderStatus(items[0].Status), User: items[0].UserId}, nil
}
