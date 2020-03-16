package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tidwall/buntdb"
)

var app = NewApp()

func main() {
	var err error

	var (
		version = flag.Bool("version", false, "version v1.0")
		config  = flag.String("config", "account.json", "config file.")
		port    = flag.Int("port", 1080, "listen port.")
	)

	flag.Parse()

	if *version {
		fmt.Println("v0.1")
		os.Exit(0)
	}

	app.SetAccounts(config)
	app.DB, err = buntdb.Open("wechat.db")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer app.DB.Close()

	InitRoute(app.Web.HttpServer)
	log.Println("Start AccessToken Server on ", *port)
	app.Web.StartServer(*port)
}
