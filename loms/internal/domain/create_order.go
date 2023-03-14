package domain

import (
	"context"
	"errors"
	"route256/loms/pkg/model"
)

var (
	ErrInsufficientStock = errors.New("insufficient stock")
)

func (dm *domainmodel) CreateOrder(ctx context.Context, user int64, items []*model.Item) (int64, error) {
	orderId, err := dm.lomsRepo.NewOrder(ctx, user)
	if err != nil {
		return -1, err
	}

	err = dm.tm.RunRepeteableRead(ctx, func(ctxTX context.Context) error {
		err = dm.lomsRepo.UpdateStatus(ctxTX, orderId, model.StatusAwaitingPayment, model.StatusNew)

		for _, item := range items {
			stocks, err := dm.lomsRepo.Stocks(ctxTX, item.Sku)
			if err != nil {
				return err
			}

			reservations := make([]Reservation, 0)
			count := uint64(item.Count)
			for _, stock := range stocks {
				if count == 0 {
					break
				}

				to_use := stock.Count
				if to_use > count {
					to_use = count
				}
				count -= to_use

				reservations = append(reservations, Reservation{OrderId: orderId, WarehouseID: stock.WarehouseID,
					Sku: item.Sku, Count: uint16(to_use)})
			}

			if count != 0 {
				return ErrInsufficientStock
			}

			for _, item := range reservations {
				err = dm.lomsRepo.Reserve(ctxTX, &item)
				if err != nil {
					return err
				}

				err = dm.lomsRepo.NewReservation(ctxTX, &item)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})

	if err != nil {
		err = dm.lomsRepo.UpdateStatus(ctx, orderId, model.StatusFalied, model.StatusNew)
		if err != nil {
			return -1, err
		}
	}

	dm.newOrderQueue.Add(orderId)

	return orderId, nil
}
