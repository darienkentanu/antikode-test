package main

import (
	"antikode-test/api/router"
	"antikode-test/config"
	"antikode-test/util"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)

	router.RegisterBrandPath(r, db)
	router.RegisterOutletPath(r, db)

	server := &http.Server{}
	server.Handler = r
	server.ReadTimeout = 10 * time.Second
	server.WriteTimeout = 10 * time.Second
	server.Addr = fmt.Sprintf(":%d", config.Port)

	log.Printf("Starting server at %s \n", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
