package controllers

import (
	"fmt"
	"net/http"

	"global-db/sqldb"
)

// HelloWorld returns Hello, World
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	if err := sqldb.DB.Ping(); err != nil {
		fmt.Println("DB Error")
	}

	w.Write([]byte("Hello, World"))
}
