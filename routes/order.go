package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

	fmt.Fprintf(w, "Delete order  %s", true)
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
