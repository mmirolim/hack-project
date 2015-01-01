package datastore

import (
	"encoding/json"
	"time"
)

type Order struct {
	ID             int
	Items          []Item
	TableID        int
	Cost           int
	PercentService float32
	Status         Status
	TotalCost      int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	ClosedAt       time.Time
	StaffID        int
}

type Orders []Order

//Create an order
func (order Order) CreateOrder() error {
	sql := ` 
			INSERT INTO orders(
							id, 
							items,
							tableID, 
							cost, 
							percentService, 
							status, 
							totalCost, 
							createdAt, 
							updatedAt, 
							closedAt, 
							staffID,
							items) 
			VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
			`
	items, err := json.Marshal(order.Items)
	if err != nil {
		return err
	}
	_, err := DB.Exec(sql,
		order.ID,
		items,
		order.TableID,
		order.Cost,
		order.PercentService,
		order.Status,
		order.TotalCost,
		order.CreatedAt,
		order.UpdatedAt,
		order.ClosedAt,
		order.StaffID,
	)
	if err != nil {
		return err
	}

	return err
}

func (order Order) GetAllOrders() error {
	//var orders Orders
	getOrdersSQL := `
			SELECT	
					id, 
					tableID, 
					cost, 
					percentService, 
					status, 
					totalCost, 
					createdAt, 
					updatedAt, 
					closedAt, 
					staffID
			FROM orders
			`
	rows, err := DB.Query(getOrdersSQL)
	if err != nil {
		return err
	}
	//i := 0
	for rows.Next() {

	}
	return err
}
func (order Order) GetOrder(id int) {
}

func (order Order) UpdateOrder(newOrder Order) {
}

func (order Order) deleteorder() {
}
