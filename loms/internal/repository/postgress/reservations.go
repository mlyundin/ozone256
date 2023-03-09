package respository

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"route256/loms/internal/domain"
)

var (
	reservationsColumns = []string{"order_id", "warehouse_id", "sku", "count"}
)

const (
	reservationsTable = "reservations"
)

var (
	ErrReservationFail = errors.New("reservation fail")
)

func (r *LomsRepo) Reserve(ctx context.Context, reservation domain.Reservation) error {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)

	sql, args, err := sq.Insert(reservationsTable).
		Columns(reservationsColumns...).
		Values(reservation.OrderId, reservation.WarehouseID, reservation.Sku, reservation.Count).
		Suffix("RETURNING order_id").
		PlaceholderFormat(sq.Dollar).
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

	return ErrReservationFail

}

func (r *LomsRepo) Release(ctx context.Context, orderId int64) error {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)

	sql, args, err := sq.Delete(reservationsTable).
		Where(sq.Eq{"order_id": orderId}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	rows, err := db.Query(ctx, sql, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
