package domain

import (
	"context"
	"route256/loms/pkg/model"
	"time"
)

func (s *domainmodel) CancelOrder(ctx context.Context, orderId int64) error {
	return s.tm.RunRepeteableRead(ctx, func(ctxTX context.Context) error {
		err := s.lomsRepo.UpdateStatus(ctxTX, orderId, model.StatusCancelled, model.StatusAwaitingPayment)
		if err != nil {
			return err
		}

		reservations, err := s.lomsRepo.GetReservations(ctxTX, orderId)
		if err != nil {
			return err
		}

		for _, r := range reservations {
			err = s.lomsRepo.Release(ctxTX, r)
			if err != nil {
				return err
			}
		}

		err = s.lomsRepo.ReleaseAllReservations(ctxTX, orderId)
		if err != nil {
			return err
		}

		return nil
	})
}

func (s *domainmodel) CancelUnpayedOrders(ctx context.Context, beforeTimestamp time.Time) ([]int64, error) {
	var res []int64
	err := s.tm.RunRepeteableRead(ctx, func(ctxTX context.Context) error {
		orders, err := s.lomsRepo.UpdateStatusBefore(ctxTX, beforeTimestamp.Unix(), model.StatusCancelled, model.StatusAwaitingPayment)
		res = orders

		if err != nil {
			return err
		}

		for _, orderId := range orders {
			reservations, err := s.lomsRepo.GetReservations(ctxTX, orderId)
			if err != nil {
				return err
			}

			for _, r := range reservations {
				err = s.lomsRepo.Release(ctxTX, r)
				if err != nil {
					return err
				}
			}

			err = s.lomsRepo.ReleaseAllReservations(ctxTX, orderId)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
