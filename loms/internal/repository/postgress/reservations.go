package respository

import (
	"context"
	"errors"
	"route256/loms/internal/domain"
	"route256/loms/internal/repository/schema"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
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

func (r *LomsRepo) NewReservation(ctx context.Context, reservation *domain.Reservation) error {
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

func (r *LomsRepo) ReleaseAllReservations(ctx context.Context, orderId int64) error {
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

func (r *LomsRepo) GetReservations(ctx context.Context, orderId int64) ([]*domain.Reservation, error) {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)

	sql, args, err := sq.Select(reservationsColumns...).From(reservationsTable).Where(sq.Eq{"order_id": orderId}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	var items []schema.Reservation
	if err := pgxscan.Select(ctx, db, &items, sql, args...); err != nil {
		return nil, err
	}

	reservations := make([]*domain.Reservation, 0, len(items))
	for _, item := range items {
		reservations = append(reservations, &domain.Reservation{OrderId: orderId, WarehouseID: item.WarhouseId,
			Sku: uint32(item.Sku), Count: uint16(item.Count)})
	}

	return reservations, nil
}
