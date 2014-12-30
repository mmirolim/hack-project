package datastore

import (
	"fmt"
	"github.com/mmirolim/hack-project/conf"
	//ds "github.com/mmirolim/hack-project/datastore"
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

	var App conf.App

	App, err := conf.Read(f)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	fmt.Printf("%+v", App.DS)
	err = Initialize(App.DS)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}

}
