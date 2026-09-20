package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/support"

	"order-service/app/facades"
	"order-service/app/http/controllers"
)

func Web() {
	facades.Route().Get("/", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("welcome.tmpl", map[string]any{
			"version": support.Version,
		})
	})

	facades.Route().Static("public", "./public")

	orderController := controllers.NewOrderController()
	facades.Route().Get("/orders", orderController.Index)
	facades.Route().Post("/orders", orderController.Store)
}
