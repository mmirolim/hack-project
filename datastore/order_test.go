package datastore

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/mmirolim/hack-project/conf"
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
	file = "foo.db"
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

	//get all orders
	_, err = order.FindAll(Where{Field: "totalCost", Crit: "=", Value: 14000}, 10)
	if err != nil {
		t.Error(err)
	}

	//debug
	//	fmt.Println(">>>>>>>>>>>>>>>")
	//	fmt.Printf("%+v\n", orders)
	//	fmt.Println("<<<<<<<<<<<<<<<")
}

func TestGet(t *testing.T) {
	//get order
	var order Order
	var err error
	err = order.FindOne(Where{Field: "ID", Value: 1, Crit: "="})
	if err != nil {
		t.Error(err)
	}
}

func TestUpdate(t *testing.T) {

	var order Order

	//get order
	err := order.FindOne(Where{Field: "id", Value: 1, Crit: "="})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(">>>>>>>>>>>>> ORDER FINDONE")
	fmt.Printf("%+v\n", order)
	fmt.Println("<<<<<<<<<<<<<")

	order.TableID = 5
	//update order
	err = order.Update()
	if err != nil {
		t.Error(err)
	}

	//check updates
	err = order.FindOne(Where{Field: "id", Value: 1, Crit: "="})
	//	fmt.Println(">>>>>>>>>>>>>")
	//	fmt.Printf("%+v\n", order)
	//	fmt.Println("<<<<<<<<<<<<<")
	if err != nil {
		//		fmt.Println("============")
		//		fmt.Println(order.TableID)
		t.Error(err)
	}
	fmt.Println(">>>>>>>>>>>>> ORDER UPDATE")
	fmt.Printf("%+v\n", order)
	fmt.Println("<<<<<<<<<<<<<")
	//	fmt.Println(">>>>>>>>>>>>>")
	//	fmt.Printf("%+v\n", order)
	//	fmt.Println("<<<<<<<<<<<<<")
}

func TestDelete(t *testing.T) {
	//order is deleted, but order.Get(1) is returns the order
	var order Order
	err := order.FindOne(Where{Field: "ID", Value: 1, Crit: "="})
	order.Delete()
	if err != nil {
		t.Error(order)
	}
	err = order.FindOne(Where{Field: "ID", Value: 2, Crit: "="})
	order.Delete()
	if err != nil {
		t.Error(order)
	}
}
