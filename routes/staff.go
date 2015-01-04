package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/zenazn/goji/web"

	ds "github.com/mmirolim/hack-project/datastore"
)

func getStaffAll(c web.C, w http.ResponseWriter, r *http.Request) {
	var st ds.Staff
	var err error
	sts, err := st.FindAll(ds.Where{"id", ">", 0}, 0)
	if err != nil {
		panic(err)
	}
	jsn, err := json.Marshal(sts)
	fmt.Fprintf(w, string(jsn))
}

func getStaff(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get user %s", c.URLParams["id"])
}

func createStaff(c web.C, w http.ResponseWriter, r *http.Request) {
	//var st ds.Staff
	var st map[string]interface{}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		fmt.Printf("%+v\n", data)
		err = json.Unmarshal(data, &st)
	}
	fmt.Printf("%+v\n", st)
	panicOnErr(err)
	fmt.Printf("%+v\n", st)
	//err = st.Create()
	//panicOnErr(err)
	fmt.Fprintf(w, "success")
}

func updateStaff(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update user %s", c.URLParams["id"])
}
