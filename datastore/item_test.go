package datastore

import (
	"log"
	"strings"
	"testing"
	"time"

	"github.com/mmirolim/hack-project/conf"
)

func TestCreateItem(t *testing.T) {
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
	var i Item
	i.Name = "Rollton"
	i.Desc = "yummy"
	i.Img = "/img/url"
	i.Serving = 123.45
	i.Cost = 1500
	i.Unit = "litre"
	i.Status = 3
	i.UpdatedAt = time.Now()
	i.StaffID = 2
	i.CatID = 1
	i.SetDefaults()

	err = i.Create()
	if err != nil {
		t.Error(err)
	}

}

func TestFindOne(t *testing.T) {
	var i Item
	err := i.FindOne(Where{Field: "ID", Value: 1, Crit: "="})
	if err != nil {
		t.Error(err)
	}
}

func TestFindAll(t *testing.T) {
	var i Item

	i.Name = "Rollton"
	i.Desc = "yummy"
	i.Img = "/img/url"
	i.Serving = 123.45
	i.Cost = 1500
	i.Unit = "litre"
	i.Status = 3
	i.UpdatedAt = time.Now()
	i.StaffID = 2
	i.CatID = 1
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

func TestItemUpdate(t *testing.T) {
	var item Item
	err := item.FindOne(Where{Field: "id", Value: 1, Crit: "="})
	if err != nil {
		t.Error(err)
	}

	item.CatID = 15

	err = item.Update()
	if err != nil {
		t.Error(err)
	}

	err = item.FindOne(Where{Field: "id", Value: 1, Crit: "="})
	if err != nil {
		t.Error(err)
	}

	log.Printf("%+v\n", item)
}

func TestDeleteItem(t *testing.T) {
	//staff is deleted, but staff.Get(1) is returns the staff
	var item Item
	err := item.FindOne(Where{"id", ">", 0})
	if err != nil {
		t.Error(err)
	}
	err = item.Delete()
	if err != nil {
		t.Error(err)
	}
}
