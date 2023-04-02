package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"route256/libs/config"
	"route256/libs/interceptors"
	"route256/libs/kafka"
	"route256/libs/logger"
	"route256/libs/postgress/transactor"
	"route256/loms/internal/api/loms"
	"route256/loms/internal/domain"
	"route256/loms/internal/notification"
	"route256/loms/internal/repository/postgress"
	desc "route256/loms/pkg/loms"
	"time"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal("config init", err)
	}

	logger.Init(config.ConfigData.Services.Logging.Devel)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	pool, err := pgxpool.Connect(ctx, config.ConfigData.Databases.Loms.Connection())
	if err != nil {
		logger.Fatal("failed to connect to postgress", zap.Error(err))
	}
	defer pool.Close()
	{
		config := pool.Config()
		config.MaxConnIdleTime = time.Minute
		config.MaxConnLifetime = time.Hour
		config.MinConns = 2
		config.MaxConns = 10

		if err := pool.Ping(ctx); err != nil {
			logger.Fatal("failed ping to postgress", zap.Error(err))
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
			logger.Fatal("failed to create kafka sync producer", zap.Error(err))
		}

		asyncProducer, err := kafka.NewAsyncProducer(brokers)
		if err != nil {
			logger.Fatal("failed to create kafka async producer", zap.Error(err))
		}

		onSuccess := func(id string) {
			logger.Info("order success", zap.String("id", id))
		}
		onFailed := func(id string) {
			logger.Error("order failed", zap.String("id", id))
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
			logger.Fatal("failed to listen", zap.Error(err))
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
							logger.Error("Falied to cancel updayed orders due to", zap.Error(err))
						} else if len(orders) == 0 {
							logger.Info("Nothing to cancel")
						} else {
							logger.Debug("Orders has been canceled", zap.Int64s("ids", orders))
						}

					case <-ctx.Done():
						logger.Info("Goodbye!")
						return
					}
				}
			}()

		}

		logger.Info("server listening at ", zap.String("addr", lis.Addr().String()))
		if err = s.Serve(lis); err != nil {
			logger.Fatal("failed to serve", zap.Error(err))
		}
	}
}
