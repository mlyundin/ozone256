package main

import (
	"context"
	"log"
	"net"
	"route256/libs/config"
	"route256/libs/interceptors"
	"route256/libs/postgress/transactor"
	"route256/loms/internal/api/loms"
	"route256/loms/internal/domain"
	"route256/loms/internal/repository/postgress"
	desc "route256/loms/pkg/loms"
	"time"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal("config init", err)
	}

	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, config.ConfigData.Databases.Loms.Connection())
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	{
		config := pool.Config()
		config.MaxConnIdleTime = time.Minute
		config.MaxConnLifetime = time.Hour
		config.MinConns = 2
		config.MaxConns = 10

		if err := pool.Ping(ctx); err != nil {
			log.Fatal(err)
		}
	}

	tr := transactor.New(pool)
	stockHandler := respository.New(tr)

	{
		lis, err := net.Listen("tcp", ":"+config.ConfigData.Services.Loms.Port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer(
			grpc.UnaryInterceptor(
				grpcMiddleware.ChainUnaryServer(
					interceptors.LoggingInterceptor,
				),
			),
		)

		reflection.Register(s)
		desc.RegisterLomsServer(s, loms.New(domain.New(stockHandler, tr)))

		log.Printf("server listening at %v", lis.Addr())

		if err = s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}
}
