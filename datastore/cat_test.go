package datastore

import (
	"log"
	"strings"
	"testing"
	"time"

	"github.com/mmirolim/hack-project/conf"
)

func TestCreateCat(t *testing.T) {
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
	var i Cat
	i.Name = "Rollton"
	i.Desc = "yummy"
	i.UpdatedAt = time.Now()
	i.StaffID = 2
	i.SetDefaults()

	err = i.Create()
	if err != nil {
		t.Error(err)
	}

}

func TestFindOneCat(t *testing.T) {
	var i Cat
	err := i.FindOne(Where{Field: "ID", Value: 1, Crit: "="})
	if err != nil {
		t.Error(err)
	}
}

func TestFindAllCat(t *testing.T) {
	var i Cat

	i.Name = "Rollton"
	i.Desc = "yummy"
	i.UpdatedAt = time.Now()
	i.StaffID = 2
	i.SetDefaults()

	err := i.Create()
	if err != nil {
		t.Error(err)
	}

	_, err = i.FindAll(Where{Field: "desc", Crit: "=", Value: "yummy"}, 10)
	if err != nil {
		t.Error(err)
	}
}

func TestCatUpdateCat(t *testing.T) {
	var cat Cat
	err := cat.FindOne(Where{Field: "id", Value: 1, Crit: "="})
	if err != nil {
		t.Error(err)
	}

	cat.Name = "Bigbon"

	err = cat.Update()
	if err != nil {
		t.Error(err)
	}

	err = cat.FindOne(Where{Field: "id", Value: 1, Crit: "="})
	if err != nil {
		t.Error(err)
	}

	log.Printf("%+v\n", cat)
}

func TestDeleteCat(t *testing.T) {
	//staff is deleted, but staff.Get(1) is returns the staff
	var cat Cat
	err := cat.FindOne(Where{"id", ">", 0})
	if err != nil {
		t.Error(err)
	}
	err = cat.Delete()
	if err != nil {
		t.Error(err)
	}
}
