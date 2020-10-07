package model

import (
	"github.com/dekklabs/apirest/src/db"
	"github.com/dekklabs/apirest/src/entities"
	"github.com/dekklabs/apirest/src/utils"
)

//RegistroUsuario registro de usuario
func RegistroUsuario(usuario *entities.Usuario) (err error) {
	db, _ := db.Conexion()

	usuario.Password, _ = utils.EncryptPassword(usuario.Password)

	usuario.Nombre = ""
	usuario.Apellido = ""
	usuario.Description = ""
	usuario.Image = ""
	usuario.Cover = ""

	results, err := db.Exec(`insert into usuario(
		nombre,
		apellido,
		username,
		password,
		email,
		description,
		image,
		cover)
		values(?, ?, ?, ?, ?, ?, ?, ?);`,
		usuario.Nombre,
		usuario.Apellido,
		usuario.Username,
		usuario.Password,
		usuario.Email,
		usuario.Description,
		usuario.Image,
		usuario.Cover)

	usuario.ID, _ = results.LastInsertId()
	return nil
}
