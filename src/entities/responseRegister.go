package entities

//ResponseRegister modelo de respuesta del registro de usuaurio
type ResponseRegister struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}
