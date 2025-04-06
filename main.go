package main

import (
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//
	//conf, err := conf.NewConfig()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//db, err := psql.NewDb(conf.PgsqlNameServe)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//u := usecase.NewUseCase(&db)
	//
	//rest.Server_Echo(conf, u)

	sql := sqlbuilder.Update("user").Set("name")
	fmt.Print(sql)
}
