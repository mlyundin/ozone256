package main

import (
	"fmt"
	"log"
	"net"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"route256/libs/interceptors"
	lomsV1 "route256/loms/internal/api/loms_v1"
	"route256/loms/internal/service/loms"
	desc "route256/loms/pkg/loms_v1"
)

const port = ":8081"
const grpcPort = 8081

func main() {
	// handler := handlers.New()
	// http.Handle("/stocks", srvwrapper.New(handler.Stocks))
	// http.Handle("/createOrder", srvwrapper.New(handler.CreateOrder))
	// http.Handle("/listOrder", srvwrapper.New(handler.ListOrder))
	// http.Handle("/cancelOrder", srvwrapper.New(handler.CancelOrder))
	// http.Handle("/orderPayed", srvwrapper.New(handler.OrderPayed))

	// log.Println("listening http at", port)
	// err := http.ListenAndServe(port, nil)
	// log.Fatal("cannot listen http", err)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
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
	desc.RegisterLomsV1Server(s, lomsV1.New(loms.NewService()))

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
