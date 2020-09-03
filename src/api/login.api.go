package api

import (
	"encoding/json"
	"net/http"

	"github.com/dekklabs/apirest/src/entities"
	"github.com/dekklabs/apirest/src/model"
	"github.com/dekklabs/apirest/src/tokenjwt"
	"github.com/dekklabs/apirest/src/tools"
)

//LoginApi api para el inicio de sesion de un usuario
func LoginApi(w http.ResponseWriter, r *http.Request) {
	var usuario entities.Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := tools.ResponseApiLogin(true, "Usuario y/o contraseña inválidas", false, "")
		json.NewEncoder(w).Encode(resp)
		return
	}

	user, exists := model.Login(usuario.Username, usuario.Password)

	if exists == false {
		w.WriteHeader(http.StatusBadRequest)
		resp := tools.ResponseApiLogin(true, "Usuario no existe", false, "")
		json.NewEncoder(w).Encode(resp)
		return
	}

	jwtKey, errjwt := tokenjwt.GenerateJWT(user)
	if errjwt != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := tools.ResponseApiLogin(true, "Ocurrió un error al intentar generar el token correspondiente", false, "")
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp := entities.ResponseLogin{
		Error:   false,
		Message: "perfecto",
		Status:  true,
		Token:   jwtKey,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
