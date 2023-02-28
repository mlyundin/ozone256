package loms_v1

import (
	"route256/loms/internal/service/loms"
	desc "route256/loms/pkg/loms_v1"
)

type Implementation struct {
	desc.UnimplementedLomsV1Server

	service loms.Service
}

func New(service loms.Service) *Implementation {
	return &Implementation{
		desc.UnimplementedLomsV1Server{},
		service,
	}
}
