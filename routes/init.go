package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"

	ds "github.com/mmirolim/hack-project/datastore"
)

const (
	OK      = "OK"
	SUCCESS = "SUCCESS"
	FAILURE = "FAILURE"
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
	m.Get("/categories", getCatAll)
	m.Get("/categories/:id", getCat)
	m.Post("/categories", createCat)
	m.Put("/categories/:id", updateCat)
	m.Delete("/categories/:id", deleteCat)
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
func reply(w http.ResponseWriter, status int, msg string) {
	rep := Reply{status, msg}
	var res string
	b, err := json.Marshal(rep)
	if err != nil {
		res = err.Error()
	} else {
		res = string(b)
	}
	fmt.Fprintf(w, res)
}

func replyOnErr(w http.ResponseWriter, status int, err error) {
	if err != nil {
		reply(w, status, err.Error())
	}
}
