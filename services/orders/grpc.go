package main

import (
	"log"
	"net"

	handler "github.com/vivalchemy/kitchen/services/orders/handler/orders"
	"github.com/vivalchemy/kitchen/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{
		addr: addr,
	}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	gRPCServer := grpc.NewServer()
	// register gRPC services
	orderService := service.NewOrderService()
	handler.NewOrdersGRPCHandler(gRPCServer, orderService)

	log.Println("Starting gRPC server on", s.addr)
	return gRPCServer.Serve(lis)
}
