package userapi

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/dekklabs/apirest/src/entities"
	"github.com/dekklabs/apirest/src/middlew"
	"github.com/dekklabs/apirest/src/model/usermodel"
)

type responseImage struct {
	Error   bool
	Message string
}

//UploadAvatarAPI api para subir avatar a un usuario
func UploadAvatarAPI(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")

	var extension = strings.Split(handler.Filename, ".")[1]

	var idUser = strconv.Itoa(int(middlew.IDUsuario))

	var archivo string = "uploads/avatars/" + idUser + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := responseImageFunc(true, "Error al subir la imagen")
		json.NewEncoder(w).Encode(resp)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := responseImageFunc(true, "Error al copiar la imagen")
		json.NewEncoder(w).Encode(resp)
		return
	}

	var usuario entities.Usuario
	var status int64
	usuario.Image = idUser + "." + extension
	status, err = usermodel.UploadAvatarUser(usuario, idUser)

	if err != nil || status == 0 {
		w.WriteHeader(http.StatusBadRequest)
		resp := responseImageFunc(true, "Error al grabar el avatar en la DB"+err.Error())
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func responseImageFunc(err bool, message string) (resp responseImage) {
	resp = responseImage{
		Error:   err,
		Message: message,
	}
	return
}
