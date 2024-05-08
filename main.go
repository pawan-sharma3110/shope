package main

import (
	"database/sql"
	"log"
	"net/http"
	"shope/router"
)

var Db *sql.DB

func main() {

	r := router.CustomerRouter()
	println("Server starting on port 8030....")
	log.Fatal(http.ListenAndServe(":8030", r))

}
