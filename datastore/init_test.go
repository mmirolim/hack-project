package datastore_test

import (
	"fmt"
	"github.com/mmirolim/hack-project/conf"
	ds "github.com/mmirolim/hack-project/datastore"
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

	db, err := ds.Initialize(App.DS)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}

	order := ds.Order{
		TableID: 1,
		StaffID: 2,
	}
	/*
		_, err = db.Query("CREATE TABLE orders ( tableID INTEGER , staffID INTEGER, ID INTEGER, PRIMARY KEY(ID ASC))")
		if err != nil {
			t.Error(err)
		}
	*/
	//_, err = db.Exec("INSERT INTO VALUES (1, ?,?)", order.TableID, order.StaffID)
	_, err = db.Query("INSERT INTO orders VALUES (1, ?,?)", order.TableID, order.StaffID)

	if err != nil {
		t.Error(err)
	}
	db.Exec("INSERT INTO orders VALUES (1, 2,3)")
	fmt.Println("initialize pass")
}
