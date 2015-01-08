package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	ds "github.com/mmirolim/hack-project/datastore"
	"github.com/zenazn/goji/web"
)

func getItemsAll(c web.C, w http.ResponseWriter, r *http.Request) {
	var item ds.Item
	items, err := item.FindAll(ds.Where{"id", ">", 0}, 0)
	panicOnErr(err)
	itemsJSON, err := json.Marshal(items)
	sendRes(w, err, itemsJSON)
}

func getItem(c web.C, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(c.URLParams["id"])
	panicOnErr(err)

	var item ds.Item
	err = item.FindOne(ds.Where{"id", "=", id})
	panicOnErr(err)

	itemJSON, err := json.Marshal(item)
	sendRes(w, err, itemJSON)
}

func createItem(c web.C, w http.ResponseWriter, r *http.Request) {
	var item ds.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	panicOnErr(err)
	err = item.Create()
	sendRes(w, err, nil)
}

func updateItem(c web.C, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(c.URLParams["id"])
	panicOnErr(err)

	var oldItem, item ds.Item
	err = oldItem.FindOne(ds.Where{"id", "=", id})
	panicOnErr(err)

	err = json.NewDecoder(r.Body).Decode(&item)
	panicOnErr(err)

	item.ID = id
	item.CreatedAt = oldItem.CreatedAt
	err = item.Update()
	panicOnErr(err)

	itemJSON, err := json.Marshal(&item)
	sendRes(w, err, itemJSON)
}

func deleteItem(c web.C, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(c.URLParams["id"])
	var item ds.Item
	err = item.FindOne(ds.Where{"id", "=", id})
	panicOnErr(err)

	err = item.Delete()
	panicOnErr(err)

	sendRes(w, err, nil)
}
