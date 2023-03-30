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

func Status2Str(status OrderStatus) string {
	switch status {
	case StatusUnknown:
		return "Unknown"

	case StatusNew:
		return "New"

	case StatusAwaitingPayment:
		return "AwaitingPayment"

	case StatusFalied:
		return "Failed"

	case StatusPayed:
		return "Payed"

	case StatusCancelled:
		return "Canceled"
	}

	return ""
}

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
