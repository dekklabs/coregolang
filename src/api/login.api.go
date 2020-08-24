package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dekklabs/apirest/src/entities"
	"github.com/dekklabs/apirest/src/model"
	"github.com/dekklabs/apirest/src/tokenjwt"
)

//LoginApi api para el inicio de sesion de un usuario
func LoginApi(w http.ResponseWriter, r *http.Request) {
	var usuario entities.Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)

	if err != nil {
		http.Error(w, "Usuario y/o contraseña inválidas", http.StatusBadRequest)
		return
	}

	user, exists := model.Login(usuario.Username, usuario.Password)

	if exists == false {
		http.Error(w, "Usuario no existe", http.StatusBadRequest)
		return
	}

	jwtKey, errjwt := tokenjwt.GenerateJWT(user)
	if errjwt != nil {
		http.Error(w, "Ocurrió un error al intentar generar el token correspondiente", http.StatusBadRequest)
		return
	}

	resp := entities.ResponseLogin{
		Token: jwtKey,
	}

	fmt.Println(resp)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
