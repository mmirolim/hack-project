package main

import (
	"flag"
	"log"
	"os"

	"github.com/mmirolim/hack-project/conf"
	"github.com/mmirolim/hack-project/routes"
	"github.com/zenazn/goji"
)

func main() {
	// read conf file
	f, err := os.Open("ex-conf.toml")
	fatalOnError(err)
	App, err := conf.Read(f)
	fatalOnError(err)
	// close conf file
	f.Close()
	// init routes
	m := routes.Initialize()
	// set goji server port
	flag.Set("bind", ":"+App.Srv.Port)
	// register routes
	goji.Handle("/*", m)
	// start server
	goji.Serve()

}

func fatalOnError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
