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
	if err != nil {
		return err
	}

	return err
}

func (orders Orders) GetAll() (Orders, error) {
	//var orders Orders
	var count int
	getOrdersCountSQL := "SELECT COUNT() FROM orders"

	rows, err := DB.Query(getOrdersCountSQL)
	for rows.Next() {
		rows.Scan(&count)
	}
	orders = make(Orders, count)
	//fmt.Printf("%+v", orders)
	getOrdersSQL := `
			SELECT	
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
					staffID
			FROM orders
			`
	rows, _ = DB.Query(getOrdersSQL)
	if err != nil {
		return nil, err
	}

	i := 0
	var items string
	for rows.Next() {
		var createdAt, updatedAt, closedAt int64
		if err := rows.Scan(
			&orders[i].ID,
			&items,
			&orders[i].TableID,
			&orders[i].Cost,
			&orders[i].PercentService,
			&orders[i].Status,
			&orders[i].TotalCost,
			&createdAt,
			&updatedAt,
			&closedAt,
			&orders[i].StaffID,
		); err != nil {
			return nil, err
		}
		orders[i].CreatedAt = time.Unix(createdAt, 0)
		orders[i].UpdatedAt = time.Unix(updatedAt, 0)
		orders[i].ClosedAt = time.Unix(closedAt, 0)

		json.Unmarshal([]byte(items), &orders[i].Items)
		i += 1
	}

	return orders, err
}

func (order Order) Get(id int) (Order, error) {
	getOrderByIDSQL := `
			SELECT	
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
					staffID
			FROM orders
			WHERE id = ?
			`

	rows, err := DB.Query(getOrderByIDSQL, id)
	if err != nil {
		return order, err
	}

	var items string
	for rows.Next() {
		var createdAt, updatedAt, closedAt int64
		if err := rows.Scan(
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
		); err != nil {
			return order, err
		}
		order.CreatedAt = time.Unix(createdAt, 0)
		order.UpdatedAt = time.Unix(updatedAt, 0)
		order.ClosedAt = time.Unix(closedAt, 0)

		json.Unmarshal([]byte(items), &order.Items)
	}
	return order, err
}

func (order Order) Update(newOrder Order) error {
	updateOrderSQL := ` UPDATE orders
						SET	items = ?,
						tableID = ?,
						percentService =?,
						status = ?,
						totalCost = ?,
						createdAt=?,
						updatedAt = ?,
						closedAt=?,
						staffID=?
						WHERE id = ?
						`
	items, err := json.Marshal(newOrder.Items)
	if err != nil {
		return err
	}

	_, err = DB.Exec(updateOrderSQL,
		items,
		newOrder.TableID,
		newOrder.PercentService,
		newOrder.Status,
		newOrder.TotalCost,
		newOrder.CreatedAt.Unix(),
		newOrder.UpdatedAt.Unix(),
		newOrder.ClosedAt.Unix(),
		newOrder.StaffID,
		order.ID,
	)
	if err != nil {
		return err
	}
	return err

}

func (order Order) Delete() error {
	deleteOrderSQL := "DELETE FROM orders WHERE id = ?"
	_, err := DB.Exec(deleteOrderSQL, order.ID)
	if err != nil {
		return err
	}

	return err
}
