package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"

	"order-service/app/facades"
	"order-service/app/models"
)

type M20260919060632CreateOrdersTable struct{}

// Signature The unique signature for the migration.
func (r *M20260919060632CreateOrdersTable) Signature() string {
	return "20260919060632_create_orders_table"
}

// Up Run the migrations.
func (r *M20260919060632CreateOrdersTable) Up() error {
	if !facades.Schema().HasTable("orders") {
		return facades.Schema().Create("orders", func(table schema.Blueprint) {
			table.Uuid("id")
			table.Primary("id")
			table.String("customer_id")
			table.Decimal("total_amount")
			table.Enum("status", []any{
				string(models.OrderStatusPending),
				string(models.OrderStatusPaid),
				string(models.OrderStatusShipped),
				string(models.OrderStatusDelivered),
				string(models.OrderStatusCancelled),
			}).Default(string(models.OrderStatusPending))
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20260919060632CreateOrdersTable) Down() error {
	return facades.Schema().DropIfExists("orders")
}
