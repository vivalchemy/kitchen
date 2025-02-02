package handler

import (
	"encoding/json"
	"net/http"

	"github.com/vivalchemy/kitchen/services/common/genproto/orders"
	"github.com/vivalchemy/kitchen/services/orders/types"
)

type OrdersHttpHandler struct {
	ordersService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewOrdersHttpHandler(ordersService types.OrderService) *OrdersHttpHandler {
	httpHandler := &OrdersHttpHandler{
		ordersService: ordersService,
	}
	return httpHandler
}

func (h *OrdersHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	order := &orders.Order{
		OrderId:    42,
		CustomerId: req.GetCustomerId(),
		ProductId:  req.GetProductId(),
		Quantity:   req.GetQuantity(),
	}

	err = h.ordersService.CreateOrder(r.Context(), order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := orders.CreateOrderResponse{
		Status: "success",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
