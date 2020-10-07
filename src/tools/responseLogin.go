package tools

import "github.com/dekklabs/apirest/src/entities"

//ResponseApi regresa un mensaje de tipo ResponseRegister
func ResponseApiLogin(err bool, message string, token string) (registro entities.ResponseLogin) {
	registro = entities.ResponseLogin{
		Error:   err,
		Message: message,
		Token:   token,
	}
	return
}
