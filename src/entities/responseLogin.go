package entities

//ResponseLogin entidad token
type ResponseLogin struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Token   string `json:"token"`
}
