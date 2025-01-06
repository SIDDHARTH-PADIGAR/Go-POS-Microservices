package main

import (
	handler "go-pos-ms/services/orders/handler/orders"
	"go-pos-ms/services/orders/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}
	gRPCServer := grpc.NewServer()

	//Register
	orderService := service.NewOrderService()
	handler.NewGrpcOrdersService(gRPCServer, orderService)

	log.Println("gRPC server is running on", s.addr)

	return gRPCServer.Serve(lis)
}
