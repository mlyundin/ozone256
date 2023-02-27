package model

type StockItem struct {
	WarehouseID int64
	Count       uint64
}

type Item struct {
	Sku   uint32
	Count uint16
}

type Order struct {
	Status string
	User   int64
	Items  []*Item
}
