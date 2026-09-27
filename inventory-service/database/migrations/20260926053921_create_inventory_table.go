package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"

	"inventory-service/app/facades"
)

type M20260926053921CreateInventoryTable struct{}

// Signature The unique signature for the migration.
func (r *M20260926053921CreateInventoryTable) Signature() string {
	return "20260926053921_create_inventory_table"
}

// Up Run the migrations.
func (r *M20260926053921CreateInventoryTable) Up() error {
	if !facades.Schema().HasTable("inventory") {
		return facades.Schema().Create("inventory", func(table schema.Blueprint) {
			table.Uuid("id")
			table.Primary("id")
			table.Uuid("product_id")
			table.Integer("stock")
			table.Integer("reserved_stock")
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20260926053921CreateInventoryTable) Down() error {
	return facades.Schema().DropIfExists("inventory")
}
