package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"

	ds "github.com/mmirolim/hack-project/datastore"
)

func getNotsAll(c web.C, w http.ResponseWriter, r *http.Request) {
	var not ds.Nots
	nots, err := not.FindAll(ds.Where{"id", ">", 0}, 0)
	panicOnErr(err)
	jsn, err := json.Marshal(nots)
	fmt.Fprintf(w, string(jsn))
}

func getNots(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get table %s", c.URLParams["id"])
}

func createNots(c web.C, w http.ResponseWriter, r *http.Request) {
	var not ds.Nots
	err := json.NewDecoder(r.Body).Decode(&not)
	panicOnErr(err)
	err = not.Create()
	panicOnErr(err)
	fmt.Fprintf(w, "success")
}

func updateNots(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update table  %s", c.URLParams["id"])
}

func deleteNots(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete table  %s", c.URLParams["id"])
}

func getTableNots(c web.C, w http.ResponseWriter, r *http.Request) {
	var not ds.Nots
	alias := c.URLParams["alias"]
	nots, err := not.FindAll(ds.Where{"tableID", "=", alias}, 0)
	panicOnErr(err)
	dat, err := json.Marshal(nots)
	panicOnErr(err)
	fmt.Fprintf(w, string(dat))
}
