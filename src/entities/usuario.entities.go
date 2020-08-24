package entities

//Usuario entidade de usuario
type Usuario struct {
	ID          int64  `json:"id"`
	Nombre      string `json:"nombre,omitempty"`
	Apellido    string `json:"apellido,omitempty"`
	Username    string `json:"username"`
	Password    string `json:"password,omitempty"`
	Email       string `json:"email"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image,omitempty"`
	Cover       string `json:"cover,omitempty"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}
