package datastore

import (
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

	var App conf.App

	App, err := conf.Read(f)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}

	err = Initialize(App.DS)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}

}
