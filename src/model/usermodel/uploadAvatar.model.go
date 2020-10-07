package usermodel

import (
	"github.com/dekklabs/apirest/src/db"
	"github.com/dekklabs/apirest/src/entities"
)

//UploadAvatarUser sube el avatar del usuario a la base de datos
func UploadAvatarUser(usuario entities.Usuario, id string) (int64, error) {
	db, _ := db.Conexion()

	row, err := db.Exec("UPDATE usuario SET image = ? WHERE id = ?;", usuario.Image, id)

	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}
