package main

import (
	"fmt"
	"log"
	"net"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"route256/libs/interceptors"
	"route256/loms/internal/api/loms"
	"route256/loms/internal/domain"
	desc "route256/loms/pkg/loms"
)

const grpcPort = 8081

func main() {

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
	desc.RegisterLomsServer(s, loms.New(domain.New()))

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
