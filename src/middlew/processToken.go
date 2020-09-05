package middlew

import (
	"errors"
	"strings"

	"github.com/dekklabs/apirest/src/entities"
	"github.com/dekklabs/apirest/src/utils"
	"github.com/dgrijalva/jwt-go"
)

//Username el valor del usarname en todos los endpoints
var Username string

//IDUsuario es el ide devuelto por el modelo
var IDUsuario int64

//ProcessToken procesa el token
func ProcessToken(tk string) (*entities.Clain, bool, int64, error) {
	miClave := []byte("apirestdekkdesarrollo")
	claims := &entities.Clain{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, 0, errors.New("Formato del token no válido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err != nil {
		_, encontrado, _ := utils.VerifyExistsUsername(claims.Username)
		if encontrado == true {
			Username = claims.Username
			IDUsuario = claims.ID
		}
		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, 0, errors.New("Token inválido")
	}

	return claims, false, 0, err
}
