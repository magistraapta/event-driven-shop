package messaging

import (
	"context"
	"fmt"

	"github.com/rabbitmq/amqp091-go"

	"inventory-service/app/facades"
)


type Consumer struct {
	conn *amqp091.Connection
	channel *amqp091.Channel
}

func NewConsumer(queueName, exchangeName string, routingKeys []string) (*Consumer, error) {
	url := facades.Config().GetString("rabbitmq.url")

	conn, err := amqp091.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("amqp dial error %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("rabbitmq channel error")
	}

	if _, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return nil, fmt.Errorf("rabbitmq queue declare: %w", err)
	}

	for _, key := range routingKeys {
		if err := ch.QueueBind(queueName, key, exchangeName, false, nil); err != nil {
			return nil, fmt.Errorf("rabbitmq queue bind: %w", err)
		}
	}

	if err := ch.Qos(1, 0, false); err != nil {
		return nil, fmt.Errorf("rabbitmq qos: %w", err)
	}

	return &Consumer{conn: conn, channel: ch}, nil
}

func (c *Consumer) Consume(ctx context.Context, queueName string ,routingKeys []string, handler func(amqp091.Delivery) error) error {
	deliveries, err := c.channel.Consume(queueName,"", false, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("rabbitmq consume: %w", err)
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case d, ok := <- deliveries:
			if !ok {
				return nil
			}
			if err := handler(d); err != nil {
				facades.Log().Errorf("event handler error: %w", err)
				d.Nack(false, false)
				continue
			}
			d.Ack(false)
		}
	}
} 