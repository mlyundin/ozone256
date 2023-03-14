package main

import (
	"context"
	"log"
	"net"
	"route256/checkout/internal/api/checkout"
	"route256/checkout/internal/clients/loms"
	"route256/checkout/internal/clients/products"
	"route256/checkout/internal/domain"
	"route256/checkout/internal/repository/postgress"
	desc "route256/checkout/pkg/checkout"
	"route256/libs/config"
	"route256/libs/interceptors"
	"route256/libs/postgress/transactor"
	lomcln "route256/loms/pkg/client/grpc/loms-service"
	productcln "route256/product/pkg/client/grpc/product-service"
	"time"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal("config init", err)
	}

	connLoms, err := grpc.DialContext(context.Background(), config.ConfigData.Services.Loms.Url(),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer connLoms.Close()
	lomsClient := loms.New(lomcln.New(connLoms))

	connProduct, err := grpc.DialContext(context.Background(), config.ConfigData.Services.Products.Url(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer connProduct.Close()
	productClient := products.New(productcln.New(connProduct), config.ConfigData.Services.Products.Token)

	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, config.ConfigData.Databases.Checkout.Connection())
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
	cartHandler := respository.New(transactor.New(pool))

	{
		lis, err := net.Listen("tcp", ":"+config.ConfigData.Services.Checkout.Port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		server := grpc.NewServer(
			grpc.UnaryInterceptor(
				grpcMiddleware.ChainUnaryServer(
					interceptors.LoggingInterceptor,
				),
			),
		)
		reflection.Register(server)

		domain := domain.New(lomsClient, productClient, cartHandler)
		desc.RegisterCheckoutServer(server, checkout.New(domain))

		log.Printf("server listening at %v", lis.Addr())
		if err = server.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}
}
