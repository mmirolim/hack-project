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
	panicOnErr(err)
	fmt.Printf("%+v\n", st)
	err = st.Create()
	panicOnErr(err)
	fmt.Fprintf(w, "success")
}

func updateCat(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update user %s", c.URLParams["id"])
}

func deleteCat(c web.C, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(c.URLParams["id"])
	var cat ds.Cat
	err = cat.FindOne(ds.Where{"id", "=", id})
	panicOnErr(err)
	err = cat.Delete()
	panicOnErr(err)
	w.Header().Set("Content-Type", "text/json")
	if err != nil {
		panic(err)
		res := "{'result': 'failure'}"
		result, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, string(result))
	}
}
