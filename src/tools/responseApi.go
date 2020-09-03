package tools

import "github.com/dekklabs/apirest/src/entities"

//ResponseApi regresa un mensaje de tipo ResponseRegister
func ResponseApi(err bool, mensaje string) (registro entities.ResponseRegister) {
	registro = entities.ResponseRegister{
		Error:   err,
		Message: mensaje,
	}
	return
}
