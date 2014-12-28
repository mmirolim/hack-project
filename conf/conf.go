package conf

import (
	"io"

	"github.com/BurntSushi/toml"
)

type Datastore struct {
	SQLite struct {
		Name string
	}
}

type App struct {
	DS Datastore
}

func Read(r io.Reader) (App, error) {
	var conf App
	_, err := toml.DecodeReader(r, &conf)
	return conf, err
}
