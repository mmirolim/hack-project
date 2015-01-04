package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	ds "github.com/mmirolim/hack-project/datastore"
	"github.com/zenazn/goji/web"
)

func getStaffAll(c web.C, w http.ResponseWriter, r *http.Request) {
	var st ds.Staff
	var err error
	sts, err := st.FindAll(ds.Where{"id", ">", 0}, 0)
	panicOnErr(err)
	jsn, err := json.Marshal(sts)
	fmt.Fprintf(w, string(jsn))
}

func getStaff(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get user %s", c.URLParams["id"])
}

func createStaff(c web.C, w http.ResponseWriter, r *http.Request) {
	var st ds.Staff
	err := json.NewDecoder(r.Body).Decode(&st)
	panicOnErr(err)
	fmt.Printf("%+v\n", st)
	err = st.Create()
	panicOnErr(err)
	fmt.Fprintf(w, "success")
}

func updateStaff(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update user %s", c.URLParams["id"])
}
