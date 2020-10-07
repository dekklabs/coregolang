package userapi

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/dekklabs/apirest/src/entities"
	"github.com/dekklabs/apirest/src/model/usermodel"

	"github.com/dekklabs/apirest/src/middlew"
)

//UpdateCoverAPI api para actualizar el cover
func UpdateCoverAPI(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("cover")

	var extension = strings.Split(handler.Filename, ".")[1]

	var id = strconv.Itoa(int(middlew.IDUsuario))

	var archivo string = "uploads/covers/" + id + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, "Error al subir la imagen", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen", http.StatusBadRequest)
		return
	}

	var usuario entities.Usuario
	var status int64
	usuario.Cover = id + "." + extension
	status, err = usermodel.UpdateCover(usuario, id)

	if err != nil || status == 0 {
		http.Error(w, "Error al grabar la imagen", http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
