package api

import (
	"encoding/json"
	"net/http"

	"github.com/dekklabs/apirest/src/entities"
	"github.com/dekklabs/apirest/src/model"
	"github.com/dekklabs/apirest/src/utils"
)

//RegistroUsuarioApi api para registrar los usuarios
func RegistroUsuarioApi(response http.ResponseWriter, request *http.Request) {
	var usuario entities.Usuario

	err := json.NewDecoder(request.Body).Decode(&usuario)

	if err != nil {
		http.Error(response, "Error con los datos enviados"+err.Error(), http.StatusBadRequest)
		return
	}

	findEmail, _ := utils.VerifyExistsEmail(usuario.Email)

	if findEmail == true {
		http.Error(response, "El correo ya se encuentra registrado", http.StatusBadRequest)
		return
	}

	_, findUsername, _ := utils.VerifyExistsUsername(usuario.Username)

	if findUsername == true {
		http.Error(response, "El username ya se encuentra registrado", http.StatusBadRequest)
		return
	}

	if len(usuario.Password) < 6 {
		http.Error(response, "La contraseña no puede ser menor a 6 caracteres", http.StatusBadRequest)
		return
	}

	err = model.RegistroUsuario(&usuario)

	if err != nil {
		http.Error(response, "Error ", http.StatusBadRequest)
		return
	}

	response.WriteHeader(http.StatusOK)
}
