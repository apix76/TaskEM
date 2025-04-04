package rest

import (
	"TaskEM/conf"
	"TaskEM/db/psql"
	"TaskEM/entities"
	"TaskEM/usecase"
	"encoding/json"
	"fmt"

	"log"
	"net"
	"net/http"
)

type HTTPHandler struct{ Db psql.DbAccess }

func Http(conf conf.Conf) {
	mux := Handler(conf.PgsqlNameServe)
	l, err := net.Listen("tcp", conf.HttpPort)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.Serve(l, mux))

}
//type User entities.User

func Handler(PgsqlNameServe string) http.Handler {

	db, err := psql.NewDb(PgsqlNameServe)
	if err != nil {
		log.Fatal(err)
	}

	handler := HTTPHandler{Db: db}

	mux := http.NewServeMux()
	mux.HandleFunc(" /", handler.ServeUser)

	fmt.Println("Start http server")

	return mux
}

func (g *HTTPHandler) ServeUser(w http.ResponseWriter, req *http.Request) {
	//TODO: Подумать по поводу возврата сообщений


	result, err := usecase.


	if err
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(ResponseURL)
	if err != nil {
		log.Fatal(err)
	}
}
