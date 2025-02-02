package service

import (
	"context"
	"log"

	"github.com/vivalchemy/kitchen/services/common/genproto/orders"
)

var orderStore = make([]*orders.Order, 0)

type OrderService struct {
	// store
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	// store
	orderStore = append(orderStore, order)
	log.Println("order created")
	log.Println(orderStore)
	return nil
}

func (s *OrderService) GetOrders(ctx context.Context) []*orders.Order {
	return orderStore
}
