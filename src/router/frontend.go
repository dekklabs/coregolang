package router

import (
	"github.com/dekklabs/apirest/src/api"
	"github.com/dekklabs/apirest/src/api/productsapi"
	"github.com/dekklabs/apirest/src/api/userapi"
	"github.com/dekklabs/apirest/src/middlew"
	"github.com/gorilla/mux"
)

//Routers rutas del frontend
func Routers(router *mux.Router) {

	// Auth
	router.HandleFunc("/api/v1/register", api.RegistroUsuarioApi).Methods("POST")
	router.HandleFunc("/api/v1/login", api.LoginApi).Methods("POST")

	// User
	router.HandleFunc("/api/v1/update-user", middlew.VerifyJwt(userapi.UpdateProfileApi)).Methods("PUT")
	router.HandleFunc("/api/v1/profile", middlew.VerifyJwt(userapi.ProfileAPI)).Methods("GET")
	router.HandleFunc("/api/v1/upload-avatar", middlew.VerifyJwt(userapi.UploadAvatarAPI)).Methods("PUT")
	router.HandleFunc("/api/v1/get-avatar", userapi.GetAvatarAPI).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/upload-cover", middlew.VerifyJwt(userapi.UpdateCoverAPI)).Methods("PUT")
	router.HandleFunc("/api/v1/get-cover", middlew.VerifyJwt(userapi.GetCoverAPI)).Methods("GET")

	// Products
	router.HandleFunc("/api/v1/create-category", middlew.VerifyJwt(productsapi.CrearCategory)).Methods("POST")
	router.HandleFunc("/api/v1/update-category", middlew.VerifyJwt(productsapi.UpdateCategory)).Methods("PUT")
}
