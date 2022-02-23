package handlers

import (
	"log"
	"net/http"
	"os"

	middleware "Red_Social/middlewares"

	routers "Red_Social/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Funcion para interactuar con la API
//mux hace un manejo de la url
func Managers() {
	router := mux.NewRouter()

	//Rutas: nombre de la ruta - funcion que realiza - metodo el cual se envia lso datos

	//Registro y login usan el mismo middleware
	router.HandleFunc("/registro", middleware.CheckDB(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middleware.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/verPerfil", middleware.CheckDB(middleware.ValidateJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/actualizarPerfil", middleware.CheckDB(middleware.ValidateJWT(routers.EditProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middleware.CheckDB(middleware.ValidateJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweet", middleware.CheckDB(middleware.ValidateJWT(routers.LeoTweet))).Methods("GET")
	router.HandleFunc("/borroTweet", middleware.CheckDB(middleware.ValidateJWT(routers.EliminarTweet))).Methods("DELETE")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router) //Permitir a todo el mundo que acceda a nuestra API y verificar si existe la ruta

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
