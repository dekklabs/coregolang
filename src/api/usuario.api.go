package api

import (
	"encoding/json"
	"net/http"

	"github.com/dekklabs/apirest/src/entities"
	"github.com/dekklabs/apirest/src/model"
)

//RegistroUsuarioApi api para registrar los usuarios
func RegistroUsuarioApi(response http.ResponseWriter, request *http.Request) {
	var usuario entities.Usuario

	err := json.NewDecoder(request.Body).Decode(&usuario)

	if err != nil {
		http.Error(response, "Error con los datos enviados", http.StatusBadRequest)
		return
	}

	if len(usuario.Nombre) <= 2 {
		http.Error(response, "El nombre debe tener más de 2 sílabas", http.StatusBadRequest)
		return
	}

	err = model.RegistroUsuario(&usuario)

	if err != nil {
		http.Error(response, "Error ", http.StatusBadRequest)
		return
	}

	response.WriteHeader(http.StatusOK)
}
