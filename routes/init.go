package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"

	ds "github.com/mmirolim/hack-project/datastore"
)

type opStat struct {
	Status string `json:"status"`
}

func Initialize(status <-chan ds.Status) *web.Mux {
	// WARNING more specific routes should be first then more general
	m := web.New()
	// show default html
	m.Get("/", http.FileServer(http.Dir("assets")))
	// serve other static files
	m.Get("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	// orders
	m.Get("/orders", getOrdersAll)
	m.Post("/orders", createOrder)
	m.Get("/orders/active/table/:alias", activeOrder)
	m.Get("/orders/status/:status", statusOrder)
	m.Get("/orders/today", todayOrder)
	m.Get("/orders/:id", getOrder)
	m.Delete("/orders/:id", deleteOrder)
	m.Put("/orders/:id", updateOrder)
	// tables
	m.Get("/tables", getTablesAll)
	m.Get("/tables/:id", getTable)
	m.Post("/tables", createTable)
	m.Put("/tables/:id", updateTable)
	m.Delete("/tables/:id", deleteTable)
	// staff
	m.Get("/staff", getStaffAll)
	m.Get("/staff/:id", getStaff)
	m.Post("/staff", createStaff)
	m.Put("/staff/:id", updateStaff)

	//categories
	m.Get("/categories", getCatsAll)
	m.Get("/categories/:id", getCat)
	m.Post("/categories", createCat)
	m.Put("/categories/:id", updateCat)
	m.Delete("/categories/:id", deleteCat)

	//items
	m.Get("/items", getItemsAll)
	m.Get("/items/:id", getItem)
	m.Post("/items", createItem)
	m.Put("/items/:id", updateItem)
	m.Delete("/items/:id", deleteItem)

	// notifications
	m.Get("/notifications/tables/:alias", getTableNots)
	return m
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

// set response to json format
func JSON(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func sendRes(w http.ResponseWriter, err error, object []byte) {
	if err != nil {
		fmt.Printf("%+v\n", err)
		r, err := json.Marshal(opStat{"failure"})
		panicOnErr(err)
		fmt.Fprintf(w, string(r))
	} else {
		if len(object) == 0 {
			r, err := json.Marshal(opStat{"success"})
			panicOnErr(err)
			fmt.Fprintf(w, string(r))
		} else {
			fmt.Fprintf(w, string(object))
		}
	}
}
