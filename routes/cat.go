package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	ds "github.com/mmirolim/hack-project/datastore"
	"github.com/zenazn/goji/web"
)

func getCatAll(c web.C, w http.ResponseWriter, r *http.Request) {
	var st ds.Cat
	var err error
	sts, err := st.FindAll(ds.Where{"id", ">", 0}, 0)
	panicOnErr(err)
	jsn, err := json.Marshal(sts)
	fmt.Fprintf(w, string(jsn))
}

func getCat(c web.C, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(c.URLParams["id"])
	panicOnErr(err)
	var cat ds.Cat
	err = cat.FindOne(ds.Where{"id", "=", id})
	panicOnErr(err)

	catJSON, _ := json.Marshal(cat)

	fmt.Fprintf(w, string(catJSON))
}

func createCat(c web.C, w http.ResponseWriter, r *http.Request) {
	var st ds.Cat
	err := json.NewDecoder(r.Body).Decode(&st)
	replyOnErr(w, 400, err)

	err = st.Create()
	replyOnErr(w, 500, err)

	reply(w, 200, OK)
}

func updateCat(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update user %s", c.URLParams["id"])
}

func deleteCat(c web.C, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(c.URLParams["id"])
	var cat ds.Cat
	err = cat.FindOne(ds.Where{"id", "=", id})
	replyOnErr(w, 500, err)

	err = cat.Delete()
	replyOnErr(w, 500, err)

	reply(w, 200, OK)
}
