package respository

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"

	"route256/loms/internal/repository/schema"
	"route256/loms/pkg/model"
)

var (
	stockColumns = []string{"warehouse_id", "sku", "count"}
)

const (
	stockTable = "stock"
)

func (r *LomsRepo) Stocks(ctx context.Context, sku uint32) ([]*model.StockItem, error) {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)

	sql, args, err := sq.Select(stockColumns...).From(stockTable).Where(sq.Eq{"sku": sku}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	var items []schema.StocktItem
	if err := pgxscan.Select(ctx, db, &items, sql, args...); err != nil {
		return nil, err
	}

	domainItems := make([]*model.StockItem, 0, len(items))
	for _, item := range items {
		domainItems = append(domainItems, &model.StockItem{WarehouseID: item.WarhouseId, Count: uint64(item.Count)})
	}

	return domainItems, nil
}

var (
	ErrInsufficientCount = errors.New("insufficient count")
	ErrUnknownStock      = errors.New("unknow stock")
)

func (r *LomsRepo) ReserveStock(ctx context.Context, sku uint32, item *model.StockItem) error {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)

	sql, args, err := sq.Update(stockTable).
		Where(sq.Eq{"sku": sku}).
		Where(sq.Eq{"warehouse_id": item.WarehouseID}).
		Set("count", sq.Expr("count - ?", item.Count)).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING count").
		ToSql()

	if err != nil {
		return err
	}

	rows, err := db.Query(ctx, sql, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		return nil
	}

	return ErrInsufficientCount
}

func (r *LomsRepo) AddStock(ctx context.Context, sku uint32, item *model.StockItem) error {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)

	sql, args, err := sq.Update(stockTable).
		Where(sq.Eq{"sku": sku}).
		Where(sq.Eq{"warehouse_id": item.WarehouseID}).
		Set("count", sq.Expr("count + ?", item.Count)).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING count").
		ToSql()

	if err != nil {
		return err
	}

	rows, err := db.Query(ctx, sql, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		return nil
	}

	return ErrUnknownStock
}
