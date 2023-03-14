package schema

type StocktItem struct {
	WarhouseId int64 `db:"warehouse_id"`
	Sku        int64 `db:"sku"`
	Count      int64 `db:"count"`
}
