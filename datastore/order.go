package datastore

import "time"

type Order struct {
	ID             int       `db:"id"`
	Items          []Item    `db:"items"`
	TableID        int       `db:"tableID"`
	Cost           int       `db:"cost"`
	PercentService float32   `db:"percentService"`
	Status         Status    `db:"status"`
	TotalCost      int       `db:"totalCost"`
	CreatedAt      time.Time `db:"createdAt"`
	UpdatedAt      time.Time `db:"updatedAt"`
	ClosedAt       time.Time `db:"closedAt"`
	StaffID        int       `db:"staffID"`
}

// models table name
func (o Order) TableName() string {
	return "orders"
}

// create table query
func (o Order) createTableQuery() string {
	q := "CREATE TABLE IF NOT EXISTS " + o.TableName()
	q += ` ( 
		id INTEGER PRIMARY	KEY AUTOINCREMENT, 
		items TEXT, 
		tableID INTEGER , 
		cost INTEGER, 
		percentService REAL, 
		status INTEGER, 
		totalCost INTEGER, 
		createdAt INTEGER, 
		updatedAt INTEGER, 
		closedAt INTEGER, 
		staffID INTEGER
	)
	`
	return q
}

// defaults to set
func (o *Order) SetDefaults() {
	o.CreatedAt = time.Now()
}

// validation logic for items
func (o Order) Validate() error {
	var err error
	return err
}
