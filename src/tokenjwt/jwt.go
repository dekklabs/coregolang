package tokenjwt

import (
	"github.com/dekklabs/apirest/src/entities"
	"github.com/dgrijalva/jwt-go"
)

//GenerateJWT generar jwt para el login
func GenerateJWT(usuario entities.Usuario) (string, error) {
	miPassword := []byte("myworldifmypassword")

	payload := jwt.MapClaims{
		"id":          usuario.ID,
		"nombre":      usuario.Nombre,
		"apellido":    usuario.Apellido,
		"username":    usuario.Username,
		"password":    usuario.Password,
		"email":       usuario.Email,
		"description": usuario.Description,
		"image":       usuario.Image,
		"cover":       usuario.Cover,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miPassword)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
