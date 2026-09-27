package main

import (
	"inventory-service/bootstrap"
)

func main() {
	app := bootstrap.Boot()

	app.Start()
}
