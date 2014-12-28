package conf

import (
	"fmt"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	// read file and check parsing
	f := `
[ds]
	[ds.sqlite]
	name = "orders"
`
	// convert to Reader interface
	r := strings.NewReader(f)
	app, err := Read(r)
	if err != nil {
		t.Error("Read error")
	}
	fmt.Printf("%+v\n", app)
	want := "orders"
	if got := app.DS.SQLite.Name; got != want {
		t.Errorf("Datastore redis port %d, want %d", got, want)
	}

}
