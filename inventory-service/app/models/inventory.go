package models

import (
	"github.com/google/uuid"
	"github.com/goravel/framework/database/orm"
)

type Inventory struct {
	orm.Model
	ID uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"product_id"`
	Stock int `json:"stock"`
	ReservedStock int `json:"reserved_stock"`
}
