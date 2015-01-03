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

func Initialize(status <-chan ds.Status) *web.Mux {
	m := web.New()
	// show default html
	m.Get("/", http.FileServer(http.Dir("assets")))
	// serve other static files
	m.Get("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	// orders
	m.Get("/orders", getOrdersAll)
	m.Get("/orders/:id", getOrder)
	m.Post("/orders", createOrder)
	m.Put("/orders/:id", updateOrder)
	m.Delete("/orders/:id", deleteOrder)
	// tables
	m.Get("/tables", getTablesAll)
	m.Get("/tables/:id", getTable)
	m.Post("/tables", createTable)
	m.Put("/tables/:id", updateTable)
	m.Delete("/tables/:id", deleteTable)
	// users
	m.Get("/users", getUsersAll)
	m.Get("/users/:id", getUser)
	m.Post("/users", createUser)
	m.Put("/users/:id", updateUser)
	return m
}

func getOrdersAll(c web.C, w http.ResponseWriter, r *http.Request) {
	var order ds.Order
	var err error
	orders, err := order.FindAll(ds.Where{}, 0)
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

func getTablesAll(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get all tables %s", "Tables")
}

func getTable(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get table %s", c.URLParams["id"])
}

func createTable(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create table %s", c.URLParams["id"])
}

func updateTable(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update table  %s", c.URLParams["id"])
}

func deleteTable(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete table  %s", c.URLParams["id"])
}
func getUsersAll(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get all users %s", "Users")
}

func getUser(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get user %s", c.URLParams["id"])
}

func createUser(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create user %s", c.URLParams["id"])
}

func updateUser(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update user %s", c.URLParams["id"])
}
