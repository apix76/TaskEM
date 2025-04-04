package main

import (
	"TaskEM/conf"
	"TaskEM/server/rest"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf, err := conf.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	//rest.Http(conf)
	rest.Server_Echo(conf)
}
