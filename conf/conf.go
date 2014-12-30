package conf

import (
	"io"

	"github.com/BurntSushi/toml"
)

type Datastore struct {
	SQLite struct {
		Name string
		File string
	}
}

type Server struct {
	Port string
}

type Raspi struct {
	AdaptorName string
	BotName     string
	Led         struct {
		Name string
		Pin  string
	}
}

type App struct {
	DS  Datastore
	Srv Server
	Rs  Raspi
}

func Read(r io.Reader) (App, error) {
	var conf App
	_, err := toml.DecodeReader(r, &conf)
	return conf, err
}
