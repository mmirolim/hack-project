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
