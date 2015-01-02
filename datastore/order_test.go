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
}

func TestGetALL(t *testing.T) {
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

	err := order.Create()
	if err != nil {
		t.Error(err)
	}
	//create array of orders
	var orders Orders

	//get all orders
	orders, err = orders.GetAll()
	if err != nil {
		t.Error(err)
	}

	//debug
	fmt.Printf("%+v\n", orders)
}

func TestGet(t *testing.T) {
	//get order
	var order Order
	var err error
	order, err = order.Get(1)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", order)
}

func TestUpdate(t *testing.T) {
	status := StatusDelivered
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
	var newOrder Order

	newOrder.Items = items
	newOrder.TableID = 5
	newOrder.Cost = 15000
	newOrder.PercentService = 34.56
	newOrder.Status = status
	newOrder.TotalCost = 18000
	newOrder.CreatedAt = time.Now()
	newOrder.UpdatedAt = time.Now()
	newOrder.ClosedAt = time.Now()
	newOrder.StaffID = 1

	//get order
	order, err := order.Get(1)
	if err != nil {
		t.Error(err)
	}

	//update order
	err = order.Update(newOrder)
	if err != nil {
		t.Error(err)
	}

	//check updates
	order, err = order.Get(1)
	if order.TableID != 5 {
		t.Error(err)
	}
	fmt.Printf("%+v\n", order)
}

func TestDelete(t *testing.T) {
	//order is deleted, but order.Get(1) is returns the order
	var order Order
	order, err := order.Get(1)
	order.Delete()
	if err != nil {
		t.Error(order)
	}
}

func TestDropTable(t *testing.T) {
	err := DropTable()
	if err != nil {
		fmt.Println(err)
	}
}
