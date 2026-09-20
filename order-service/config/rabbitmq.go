package config

import "order-service/app/facades"

func init() {
	config := facades.Config()

	config.Add("rabbitmq", map[string]any{
		"url":      config.Env("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"),
		"exchange": config.Env("RABBITMQ_EXCHANGE", "order_events"),
	})
}
