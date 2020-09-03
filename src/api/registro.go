package api

import (
	"encoding/json"
	"net/http"

	"github.com/dekklabs/apirest/src/entities"
	"github.com/dekklabs/apirest/src/model"
	"github.com/dekklabs/apirest/src/tools"
	"github.com/dekklabs/apirest/src/utils"
)

//RegistroUsuarioApi api para registrar los usuarios
func RegistroUsuarioApi(response http.ResponseWriter, request *http.Request) {
	var usuario entities.Usuario

	err := json.NewDecoder(request.Body).Decode(&usuario)

	if err != nil { // Error en caso los datos que el usuario envia son incorrectos
		response.WriteHeader(http.StatusBadRequest)
		resp := tools.ResponseApi(true, "Error con los datos enviados"+err.Error())
		json.NewEncoder(response).Encode(resp)
		return
	}

	if len(usuario.Username) == 0 { // Error en caso no se reciba un username
		response.WriteHeader(http.StatusBadRequest)
		resp := tools.ResponseApi(true, "Es necesario un nombre de usuario")
		json.NewEncoder(response).Encode(resp)
		return
	}

	findEmail, _ := utils.VerifyExistsEmail(usuario.Email) // Función que verifica si el correo recibido existe en la base de datos

	if findEmail == true { // Error en caso el correo exista
		response.WriteHeader(http.StatusBadRequest)
		resp := tools.ResponseApi(true, "El correo ya se encuentra registrado")
		json.NewEncoder(response).Encode(resp)
		return
	}

	_, findUsername, _ := utils.VerifyExistsUsername(usuario.Username) // Función que verifica si el username recibido existe en la base de datos

	if findUsername == true { // Error en caso el usuario exista
		response.WriteHeader(http.StatusBadRequest)
		resp := tools.ResponseApi(true, "El username ya se encuentra registrado")
		json.NewEncoder(response).Encode(resp)
		return
	}

	if len(usuario.Password) < 6 { // Error en caso el password tenga menos de 6 caracteres
		response.WriteHeader(http.StatusBadRequest)
		resp := tools.ResponseApi(true, "Password no puede tener menos de 6 caracteres")
		json.NewEncoder(response).Encode(resp)
		return
	}

	err = model.RegistroUsuario(&usuario)

	if err != nil { // Error entidad
		response.WriteHeader(http.StatusBadRequest)
		resp := tools.ResponseApi(true, "Error")
		json.NewEncoder(response).Encode(resp)
		return
	}

	resp := tools.ResponseApi(false, "Registro completo") // Paso sin errores

	response.WriteHeader(http.StatusOK)
	response.Header().Add("Content-Type", "application/json")
	json.NewEncoder(response).Encode(resp)
}
