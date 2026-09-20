package main

import (
	"order-service/bootstrap"
)

func main() {
	app := bootstrap.Boot()

	app.Start()
}
