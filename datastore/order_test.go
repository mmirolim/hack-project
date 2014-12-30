package datastore

import (
	"fmt"
	"github.com/mmirolim/hack-project/conf"
	"log"
	"strings"
	"testing"
	"time"
)

func TestOrder(t *testing.T) {

	var osh Item
	osh.Name = "Osh"
	osh.Desc = "very tasty"
	osh.Img = "/url/url/url"
	osh.Serving = 12.34
	osh.Cost = 10000
	osh.Unit = "portion"
	osh.Status = 1

	var order Order
	var items = make([]Item, 1)
	items[0] = osh
	order.Items = items
	order.TableID = 1
	order.Cost = 10000
	order.PercentService = 12
	order.Status = 2
	order.TotalCost = 11200
	order.StaffID = 3
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	order.ClosedAt = time.Now()

	//DB initialize
	mockConf := `
[ds]
	[ds.sqlite]
	name = "sqlite3"
	file = "../foo.db"
[srv]
	port = "3000"
`

	f := strings.NewReader(mockConf)

	var App conf.App

	App, err := conf.Read(f)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}

	err = Initialize(App.DS)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	fmt.Printf("%+v", DB)

	order.CreateOrder()

	/*getOrder test
	var getOrder Order
	getOrder.GetOrder(1)
	fmt.Printf("%+v", getOrder)

	var allOrders Order
	allOrders.GetAllOrders()
	fmt.Printf("%+v", allOrders)
	*/
}
