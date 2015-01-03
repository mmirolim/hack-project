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
	log.Println(i.createTableQuery())
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
		log.Printf("%+v\n", i)
		t.Error(err)
	}

}
