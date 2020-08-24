package model

import (
	"github.com/dekklabs/apirest/src/entities"
	"github.com/dekklabs/apirest/src/utils"
	"golang.org/x/crypto/bcrypt"
)

//Login inicia sesion con un usuario
func Login(username, password string) (entities.Usuario, bool) {
	user, encontrado, _ := utils.VerifyExistsUser(username)

	usuario := entities.Usuario{}

	if encontrado == false {
		return usuario, false
	}

	userPassword := []byte(user.Password)
	passwordBody := []byte(password)

	user.Password = ""

	err := bcrypt.CompareHashAndPassword(userPassword, passwordBody)
	if err != nil {
		return usuario, false
	}

	return user, true
}
