package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/zenazn/goji/web"

	ds "github.com/mmirolim/hack-project/datastore"
)

func getStaffAll(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get all users %s", "Staffs")
}

func getStaff(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get user %s", c.URLParams["id"])
}

func createStaff(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create user %s", c.URLParams["id"])
}

func updateStaff(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update user %s", c.URLParams["id"])
}
