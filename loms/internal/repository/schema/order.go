package schema

type Order struct {
	OrderId int64 `db:"order_id"`
	Status  int32 `db:"status"`
	UserId  int64 `db:"user_id"`
}
