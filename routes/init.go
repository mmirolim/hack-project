package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/zenazn/goji/web"

	ds "github.com/mmirolim/hack-project/datastore"
)

const (
	OK       = "OK"
	SUCCESS  = "SUCCESS"
	FAILURE  = "FAILURE"
	NOTFOUND = "NOT FOUND"
)

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
		url := fmt.Sprintf("%s", r.URL)
		// @TODO temp fix to content type
		if !strings.Contains(url, "/assets/") {
			w.Header().Set("Content-Type", "application/json")
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

type Reply struct {
	StatusCode int    `json:"status"`
	Msg        string `json:"msg"`
}

func (p *Reply) Set(status int, msg string) {
	p.StatusCode = status
	p.Msg = msg
}

// @todo refactor should be configurabe from header
// now just json
func replyJson(w http.ResponseWriter, data interface{}) {
	var res string
	b, err := json.Marshal(data)
	if err != nil {
		res = err.Error()
	} else {
		res = string(b)
	}
	switch v := data.(type) {
	case Reply:
		w.WriteHeader(v.StatusCode)
	}
	fmt.Fprintf(w, res)
}
