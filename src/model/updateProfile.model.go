package model

import (
	"github.com/dekklabs/apirest/src/db"
	"github.com/dekklabs/apirest/src/entities"
)

//UpdateProfileModel actualiza el perfil de un usuario
func UpdateProfileModel(usuario *entities.Usuario) (int64, error) {
	db, _ := db.Conexion()

	row, err := db.Exec(
		"UPDATE user SET nombre = ?, apellido = ?, username = ?, description = ? WHERE id = ?",
		usuario.Nombre, usuario.Apellido, usuario.Username, usuario.Description, usuario.ID)

	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}
