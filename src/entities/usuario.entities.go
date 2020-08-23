package entities

//Usuario entidade de usuario
type Usuario struct {
	Id       int64  `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int64  `json:"edad"`
}
