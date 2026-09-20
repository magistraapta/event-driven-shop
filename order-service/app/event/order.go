package event

import (
	"time"

	"order-service/app/models"
)

type OrderEvent struct {
	EventID   string       `json:"evnet_id"`
	EventType string       `json:"event_type"`
	Order     models.Order `json:"order"`
	CreatedAt time.Time    `json:"created_at"`
}
