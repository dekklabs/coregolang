package entities

//ResponseLogin entidad token
type ResponseLogin struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
	Token   string `json:"token"`
}
