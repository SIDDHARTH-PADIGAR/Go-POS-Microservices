package handler

import (
	util "go-pos-ms/services/common"
	"go-pos-ms/services/common/genproto/orders"
	"go-pos-ms/services/orders/types"
	"net/http"
)

type OrdersHttpHandler struct {
	orderService types.OrderService
}

func NewhttpOrdersHandler(orderService types.OrderService) *OrdersHttpHandler {
	handler := &OrdersHttpHandler{
		orderService: orderService,
	}

	return handler
}

func (h *OrdersHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	// Implement the CreateOrder handler
	var req orders.CreateOrderRequest
	err := util.ParseJSON(r, &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	order := &orders.Order{
		OrderID:    42,
		CustomerID: req.GetCustomerID(),
		ProductID:  req.GetProductID(),
		Quantity:   req.GetQuantity(),
	}

	err = h.orderService.CreateOrder(r.Context(), order)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
	}

	res := &orders.CreateOrderResponse{
		Status: "Order created successfully",
	}

	util.WriteJSON(w, http.StatusOK, res)
}
