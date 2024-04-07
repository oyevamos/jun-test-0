package main

import (
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/oyevamos/jun-test-0.git/internal/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.MembersRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":7500", r))
}
