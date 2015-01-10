package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	ds "github.com/mmirolim/hack-project/datastore"
	"github.com/mmirolim/hack-project/services"
	"github.com/zenazn/goji/web"
)

func getCatsAll(c web.C, w http.ResponseWriter, r *http.Request) {
	var st ds.Cat
	sts, err := st.FindAll(ds.Where{"id", ">", 0}, 0)
	if err != nil {
		replyJson(w, Reply{500, err.Error()})
	}
	jsn, err := json.Marshal(sts)
	if err != nil {
		replyJson(w, Reply{400, err.Error()})
	}
	fmt.Fprintf(w, string(jsn))
}

func getCat(c web.C, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(c.URLParams["id"])
	if err != nil {
		replyJson(w, Reply{400, err.Error()})
		return
	}

	var cat ds.Cat
	err = cat.FindOne(ds.Where{"id", "=", id})
	if err != nil {
		replyJson(w, Reply{500, err.Error()})
		return
	}

	jsn, err := json.Marshal(cat)
	if err != nil {
		replyJson(w, Reply{500, err.Error()})
		return
	}
	fmt.Fprintf(w, string(jsn))
}

func createCat(c web.C, w http.ResponseWriter, r *http.Request) {
	var st ds.Cat
	err := json.NewDecoder(r.Body).Decode(&st)
	if err != nil {
		replyJson(w, Reply{400, err.Error()})
		services.LogDeb(err)
		return
	}

	err = st.Create()
	if err != nil {
		replyJson(w, Reply{500, err.Error()})
		return
	}

	replyJson(w, Reply{200, OK})

}

func updateCat(c web.C, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(c.URLParams["id"])
	if err != nil {
		replyJson(w, Reply{500, err.Error()})
		return
	}
	var oldCat, cat ds.Cat
	err = oldCat.FindOne(ds.Where{"id", "=", id})
	if err != nil {
		replyJson(w, Reply{500, err.Error()})
		return
	}
	cat.ID = id
	err = cat.Update()
	if err != nil {
		replyJson(w, Reply{500, err.Error()})
		return
	}
	jsn, err := json.Marshal(&cat)
	if err != nil {
		replyJson(w, Reply{500, err.Error()})
		return
	}

	fmt.Fprintf(w, string(jsn))
}

func deleteCat(c web.C, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(c.URLParams["id"])
	var cat ds.Cat
	err = cat.FindOne(ds.Where{"id", "=", id})
	if err != nil {
		replyJson(w, Reply{500, err.Error()})
		return
	}

	err = cat.Delete()
	if err != nil {
		replyJson(w, Reply{500, err.Error()})
		return
	}
	replyJson(w, Reply{200, OK})
}
