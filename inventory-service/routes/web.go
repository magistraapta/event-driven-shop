package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/support"

	"inventory-service/app/http/controllers"
	"inventory-service/app/facades"
)

func Web() {
	facades.Route().Get("/", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("welcome.tmpl", map[string]any{
			"version": support.Version,
		})
	})

	facades.Route().Static("public", "./public")

	inventoryController := controllers.NewInventoryController()
	facades.Route().Get("/", inventoryController.Index)
}
