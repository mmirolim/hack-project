package routes

import (
	"fmt"
	"net/http"

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
	// tables
	m.Get("/tables", getTablesAll)
	m.Get("/tables/:id", getTable)
	m.Post("/tables", createTable)
	m.Put("/tables/:id", updateTable)
	// users
	m.Get("/users", getUsersAll)
	m.Get("/users/:id", getUser)
	m.Post("/users", createUser)
	m.Put("/users/:id", updateUser)
	return m
}

func getOrdersAll(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All orders %s", "ORDERS")
}

func getOrder(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Order %s", c.URLParams["id"])
}

func createOrder(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create order %s", c.URLParams["id"])
}

func updateOrder(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update order  %s", c.URLParams["id"])
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
