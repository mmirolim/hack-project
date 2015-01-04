package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	ds "github.com/mmirolim/hack-project/datastore"
	"github.com/zenazn/goji/web"
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

	order.SetDefaults()
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
	w.Header().Set("Content-Type", "text/json")
	if err != nil {
		panic(err)
		res := "{'result': 'failure'}"
		result, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, string(result))
	}

	res := "{'result': 'success'}"
	result, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(result))
}
func updateOrder(c web.C, w http.ResponseWriter, r *http.Request) {
	var order ds.Order
	id, err := strconv.Atoi(c.URLParams["id"])
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(r.Body)
	log.Println(r.Body)
	err = decoder.Decode(&order)
	if err != nil {
		panic(err)
	}
	order.ID = id
	err = order.Update()
	if err != nil {
		panic(err)
	}
	orderJSON, _ := json.Marshal(&order)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "text/json")
	fmt.Fprintf(w, string(orderJSON))
}

func statusOrder(c web.C, w http.ResponseWriter, r *http.Request) {
	var order ds.Order
	status, err := strconv.Atoi(c.URLParams["status"])
	fmt.Println(status)
	if err != nil {
		panic(err)
	}
	orders, err := order.FindAll(ds.Where{
		Field: "status",
		Crit:  " = ",
		Value: status}, 10)

	fmt.Println(orders)
	if err != nil {
		panic(err)
	}
	ordersJSON, err := json.Marshal(orders)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(ordersJSON))
}

func todayOrder(c web.C, w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	t := time.Date(now.Year(), now.Month(), now.Day(), 00, 0, 0, 0, time.UTC)
	today := int(t.Unix())

	var order ds.Order
	orders, err := order.FindAll(ds.Where{Field: "createdAt", Crit: ">", Value: today}, 10)
	ordersJSON, err := json.Marshal(orders)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(ordersJSON))
}

func activeOrder(c web.C, w http.ResponseWriter, r *http.Request) {
	_, err := strconv.Atoi(c.URLParams["alias"])
	panicOnErr(err)

	var order ds.Order
	orders, err := order.FindAll(ds.Where{Field: "status", Crit: "!=", Value: ds.StatusPaid}, 0)

	ordersJSON, err := json.Marshal(orders)
	panicOnErr(err)

	w.Header().Set("Content-type", "application/json")

	fmt.Fprintf(w, string(ordersJSON))
}
