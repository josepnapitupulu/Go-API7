package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/josepnapitupulu/API_Pelayan/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterPelayansRoutes(r)
	http.Handle("/", r)

	fmt.Print("Starting Server localhost:7074")
	log.Fatal(http.ListenAndServe("localhost:7074", r))
}