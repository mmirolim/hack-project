package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	ds "github.com/mmirolim/hack-project/datastore"
	"github.com/zenazn/goji/web"
)

func getCatsAll(c web.C, w http.ResponseWriter, r *http.Request) {
	var st ds.Cat
	var err error

	sts, err := st.FindAll(ds.Where{"id", ">", 0}, 0)
	panicOnErr(err)

	catsJSON, err := json.Marshal(sts)
	sendRes(w, err, catsJSON)
}

func getCat(c web.C, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(c.URLParams["id"])
	panicOnErr(err)

	var cat ds.Cat
	err = cat.FindOne(ds.Where{"id", "=", id})
	panicOnErr(err)

	catJSON, err := json.Marshal(cat)
	sendRes(w, err, catJSON)
}

func createCat(c web.C, w http.ResponseWriter, r *http.Request) {
	var st ds.Cat
	err := json.NewDecoder(r.Body).Decode(&st)
	panicOnErr(err)
	err = st.Create()
	sendRes(w, err, nil)
}

func updateCat(c web.C, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(c.URLParams["id"])
	panicOnErr(err)

	var oldCat, cat ds.Cat
	err = oldCat.FindOne(ds.Where{"id", "=", id})
	panicOnErr(err)

	err = json.NewDecoder(r.Body).Decode(&cat)
	panicOnErr(err)

	cat.ID = id
	err = cat.Update()
	panicOnErr(err)

	orderJSON, err := json.Marshal(&cat)
	sendRes(w, err, orderJSON)
}

func deleteCat(c web.C, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(c.URLParams["id"])
	var cat ds.Cat
	err = cat.FindOne(ds.Where{"id", "=", id})
	panicOnErr(err)

	err = cat.Delete()
	panicOnErr(err)

	sendRes(w, err, nil)
}
