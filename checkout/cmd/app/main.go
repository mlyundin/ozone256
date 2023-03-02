package main

import (
	"context"
	"fmt"
	"log"
	"net"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"route256/libs/interceptors"
	lomcln "route256/loms/pkg/client/grpc/loms-service"

	"route256/checkout/internal/api/checkout"
	"route256/checkout/internal/clients/loms"
	"route256/checkout/internal/clients/products"
	"route256/checkout/internal/config"
	"route256/checkout/internal/domain"

	desc "route256/checkout/pkg/checkout"
)

const grpcPort = 8080

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal("config init", err)
	}

	conn, err := grpc.DialContext(context.Background(), config.ConfigData.Services.Loms, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	lomsClient := loms.New(lomcln.New(conn))
	productClient := products.New(config.ConfigData.Services.Products.Url, config.ConfigData.Services.Products.Token)
	domain := domain.New(lomsClient, productClient)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
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
	desc.RegisterCheckoutServer(server, checkout.New(domain))

	log.Printf("server listening at %v", lis.Addr())

	if err = server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
