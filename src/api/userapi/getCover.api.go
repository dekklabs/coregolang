package userapi

import (
	"io"
	"net/http"
	"os"

	"github.com/dekklabs/apirest/src/model/usermodel"
)

//GetCoverAPI api para obtener el cover
func GetCoverAPI(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if len(id) < 1 {
		http.Error(w, "Debe ingresar un id", http.StatusBadRequest)
		return
	}

	usuario, err := usermodel.GetCover(id)

	if err != nil {
		http.Error(w, "Error al traer el archivo de la base de datos", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/covers/" + usuario.Cover)

	if err != nil {
		http.Error(w, "Error al abrir el archivo", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error al copiar el archivo", http.StatusBadRequest)
		return
	}
}
