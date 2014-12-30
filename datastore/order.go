package datastore

import (
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

func (order *Order) CreateOrder() {
	if order.ID == 0 {
		DB.Create(order)
		DB.Save(order)
	}
}

func (order Order) GetAllOrders() {
	DB.Find(&order)
}

func (order Order) GetOrder(id int) {
	DB.First(&order, id)
}

func (order Order) UpdateOrder(newOrder Order) {
	DB.First(&order, order.ID)
	order = newOrder
	DB.Save(&order)
}

func (order Order) DeleteOrder() {
	DB.Delete(&order)
}
