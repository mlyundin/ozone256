package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"route256/libs/config"
	"route256/libs/interceptors"
	"route256/libs/postgress/transactor"
	"route256/loms/internal/api/loms"
	"route256/loms/internal/domain"
	"route256/loms/internal/repository/postgress"
	"route256/loms/internal/utils"
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

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

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

		newOrderQueue := utils.NewTimeQueue[int64]()
		dmn := domain.New(stockHandler, tr, newOrderQueue)
		reflection.Register(s)
		desc.RegisterLomsServer(s, loms.New(dmn))

		{
			ticker := time.NewTicker(time.Second * 1)

			go func() {
				for {
					select {
					case <-ticker.C:
						for _, orderID := range newOrderQueue.Before(time.Now().Add(time.Duration(-5) * time.Second)) {
							err := dmn.CancelOrder(ctx, orderID)
							if err != nil {
								log.Println(err)
							} else {
								log.Printf("Order %d has been canceled", orderID)
							}
						}

					case <-ctx.Done():
						fmt.Println("Goodbye!")
						return
					}
				}
			}()

		}

		log.Printf("server listening at %v", lis.Addr())

		if err = s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}
}
