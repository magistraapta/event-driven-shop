package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type StoreOrderRequest struct {
	CustomerID  string  `form:"customer_id" json:"customer_id"`
	TotalAmount float32 `form:"total_amount" json:"total_amount"`
}

func (r *StoreOrderRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreOrderRequest) Filters(ctx http.Context) map[string]any {
	return map[string]any{}
}

func (r *StoreOrderRequest) Rules(ctx http.Context) map[string]any {
	return map[string]any{
		"customer_id":  "required|string",
		"total_amount": "required|numeric|min:0",
	}
}

func (r *StoreOrderRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *StoreOrderRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *StoreOrderRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
