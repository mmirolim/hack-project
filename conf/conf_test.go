package conf

import (
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	// read file and check parsing
	f := `
[ds]
	[ds.sqlite]
	name = "sqlite3"
        file = "./foo.db"
[srv]
port = "3000"
`
	// convert to Reader interface
	r := strings.NewReader(f)
	app, err := Read(r)
	if err != nil {
		t.Error(err)
	}
	want := "./foo.db"
	if got := app.DS.SQLite.File; got != want {
		t.Errorf("SQLite db file %s, want %s", got, want)
	}
	want = "3000"
	if got := app.Srv.Port; got != want {
		t.Errorf("Server port %s, want %s", got, want)
	}

}
