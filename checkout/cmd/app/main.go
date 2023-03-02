package main

import (
	"context"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"route256/checkout/internal/clients/loms"
	"route256/checkout/internal/clients/products"
	"route256/checkout/internal/config"
	"route256/checkout/internal/domain"
	"route256/checkout/internal/handlers"
	"route256/libs/srvwrapper"
	lomcln "route256/loms/pkg/client/grpc/loms-service"
)

const port = ":8080"

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
	businessLogic := domain.New(lomsClient, productClient)

	handler := handlers.New(businessLogic)
	http.Handle("/addToCart", srvwrapper.New(handler.AddToCart))
	http.Handle("/deleteFromCart", srvwrapper.New(handler.DeleteFromCart))
	http.Handle("/listCart", srvwrapper.New(handler.ListCart))
	http.Handle("/purchase", srvwrapper.New(handler.Purchase))

	log.Println("listening http at", port)
	err = http.ListenAndServe(port, nil)
	log.Fatal("cannot listen http", err)
}
