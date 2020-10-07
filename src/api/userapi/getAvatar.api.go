package userapi

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/dekklabs/apirest/src/model/usermodel"
)

type responseGetAvatar struct {
	Error   bool
	Message string
}

//GetAvatarAPI api devuelve el avatar de un usuario
func GetAvatarAPI(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if len(id) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		resp := responseGetAvatarFunc(true, "Debe ingresar un id")
		json.NewEncoder(w).Encode(resp)
		return
	}

	usuario, err := usermodel.GetAvatar(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := responseGetAvatarFunc(true, "Error al traer el archivo de la base de datos")
		json.NewEncoder(w).Encode(resp)
		return
	}

	OpenFile, err := os.Open("uploads/avatars/" + usuario.Image)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := responseGetAvatarFunc(true, "Error al abrir el archivo")
		json.NewEncoder(w).Encode(resp)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := responseGetAvatarFunc(true, "Error al copiar el archivo")
		json.NewEncoder(w).Encode(resp)
		return
	}
}

func responseGetAvatarFunc(err bool, messsage string) (resp responseGetAvatar) {
	resp = responseGetAvatar{
		Error:   err,
		Message: messsage,
	}
	return
}
