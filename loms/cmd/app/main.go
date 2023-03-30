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
	"route256/libs/kafka"
	"route256/libs/postgress/transactor"
	"route256/loms/internal/api/loms"
	"route256/loms/internal/domain"
	"route256/loms/internal/notification"
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

	var ns domain.NotificationSender
	{
		brokers := make([]string, 0, len(config.ConfigData.Kafka.Brokers))
		for _, broker := range config.ConfigData.Kafka.Brokers {
			brokers = append(brokers, broker.Url())
		}

		producer, err := kafka.NewSyncProducer(brokers)
		if err != nil {
			log.Fatalln(err)
		}

		asyncProducer, err := kafka.NewAsyncProducer(brokers)
		if err != nil {
			log.Fatalln(err)
		}

		onSuccess := func(id string) {
			log.Println("order success", id)
		}
		onFailed := func(id string) {
			log.Println("order failed", id)
		}

		ns = notification.NewOrderSender(
			producer,
			asyncProducer,
			config.ConfigData.Kafka.OrderStatusTopic,
			onSuccess, onFailed,
		)
	}

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

		dmn := domain.New(stockHandler, tr, ns)
		desc.RegisterLomsServer(s, loms.New(dmn))

		{
			ticker := time.NewTicker(time.Minute * 1)

			go func() {
				for {
					select {
					case <-ticker.C:
						orders, err := dmn.CancelUnpayedOrders(ctx, time.Now().Add(time.Duration(-10)*time.Minute))
						if err != nil {
							log.Printf("Falied to cancel updayed orders due to %v \n", err)
						} else if len(orders) == 0 {
							log.Println("Nothing to cancel")
						} else {
							for _, orderId := range orders {
								log.Printf("Order %d has been canceled\n", orderId)
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
