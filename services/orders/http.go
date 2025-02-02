package main

import (
	"log"
	"net/http"

	handler "github.com/vivalchemy/kitchen/services/orders/handler/orders"
	"github.com/vivalchemy/kitchen/services/orders/service"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()
	orderService := service.NewOrderService()

	orderHandler := handler.NewOrdersHttpHandler(orderService)
	orderHandler.RegisterRouter(router)

	log.Println("Starting HTTP server on", s.addr)

	return http.ListenAndServe(s.addr, router)

}
