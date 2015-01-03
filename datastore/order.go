package datastore

import (
	"encoding/json"
	"time"
)

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

func (o *Order) FindAll(wh Where, lim int) ([]Order, error) {
	ords := []Order{}
	rows, err := findAllRows(o, lim, wh)
	// don't forget to close rows
	defer rows.Close()
	// temp store for scan
	var order Order
	var items string
	for rows.Next() {
		var createdAt, updatedAt, closedAt int64
		err := rows.Scan(
			&order.ID,
			&items,
			&order.TableID,
			&order.Cost,
			&order.PercentService,
			&order.Status,
			&order.TotalCost,
			&createdAt,
			&updatedAt,
			&closedAt,
			&order.StaffID,
		)
		if err != nil {
			return ords, err
		}
		order.CreatedAt = time.Unix(createdAt, 0)
		order.UpdatedAt = time.Unix(updatedAt, 0)
		order.ClosedAt = time.Unix(closedAt, 0)

		err = json.Unmarshal([]byte(items), &order.Items)
		if err != nil {
			return ords, err
		}
		ords = append(ords, order)
	}
	return ords, err
}

func (o *Order) FindOne(wh Where) error {
	err := findOne(o, wh, o)
	return err
}
