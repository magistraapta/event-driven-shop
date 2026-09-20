package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"

	"order-service/app/facades"
)

const ExchangeName = "order_events"

type Publisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

type Event struct {
	EventID   string    `json:"event_id"`
	EventType string    `json:"event_type"`
	OccuredAt time.Time `json:"occured_at"`
	Payload   any       `json:"payload"`
}

func NewPublisher() (*Publisher, error) {
	url := facades.Config().GetString("rabbitmq.url")

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("amqp dial error: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("rabbitmq channel error: %w", err)
	}

	if err := ch.ExchangeDeclare(
		ExchangeName,
		"topic",
		true,
		false,
		false,
		false,
		nil); err != nil {
		return nil, fmt.Errorf("rabbitmq exchange declare %w", err)
	}

	return &Publisher{conn: conn, channel: ch}, nil
}

func (r *Publisher) Publish(ctx context.Context, routingKey string, payload any) error {
	event := Event{
		EventID:   uuid.NewString(),
		EventType: routingKey,
		OccuredAt: time.Now().UTC(),
		Payload:   payload,
	}

	body, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("marshal event %w", err)
	}

	return r.channel.PublishWithContext(ctx,
		ExchangeName, routingKey, false, false, amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Body:         body,
		})
}

func (r *Publisher) Close() {
	if r.channel != nil {
		r.channel.Close()
	}

	if r.conn != nil {
		r.conn.Close()
	}
}

var (
	orderEventPublisher     *Publisher
	orderEventPublisherOnce sync.Once
)

// OrderPublisher returns the shared publisher instance used across the app,
// connecting to RabbitMQ lazily on first use. It must not be called from a
// package init() — the app (facades.App()) is only ready once bootstrap.Boot()
// has run inside main().
func OrderPublisher() *Publisher {
	orderEventPublisherOnce.Do(func() {
		p, err := NewPublisher()
		if err != nil {
			facades.Log().Fatalf("Error initializing OrderEventPublisher: %v", err)
		}

		orderEventPublisher = p
	})

	return orderEventPublisher
}
