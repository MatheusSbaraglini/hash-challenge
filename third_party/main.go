package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/matheussbaraglini/hash-challenge/discount"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type discountServer struct{}

func main() {
	log := log.New(os.Stderr, "", log.LstdFlags)

	listner, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	discount.RegisterDiscountServer(grpcServer, &discountServer{})
	reflection.Register(grpcServer)

	log.Printf("Discount service is running on: %s", listner.Addr().String())

	if err := grpcServer.Serve(listner); err != nil {
		panic(err)
	}
}

func (ds *discountServer) GetDiscount(ctx context.Context, in *discount.GetDiscountRequest) (*discount.GetDiscountResponse, error) {
	return &discount.GetDiscountResponse{
		Percentage: 10,
	}, nil
}
