package models

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/goravel/framework/database/orm"
	"gorm.io/gorm"
)

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusPaid      OrderStatus = "paid"
	OrderStatusShipped   OrderStatus = "shipped"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusCancelled OrderStatus = "cancelled"
)

func (s OrderStatus) Valid() bool {
	switch s {
	case OrderStatusPending, OrderStatusPaid, OrderStatusShipped, OrderStatusDelivered, OrderStatusCancelled:
		return true
	default:
		return false
	}
}

// Scan implements sql.Scanner so the value read from the orders.status column
// is converted into an OrderStatus.
func (s *OrderStatus) Scan(value any) error {
	str, ok := value.(string)
	if !ok {
		if b, ok := value.([]byte); ok {
			str = string(b)
		} else {
			return fmt.Errorf("cannot scan %T into OrderStatus", value)
		}
	}

	*s = OrderStatus(str)

	return nil
}

// Value implements driver.Valuer so an OrderStatus is stored as a plain string.
func (s OrderStatus) Value() (driver.Value, error) {
	return string(s), nil
}

type Order struct {
	orm.Model
	ID          uuid.UUID   `json:"id"`
	CustomerID  string      `json:"customer_id"`
	TotalAmount float32     `json:"total_amount"`
	Status      OrderStatus `json:"status"`
	CreatedAt   time.Time   `json:"created_at"`
}

// BeforeCreate generates the primary key before the order is inserted.
func (o *Order) BeforeCreate(tx *gorm.DB) error {
	if o.ID == uuid.Nil {
		o.ID = uuid.New()
	}

	return nil
}
