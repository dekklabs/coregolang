package router

import (
	"github.com/dekklabs/apirest/src/api"
	"github.com/gorilla/mux"
)

//Routers rutas del frontend
func Routers(router *mux.Router) {
	router.HandleFunc("/api/v1/register", api.RegistroUsuarioApi).Methods("POST")
	router.HandleFunc("/api/v1/login", api.LoginApi).Methods("POST")
}
