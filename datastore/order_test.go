package datastore

import (
	"log"
	"testing"
)

func TestCreateOrder(t *testing.T) {
	items := []Item{
		{
			ID:      1,
			Name:    "Osh",
			Desc:    "Very tasty",
			Img:     "/url/",
			Serving: 123.3,
			Cost:    12000,
			Unit:    "portion",
			Status:  1,
		},

		{
			ID:      2,
			Name:    "Osh",
			Desc:    "Very tasty",
			Img:     "/url/",
			Serving: 123.3,
			Cost:    12000,
			Unit:    "portion",
			Status:  1,
		},
	}
	var order Order
	var status StatusIssued

	order.Items = items
	order.TableID = 2
	order.Cost = 12000
	order.PercentService = 12.34
	order.Status = status
	order.TotalCost = 14000
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	order.ClosedAt = time.Now()
	StaffID = 1

	mockConf := `
[ds]
	[ds.sqlite]
	name = "sqlite3"
	file = "../foo.db"
[srv]
	port = "3000"
`

	f := strings.NewReader(mockConf)

	App, _ := conf.Read(f)

	db, err := Initialize(App.DS)
	if err != nil {
		t.Error(err)
	}

	err := CreateOrder(order)
	if err != nil {
		t.Error(err)
	}
}
