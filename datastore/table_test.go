package datastore

import (
	"log"
	"strings"
	"testing"
	"time"

	"github.com/mmirolim/hack-project/conf"
)

func TestCreateTable(t *testing.T) {
	mockConf := `
[ds]
	[ds.sqlite]
	name = "sqlite3"
	file = "foo.db"
[srv]
	port = "3000"
`

	f := strings.NewReader(mockConf)

	App, _ := conf.Read(f)

	_, err := Initialize(App.DS)
	if err != nil {
		t.Error(err)
	}

	var table Table

	table.Alias = "Alias"
	table.Desc = "Description"
	table.SetDefaults()
	table.CreatedAt = time.Now()
	table.StaffID = 12

	err = table.Create()
	if err != nil {
		t.Error(err)
	}
}

func TestFindOneTable(t *testing.T) {
	var table Table
	err := table.FindOne(Where{Field: "ID", Value: 1, Crit: "="})
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v\n", table)
}

func TestFindAllTable(t *testing.T) {
	var table Table

	table.Alias = "Alias"
	table.Desc = "Description"
	table.SetDefaults()
	table.CreatedAt = time.Now()
	table.StaffID = 12

	err := table.Create()
	if err != nil {
		t.Error(err)
	}

	tables, err := table.FindAll(Where{Field: "Alias", Crit: "=", Value: "Alias"}, 10)
	if err != nil {
		t.Error(err)
	}

	log.Printf("%+v\n", tables)
}

func TestUpdateTable(t *testing.T) {
	var table Table
	err := table.FindOne(Where{Field: "id", Value: 1, Crit: "="})
	if err != nil {
		t.Error(err)
	}

	table.StaffID = 1500

	err = table.Update()
	if err != nil {
		t.Error(err)
	}

	err = table.FindOne(Where{Field: "id", Value: 1, Crit: "="})
	if err != nil {
		t.Error(err)
	}

	log.Printf("%+v\n", table)
}

func TestDeleteTable(t *testing.T) {
	//staff is deleted, but staff.Get(1) is returns the staff
	var table Table
	err := table.FindOne(Where{"id", ">", 0})
	if err != nil {
		t.Error(err)
	}
	err = table.Delete()
	if err != nil {
		t.Error(err)
	}
}
