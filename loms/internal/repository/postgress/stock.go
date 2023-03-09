package respository

import (
	"context"
	// "errors"
	// "fmt"

	"route256/libs/postgress/transactor"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"

	// "route256/loms/internal/domain"
	"route256/loms/internal/repository/schema"
	"route256/loms/pkg/model"
)

type StockRepo struct {
	transactor.QueryEngineProvider
}

func New(provider transactor.QueryEngineProvider) *StockRepo {
	return &StockRepo{
		QueryEngineProvider: provider,
	}
}

var (
	stockColumns = []string{"warehouse_id", "sku", "count"}
)

const (
	stockTable = "stock"
)

func (r *StockRepo) Stocks(ctx context.Context, sku uint32) ([]*model.StockItem, error) {
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
