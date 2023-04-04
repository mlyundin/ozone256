package main

import (
	"context"
	"log"
	"net"
	"route256/checkout/internal/api/checkout"
	"route256/checkout/internal/clients/loms"
	"route256/checkout/internal/clients/products"
	"route256/checkout/internal/domain"
	respository "route256/checkout/internal/repository/postgress"
	desc "route256/checkout/pkg/checkout"
	"route256/libs/config"
	"route256/libs/interceptors"
	"route256/libs/logger"
	"route256/libs/metrics"
	"route256/libs/postgress/transactor"
	lomcln "route256/loms/pkg/client/grpc/loms-service"
	productcln "route256/product/pkg/client/grpc/product-service"
	"time"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

const (
	productRateLimit = 10
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal("config init", err)
	}

	logger.Init(config.ConfigData.Services.Logging.Devel)

	connLoms, err := grpc.DialContext(context.Background(), config.ConfigData.Services.Loms.Url(),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("failed to connect to server", zap.Error(err))
	}
	defer connLoms.Close()
	lomsClient := loms.New(lomcln.New(connLoms))

	inter := interceptors.NewClientRateLimiterInterceptor(rate.NewLimiter(rate.Every(time.Second/productRateLimit), productRateLimit))
	connProduct, err := grpc.DialContext(context.Background(),
		config.ConfigData.Services.Products.Url(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(inter.Intercept))
	if err != nil {
		logger.Fatal("failed to connect to server", zap.Error(err))
	}
	defer connProduct.Close()
	productClient := products.New(productcln.New(connProduct), config.ConfigData.Services.Products.Token)

	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, config.ConfigData.Databases.Checkout.Connection())
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
	cartHandler := respository.New(transactor.New(pool))

	{
		lis, err := net.Listen("tcp", ":"+config.ConfigData.Services.Checkout.Port)
		if err != nil {
			logger.Fatal("failed to listen", zap.Error(err))
		}

		server := grpc.NewServer(
			grpc.UnaryInterceptor(
				grpcMiddleware.ChainUnaryServer(
					interceptors.LoggingInterceptor,
					metrics.Intercept,
				),
			),
		)
		reflection.Register(server)

		domain := domain.New(lomsClient, productClient, cartHandler)
		desc.RegisterCheckoutServer(server, checkout.New(domain))

		go metrics.RunHttpServer(config.ConfigData.Services.Checkout.MetricsPort)

		logger.Info("server listening at", zap.String("adress", lis.Addr().String()))
		if err = server.Serve(lis); err != nil {
			logger.Fatal("failed to serve", zap.Error(err))

		}
	}
}
