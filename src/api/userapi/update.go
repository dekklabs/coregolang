package userapi

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dekklabs/apirest/src/entities"
	"github.com/dekklabs/apirest/src/model"
)

type messageUpdate struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

//UpdateProfileApi api para actualizar u usuario
func UpdateProfileApi(w http.ResponseWriter, r *http.Request) {
	var usuario entities.Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)

	idUser := strconv.FormatInt(usuario.ID, 10)

	if idUser == "0" {
		w.WriteHeader(http.StatusBadRequest)
		resp := messageUpdate{
			Error:   true,
			Message: "Es necesario un ID",
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := messageUpdate{
			Error:   true,
			Message: "Datos con los datos enviados",
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	_, err = model.UpdateProfileModel(&usuario)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := messageUpdate{
			Error:   true,
			Message: "Error al actualizar los datos del perfil",
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp := messageUpdate{
		Error:   false,
		Message: "Correcto",
	}
	json.NewEncoder(w).Encode(resp)
}
