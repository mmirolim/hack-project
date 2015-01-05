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

	catsJSON, err := json.Marshal(sts)
	if err != nil {
		res := "{'result': 'failure'}"
		result, err := json.Marshal(res)
		fmt.Fprintf(w, string(result))
		panicOnErr(err)
	} else {
		fmt.Fprintf(w, string(catsJSON))
	}
}

func getCat(c web.C, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(c.URLParams["id"])
	panicOnErr(err)

	var cat ds.Cat
	err = cat.FindOne(ds.Where{"id", "=", id})
	panicOnErr(err)

	catJSON, err := json.Marshal(cat)

	w.Header().Set("Content-Type", "text/json")

	if err != nil {
		res := "{'result': 'failure'}"
		result, err := json.Marshal(res)
		panicOnErr(err)
		fmt.Fprintf(w, string(result))
		panic(err)
	} else {
		fmt.Fprintf(w, string(catJSON))
	}
}

func createCat(c web.C, w http.ResponseWriter, r *http.Request) {
	var st ds.Cat
	err := json.NewDecoder(r.Body).Decode(&st)
	panicOnErr(err)
	fmt.Printf("%+v\n", st)
	err = st.Create()
	if err != nil {
		res := "{'result': 'failure'}"
		result, err := json.Marshal(res)
		fmt.Fprintf(w, string(result))
		panic(err)
	} else {
		res := "{'result': 'failure'}"
		result, err := json.Marshal(res)
		panicOnErr(err)
		fmt.Fprintf(w, string(result))
	}
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
	w.Header().Set("Content-Type", "text/json")

	if err != nil {
		res := "{'result': 'failure'}"
		result, err := json.Marshal(res)
		panicOnErr(err)
		fmt.Fprintf(w, string(result))
		panic(err)
	} else {
		fmt.Fprintf(w, string(orderJSON))
	}
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
		res := "{'result': 'failure'}"
		result, err := json.Marshal(res)
		panicOnErr(err)
		fmt.Fprintf(w, string(result))
		panic(err)
	} else {
		res := "{'result': 'success'}"
		result, err := json.Marshal(res)
		panicOnErr(err)
		fmt.Fprintf(w, string(result))
	}
}
