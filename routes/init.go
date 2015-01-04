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
	m.Get("/staff", getStaffsAll)
	m.Get("/staff/:id", getStaff)
	m.Post("/staff", createStaff)
	m.Put("/staff/:id", updateStaff)
	return m
}
