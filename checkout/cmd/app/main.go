package main

import (
	"context"
	"log"
	"net"
	"route256/checkout/internal/api/checkout"
	"route256/checkout/internal/clients/loms"
	"route256/checkout/internal/clients/products"
	"route256/checkout/internal/domain"
	desc "route256/checkout/pkg/checkout"
	"route256/libs/config"
	"route256/libs/interceptors"
	lomcln "route256/loms/pkg/client/grpc/loms-service"
	productcln "route256/product/pkg/client/grpc/product-service"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
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

	domain := domain.New(lomsClient, productClient)

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
	desc.RegisterCheckoutServer(server, checkout.New(domain))

	log.Printf("server listening at %v", lis.Addr())

	if err = server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
