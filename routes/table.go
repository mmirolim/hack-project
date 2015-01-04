package routes

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"
)

func getTablesAll(c web.C, w http.ResponseWriter, r *http.Request) {
	var tbl ds.Table
	tbls, err := tbl.FindAll(ds.Where{"id", ">", 0}, 0)
	panicOnErr(err)
	jsn, err := json.Marshal(tbls)
	fmt.Fprintf(w, string(jsn))
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
