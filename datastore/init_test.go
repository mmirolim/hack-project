package datastore

import (
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

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS orders ( tableID INTEGER , staffID INTEGER, ID INTEGER, PRIMARY KEY(ID ASC))")
	if err != nil {
		t.Error(err)
	}

	_, err = db.Exec("INSERT INTO orders VALUES (1, ?,?)", order.TableID, order.StaffID)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("initialize pass")
}
