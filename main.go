package main

import (
	"flag"
	"log"
	"os"

	"github.com/mmirolim/hack-project/conf"
	ds "github.com/mmirolim/hack-project/datastore"
	"github.com/mmirolim/hack-project/routes"
	"github.com/mmirolim/hack-project/services"
	"github.com/zenazn/goji"
)

var (
	mode = flag.String("mode", "dev", "debug, dev, prod")
)

func init() {
	flag.Parse()
	flag.Usage()
}

func main() {
	// start logger
	services.Initialize(*mode)
	// read conf file
	f, err := os.Open("conf.toml")
	fatalOnError(err)
	App, err := conf.Read(f)
	fatalOnError(err)
	// close conf file
	f.Close()
	// init datastore
	_, err = ds.Initialize(App.DS)
	fatalOnError(err)
	// start status bot
	statusChan := make(chan ds.Status)
	go services.StartStatusBot(App.Rs, statusChan)
	// init routes
	m := routes.Initialize(statusChan)
	// set response format
	goji.Use(routes.JSON)
	// set goji server port
	flag.Set("bind", App.Srv.IP+":"+App.Srv.Port)
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
