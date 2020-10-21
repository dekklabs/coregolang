package productsapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dekklabs/apirest/src/model/products"

	"github.com/dekklabs/apirest/src/entities"
)

type responseCategory struct {
	Error   bool
	Message string
}

//CrearCategory api para crear categorías
func CrearCategory(w http.ResponseWriter, r *http.Request) {

	var category entities.CategoryProduct

	err := json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := responseFuncCategory(true, "Error al obtener el nombre")
		json.NewEncoder(w).Encode(resp)
	}

	err = products.CreateCategory(&category)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := responseFuncCategory(true, "Error al guardar la categoría")
		json.NewEncoder(w).Encode(resp)
	}

	resp := responseFuncCategory(false, "Correcto")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

//UpdateCategory api para actualizar una categoría
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var category entities.CategoryProduct

	err := json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := responseFuncCategory(true, "Error al recibir los datos")
		json.NewEncoder(w).Encode(resp)
	}

	status, err := products.UpdateCategory(&category)

	fmt.Println(status)

	if err != nil || status == 0 {
		w.WriteHeader(http.StatusBadRequest)
		resp := responseFuncCategory(true, "Error al actualizar")
		json.NewEncoder(w).Encode(resp)
	}

	w.WriteHeader(http.StatusOK)
}

func responseFuncCategory(status bool, message string) (resp responseCategory) {
	resp = responseCategory{
		Error:   status,
		Message: message,
	}
	return
}
