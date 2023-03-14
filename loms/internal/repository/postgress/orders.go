package respository

import (
	"context"
	"errors"
	"route256/loms/internal/repository/schema"
	"route256/loms/pkg/model"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

var (
	ordersColumns = []string{"status", "user_id", "creation_time"}
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
		Values(model.StatusNew, user, time.Now().Unix()).
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

func (r *LomsRepo) UpdateStatusBefore(ctx context.Context, beforeTimestamp int64, newStatus model.OrderStatus, currStatus model.OrderStatus) ([]int64, error) {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)

	sql, args, err := sq.Update(ordersTable).
		Where(sq.LtOrEq{"creation_time": beforeTimestamp}).
		Where(sq.Eq{"status": currStatus}).
		Set("status", newStatus).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING order_id").
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]int64, 0)
	for rows.Next() {
		var orderID int64
		rows.Scan(&orderID)
		res = append(res, orderID)
	}

	return res, nil
}

func (r *LomsRepo) GetOrder(ctx context.Context, orderId int64) (*model.Order, error) {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)

	sql, args, err := sq.Select("order_id", "status", "user_id").
		From(ordersTable).
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
