package model

type OrderStatus int32

const (
	StatusUnknown OrderStatus = iota
	StatusNew
	StatusAwaitingPayment
	StatusFalied
	StatusPayed
	StatusCancelled
)

type StockItem struct {
	WarehouseID int64
	Count       uint64
}

type Item struct {
	Sku   uint32
	Count uint16
}

type Order struct {
	Status OrderStatus
	User   int64
	Items  []*Item
}
