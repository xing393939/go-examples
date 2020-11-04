package main

import (
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"global-db/controllers"
	"global-db/sqldb"
)

func main() {
	sqldb.ConnectDB()

	http.HandleFunc("/", controllers.HelloWorld)

	s := &http.Server{
		Addr: fmt.Sprintf("%s:%s", "localhost", "5000"),
	}

	s.ListenAndServe()
}
