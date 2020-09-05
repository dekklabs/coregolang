package middlew

import (
	"encoding/json"
	"net/http"
)

type errorVerifyJwt struct {
	Error   bool   `json:"error"`
	Message string `json:"menssage"`
}

//VerifyJwt permite valor el JWT que nos llega con la petición
func VerifyJwt(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp := errorVerifyJwt{
				Error:   true,
				Message: "Error con el token " + err.Error(),
			}
			json.NewEncoder(w).Encode(resp)
			return
		}
		next.ServeHTTP(w, r)
	}
}
