package domain

import (
	"context"
	"route256/loms/pkg/model"
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
