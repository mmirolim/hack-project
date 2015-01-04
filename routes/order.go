package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/zenazn/goji/web"

	ds "github.com/mmirolim/hack-project/datastore"
)

func getOrdersAll(c web.C, w http.ResponseWriter, r *http.Request) {
	var order ds.Order
	var err error
	orders, err := order.FindAll(ds.Where{Field: "id", Crit: ">", Value: "0"}, 10)
	if err != nil {
		panic(err)
	}
	ordersJSON, err := json.Marshal(orders)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "text/json")
	fmt.Fprintf(w, string(ordersJSON))
}

func getOrder(c web.C, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(c.URLParams["id"])
	if err != nil {
		panic(err)
	}
	var order ds.Order
	err = order.FindOne(ds.Where{"id", "=", id})
	log.Println(order)
	if err != nil {
		panic(err)
	}
	orderJSON, _ := json.Marshal(order)
	w.Header().Set("Content-Type", "text/json")
	fmt.Fprintf(w, string(orderJSON))
}

func createOrder(c web.C, w http.ResponseWriter, r *http.Request) {
	var order ds.Order

	decoder := json.NewDecoder(r.Body)
	log.Println(r.Body)
	err := decoder.Decode(&order)
	if err != nil {
		panic(err)
	}
	err = order.Create()
	if err != nil {
		panic(err)
	}
	orderJSON, err := json.Marshal(order)
	if err != nil {
		panic(err)
	}
	// create notifications
	var not ds.Nots
	not.TableID = order.TableID
	not.OrderID = order.ID
	not.Msg = "You order status"
	not.Status = int(ds.StatusIssued)
	not.StaffID = order.StaffID
	not.CreatedAt = time.Now()
	not.UpdatedAt = time.Now()
	err = not.Create()
	panicOnErr(err)
	fmt.Fprintf(w, string(orderJSON))
}

func deleteOrder(c web.C, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(c.URLParams["id"])
	if err != nil {
		panic(err)
	}
	var order ds.Order
	err = order.FindOne(ds.Where{"id", "=", id})
	if err != nil {
		panic(err)
	}
	err = order.Delete()
	if err != nil {
		panic(err)
		fmt.Fprintf(w, "Delete order  %s", false)
	}
	// @todo don't forget optimize
	// del all notifications for order
	var not ds.Nots
	err = not.FindOne(ds.Where{"orderID", "=", order.ID})
	panicOnErr(err)
	err = not.Delete()
	panicOnErr(err)
	fmt.Fprintf(w, "Delete order  %s", true)
}
func updateOrder(c web.C, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(c.URLParams["id"])
	panicOnErr(err)
	var oldOrd, order ds.Order
	err = oldOrd.FindOne(ds.Where{"id", "=", id})
	panicOnErr(err)
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(r.Body).Decode(&order)
	panicOnErr(err)
	order.ID = id
	err = order.Update()
	panicOnErr(err)
	if oldOrd.Status != order.Status {
		// if no err update status of notification
		var not ds.Nots
		err = not.FindOne(ds.Where{"orderID", "=", order.ID})
		panicOnErr(err)
		not.Status = int(order.Status)
		not.UpdatedAt = time.Now()
		err = not.Update()
		panicOnErr(err)
	}
	orderJSON, _ := json.Marshal(&order)
	panicOnErr(err)
	fmt.Fprintf(w, string(orderJSON))
}
