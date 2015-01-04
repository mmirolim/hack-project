package routes

import (
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
	m.Delete("/orders/:id", deleteOrder)
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
