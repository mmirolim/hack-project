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

//Helper methods
//get ID
func (order Order) GetID() int {
	return order.ID
}

func (order Order) FieldNames() []string {
	return []string{"items", "tableID", "cost", "percentService", "status", "totalCost", "createdAt", "updatedAt", "closedAt", "staffID"}
}

func (order *Order) SetDefaults() {
	order.CreatedAt = time.Now()
}

func (order Order) Validate() error {
	var err error
	return err
}

func (order Order) TableName() string {
	return "orders"
}

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

//Create an order
func (order Order) Create() error {
	sql := ` 
			INSERT INTO orders(
							items,
							tableID, 
							cost, 
							percentService, 
							status, 
							totalCost, 
							createdAt, 
							updatedAt, 
							closedAt, 
							staffID
							) 
			VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
			`
	items, err := json.Marshal(order.Items)
	if err != nil {
		return err
	}
	_, err = DB.Exec(sql,
		items,
		order.TableID,
		order.Cost,
		order.PercentService,
		order.Status,
		order.TotalCost,
		order.CreatedAt.Unix(),
		order.UpdatedAt.Unix(),
		order.ClosedAt.Unix(),
		order.StaffID,
	)
	return err
}

func (o *Order) FindOne(wh Where) error {
	var createdAt, updatedAt, closedAt int64
	var items string
	err := findOne(o, wh, &o.ID,
		&items,
		&o.TableID,
		&o.Cost,
		&o.PercentService,
		&o.Status,
		&o.TotalCost,
		&createdAt,
		&updatedAt,
		&closedAt,
		&o.StaffID,
	)
	o.CreatedAt = time.Unix(createdAt, 0)
	o.UpdatedAt = time.Unix(updatedAt, 0)
	o.CreatedAt = time.Unix(closedAt, 0)

	return err
}

func (o *Order) FindAll(wh Where, lim int) ([]Order, error) {
	ords := []Order{}
	rows, err := findAllRows(o, lim, wh)
	if err != nil {
		return ords, err
	}
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

func (order *Order) Update() error {
	items, err := json.Marshal(order.Items)
	if err != nil {
		return err
	}

	err = update(order,
		items,
		order.TableID,
		order.Cost,
		order.PercentService,
		order.Status,
		order.TotalCost,
		order.CreatedAt.Unix(),
		order.UpdatedAt.Unix(),
		order.ClosedAt.Unix(),
		order.StaffID,
	)
	return err
}

func (order *Order) Delete() error {
	err := del(order)
	return err
}
