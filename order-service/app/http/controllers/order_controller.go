package controllers

import (
	"order-service/app/facades"
	"order-service/app/http/requests"
	"order-service/app/messaging"
	"order-service/app/models"

	"github.com/goravel/framework/contracts/http"
)

type OrderController struct {
	// Dependent services
}

func NewOrderController() *OrderController {
	return &OrderController{}
}

func (r *OrderController) Index(ctx http.Context) http.Response {
	return ctx.Response().Status(http.StatusOK).Json(http.Json{
		"message": "work in progress",
	})
}

func (r *OrderController) Store(ctx http.Context) http.Response {
	var orderRequests requests.StoreOrderRequest

	order := models.Order{
		CustomerID:  orderRequests.CustomerID,
		TotalAmount: orderRequests.TotalAmount,
		Status:      models.OrderStatusPending,
	}

	if err := facades.Orm().Query().Create(&order); err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{
			"error": err.Error(),
		})
	}

	if err := messaging.OrderPublisher().Publish(ctx.Context(), "order.created", order); err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{
			"error": err.Error(),
		})
	}

	return ctx.Response().Status(http.StatusCreated).Json(http.Json{
		"response_data": order,
	})
}
