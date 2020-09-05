package entities

import "github.com/dgrijalva/jwt-go"

//Clain es la estructua usada para procesar el JWT
type Clain struct {
	Username string `json:"username"`
	ID       int64  `json:"id"`
	jwt.StandardClaims
}
