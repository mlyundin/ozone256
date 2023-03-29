package domain

import (
	"context"
	"route256/loms/pkg/model"
)

func (s *domainmodel) OrderPayed(ctx context.Context, orderId int64) error {
	return s.tm.RunRepeteableRead(ctx, func(ctxTX context.Context) error {
		err := s.UpdateStatus(ctxTX, orderId, model.StatusPayed, model.StatusAwaitingPayment)
		if err != nil {
			return err
		}

		err = s.lomsRepo.ReleaseAllReservations(ctxTX, orderId)
		if err != nil {
			return err
		}

		return nil
	})
}
