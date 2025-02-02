package handler

import (
	"context"

	"github.com/vivalchemy/kitchen/services/common/genproto/orders"
	"github.com/vivalchemy/kitchen/services/orders/types"
	"google.golang.org/grpc"
)

type OrdersGRPCHandler struct {
	ordersService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewOrdersGRPCHandler(grpcServer *grpc.Server, ordersService types.OrderService) {
	grpcHandler := &OrdersGRPCHandler{
		ordersService: ordersService,
	}
	orders.RegisterOrderServiceServer(grpcServer, grpcHandler)
}

func (h *OrdersGRPCHandler) GetOrders(ctx context.Context, req *orders.GetOrdersRequest) (*orders.GetOrderResponse, error) {
	o := h.ordersService.GetOrders(ctx)
	res := &orders.GetOrderResponse{
		Orders: o,
	}
	return res, nil
}

func (h *OrdersGRPCHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderId:    1,
		CustomerId: 1,
		ProductId:  1,
		Quantity:   1,
	}

	err := h.ordersService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return &orders.CreateOrderResponse{
		Status: "success",
	}, nil
}
