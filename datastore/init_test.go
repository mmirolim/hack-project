package datastore

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/mmirolim/hack-project/conf"
	"log"
	"strings"
	"testing"
)

func TestInitialize(t *testing.T) {

	mockConf := `
[ds]
	[ds.sqlite]
	name = "sqlite3"
	file = "../foo.db"
[srv]
	port = "3000"
`

	f := strings.NewReader(mockConf)

	App, _ := conf.Read(f)

	fmt.Printf("%+v", App.DS)

	db, err := Initialize(App.DS)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}

	order := Order{
		TableID: 1,
		StaffID: 12,
	}

	var result sql.Result
	result, err = db.Exec("CREATE TABLE IF NOT EXISTS orders ( tableID INTEGER , staffID INTEGER, ID INTEGER, PRIMARY KEY(ID ASC))")

	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v", result)
	_, err = db.Exec("INSERT INTO orders VALUES (0, ?,?)", order.TableID, order.StaffID)
	if err != nil {
		t.Error(err)
	}

	items := []Item{
		{ID: 1, Name: "Osh", Desc: "Very tasty", Img: "/url/", Serving: 123.3,
			Cost: 12000, Unit: "portion", Status: 1},
		{ID: 2, Name: "Osh", Desc: "Very tasty", Img: "/url/", Serving: 123.3,
			Cost: 12000, Unit: "portion", Status: 1},
	}

	itemsJSON, err := json.Marshal(items)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(itemsJSON))

	fmt.Println("initialize pass")
}
