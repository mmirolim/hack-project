package datastore

import (
	"fmt"
	"github.com/mmirolim/hack-project/conf"
	"strings"
	"testing"
)

func TestCreateTable(t *testing.T) {
	var table Table

	table.Alias = "Birinchi stol"
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

	_, err := Initialize(App.DS)
	if err != nil {
		t.Error(err)
	}

	err = table.Create()
	if err != nil {
		t.Error(err)
	}
}

func TestGetAllTables(t *testing.T) {
	var tables Tables

	//get all tables
	tables, err := tables.GetAll()
	if err != nil {
		t.Error(err)
	}

	//debug
	fmt.Printf("%+v\n", tables)
}

func TestGetTable(t *testing.T) {
	//get table
	var table Table
	var err error
	table, err = table.Get(1)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", table)
}

func TestUpdateTable(t *testing.T) {

	var table Table
	var newTable Table

	newTable.Alias = "ikkinchi stolik"

	//get table
	table, err := table.Get(1)
	if err != nil {
		t.Error(err)
	}

	//update table
	err = table.Update(newTable)
	if err != nil {
		t.Error(err)
	}

	//check updates
	table, err = table.Get(1)
	fmt.Printf("%+v\n", table)
}

func TestDeleteTable(t *testing.T) {
	//table is deleted, but table.Get(1) is returns the table
	var table Table
	table, err := table.Get(1)
	table.Delete()
	if err != nil {
		t.Error(table)
	}
}

func TestDropTableTables(t *testing.T) {
	err := DropTable()
	if err != nil {
		fmt.Println(err)
	}
}
