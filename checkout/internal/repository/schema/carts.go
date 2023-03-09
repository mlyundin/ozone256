package schema

type CartItem struct {
	User  int64 `db:"user_id"`
	Sku   int64 `db:"sku"`
	Count int32 `db:"count"`
}
