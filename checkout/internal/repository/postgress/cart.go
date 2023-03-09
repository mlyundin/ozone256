package respository

import (
	"context"
	"errors"
	"fmt"

	//"fmt"

	"route256/libs/postgress/transactor"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"

	"route256/checkout/internal/domain"
	"route256/checkout/internal/repository/schema"
)

type CartRepo struct {
	transactor.QueryEngineProvider
}

func New(provider transactor.QueryEngineProvider) *CartRepo {
	return &CartRepo{
		QueryEngineProvider: provider,
	}
}

var (
	cartsColumns = []string{"user_id", "sku", "count"}
)

const (
	cartsTable = "carts"
)

func (r *CartRepo) AddToCart(ctx context.Context, item *domain.CartItem) error {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)

	sql, args, err := sq.Insert(cartsTable).
		Columns(cartsColumns...).
		Values(item.User, item.Sku, item.Count).
		Suffix(fmt.Sprintf("ON CONFLICT (user_id, sku) DO UPDATE SET count = %s.count + EXCLUDED.count", cartsTable)).
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

var (
	ErrNothingToDelete = errors.New("insufficient count")
)

func (r *CartRepo) DeleteFromCart(ctx context.Context, item *domain.CartItem) error {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)

	sql, args, err := sq.Update(cartsTable).
		Where(sq.Eq{"user_id": item.User}).
		Where(sq.Eq{"sku": item.Sku}).
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

	// TODO how to split no data case and count constrain violation case
	return ErrNothingToDelete
}

func (r *CartRepo) GetItemCount(ctx context.Context, userId int64, sku uint32) (uint16, error) {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)

	sql, args, err := sq.Select(cartsColumns...).
		From(cartsTable).
		Where(sq.Eq{"user_id": userId}).
		Where(sq.Eq{"sku": sku}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return 0, err
	}

	var items []schema.CartItem
	if err := pgxscan.Select(ctx, db, &items, sql, args...); err != nil {
		return 0, err
	}

	if len(items) == 0 {
		return 0, nil
	}

	return uint16(items[0].Count), nil
}

func (r *CartRepo) ListCart(ctx context.Context, userId int64) (*domain.Cart, error) {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)

	sql, args, err := sq.Select(cartsColumns...).From(cartsTable).Where(sq.Eq{"user_id": userId}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	var items []schema.CartItem
	if err := pgxscan.Select(ctx, db, &items, sql, args...); err != nil {
		return nil, err
	}

	domainItems := make([]domain.Product, 0, len(items))
	for _, item := range items {
		domainItems = append(domainItems, domain.Product{Sku: uint32(item.Sku), Count: uint16(item.Count)})
	}

	return &domain.Cart{Items: domainItems}, nil
}
