package datastore

import (
	"fmt"
	"github.com/mmirolim/hack-project/conf"
	"strings"
	"testing"
	"time"
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
	status := StatusIssued

	order.Items = items
	order.TableID = 2
	order.Cost = 12000
	order.PercentService = 12.34
	order.Status = status
	order.TotalCost = 14000
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	order.ClosedAt = time.Now()
	order.StaffID = 1

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

	_, err := Initialize(App.DS)
	if err != nil {
		t.Error(err)
	}

	err = order.Create()
	if err != nil {
		t.Error(err)
	}

	var orders Orders
	orders, err = orders.GetAll()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", orders)

	order, err = order.Get(3)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", order)
}
