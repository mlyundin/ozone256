package schema

type Reservation struct {
	OrderId    int64 `db:"order_id"`
	WarhouseId int64 `db:"warehouse_id"`
	Sku        int64 `db:"sku"`
	Count      int32 `db:"count"`
}
