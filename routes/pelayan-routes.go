package routes

import (
	"github.com/gorilla/mux"
	"github.com/josepnapitupulu/API_Pelayan/controllers"
)

var RegisterPelayansRoutes = func(router *mux.Router) {
	// router.HandleFunc("/", controllers.Index).Methods("GET")
	// router.HandleFunc("/jemaatBaru", controllers.Create).Methods("POST")
	router.HandleFunc("/pelayan/", controllers.CreatePelayan).Methods("POST")
	router.HandleFunc("/pelayan/", controllers.GetPelayan).Methods("GET")
	router.HandleFunc("/pelayan/{pelayanId}", controllers.GetPelayanById).Methods("GET")
	router.HandleFunc("/pelayan/{pelayanId}", controllers.UpdatePelayan).Methods("PUT")
	router.HandleFunc("/pelayan/{pelayanId}", controllers.DeletePelayan).Methods("DELETE")
}
