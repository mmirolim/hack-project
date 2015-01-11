package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	ds "github.com/mmirolim/hack-project/datastore"
	"github.com/zenazn/goji/web"
)

func getItemsAll(c web.C, w http.ResponseWriter, r *http.Request) {
	var item ds.Item
	items, err := item.FindAll(ds.Where{"id", ">", 0}, 0)
	if err != nil {
		replyJson(w, Reply{500, err.Error()})
		return
	}
	jsn, err := json.Marshal(items)
	if err != nil {
		replyJson(w, Reply{500, err.Error()})
		return
	}

	fmt.Fprintf(w, string(jsn))
}

func getItem(c web.C, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(c.URLParams["id"])
	if err != nil {
		replyJson(w, Reply{500, err.Error()})
		return
	}

	var item ds.Item
	err = item.FindOne(ds.Where{"id", "=", id})
	if err != nil {
		replyJson(w, Reply{500, err.Error()})
		return
	}

	jsn, err := json.Marshal(item)
	if err != nil {
		replyJson(w, Reply{500, err.Error()})
		return
	}

	fmt.Fprintf(w, string(jsn))
}

func createItem(c web.C, w http.ResponseWriter, r *http.Request) {
	var item ds.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		replyJson(w, Reply{400, err.Error()})
		return
	}
	err = item.Create()
	if err != nil {
		replyJson(w, Reply{500, err.Error()})
		return
	}

	replyJson(w, Reply{200, OK})
}

func updateItem(c web.C, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(c.URLParams["id"])
	if err != nil {
		replyJson(w, Reply{500, err.Error()})
		return
	}

	var oldItem, item ds.Item
	err = oldItem.FindOne(ds.Where{"id", "=", id})
	if err != nil {
		replyJson(w, Reply{500, err.Error()})
		return
	}

	err = json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		replyJson(w, Reply{500, err.Error()})
		return
	}

	item.ID = id
	item.CreatedAt = oldItem.CreatedAt
	err = item.Update()
	if err != nil {
		replyJson(w, Reply{500, err.Error()})
		return
	}

	jsn, err := json.Marshal(&item)
	if err != nil {
		replyJson(w, Reply{500, err.Error()})
		return
	}

	fmt.Fprintf(w, string(jsn))
}

func deleteItem(c web.C, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(c.URLParams["id"])
	if err != nil {
		replyJson(w, Reply{500, err.Error()})
		return
	}
	var item ds.Item
	err = item.FindOne(ds.Where{"id", "=", id})
	if err != nil {
		replyJson(w, Reply{500, err.Error()})
		return
	}

	err = item.Delete()
	if err != nil {
		replyJson(w, Reply{500, err.Error()})
		return
	}

	replyJson(w, Reply{200, OK})
}
