package router

import (
	"github.com/dekklabs/apirest/src/api"
	"github.com/dekklabs/apirest/src/api/userApi"
	"github.com/dekklabs/apirest/src/middlew"
	"github.com/gorilla/mux"
)

//Routers rutas del frontend
func Routers(router *mux.Router) {

	// Auth
	router.HandleFunc("/api/v1/register", api.RegistroUsuarioApi).Methods("POST")
	router.HandleFunc("/api/v1/login", api.LoginApi).Methods("POST")

	// User
	router.HandleFunc("/api/v1/update-user", middlew.VerifyJwt(userApi.UpdateProfileApi)).Methods("PUT")
}
