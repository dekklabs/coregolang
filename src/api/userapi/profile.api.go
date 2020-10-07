package userapi

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dekklabs/apirest/src/model/usermodel"

	"github.com/dekklabs/apirest/src/entities"
)

type responseUserProfile struct {
	Error   bool             `json:"error"`
	Message string           `json:"message"`
	User    entities.Usuario `json:"user"`
}

//ProfileAPI duelve datos del usuario
func ProfileAPI(w http.ResponseWriter, r *http.Request) {
	idTemp := r.URL.Query().Get("id")

	userTemp := entities.Usuario{}

	if len(idTemp) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		resp := responseFuncProfile(true, "Es necesario un ID1", userTemp)
		json.NewEncoder(w).Encode(resp)
	}

	id, _ := strconv.Atoi(idTemp)

	user, err := usermodel.Profile(id)
	user.Password = ""

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := responseFuncProfile(true, "No existe el usuario con el ID", userTemp)
		json.NewEncoder(w).Encode(resp)
	}

	resp := responseFuncProfile(false, "OK", user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func responseFuncProfile(err bool, message string, user entities.Usuario) (resp responseUserProfile) {
	resp = responseUserProfile{
		Error:   err,
		Message: message,
		User:    user,
	}
	return
}
