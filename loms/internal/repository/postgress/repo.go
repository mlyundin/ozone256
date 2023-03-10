package respository

import "route256/libs/postgress/transactor"

type LomsRepo struct {
	transactor.QueryEngineProvider
}

func New(provider transactor.QueryEngineProvider) *LomsRepo {
	return &LomsRepo{
		QueryEngineProvider: provider,
	}
}
