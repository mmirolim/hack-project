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

func (order Order) CreateOrder() error {
	db.NewRecord(order)
	db.Create(&order)
}
