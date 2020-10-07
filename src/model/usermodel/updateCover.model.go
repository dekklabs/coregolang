package usermodel

import (
	"github.com/dekklabs/apirest/src/db"
	"github.com/dekklabs/apirest/src/entities"
)

//UpdateCover actualización del cover
func UpdateCover(usuario entities.Usuario, id string) (int64, error) {
	db, _ := db.Conexion()

	row, err := db.Exec("UPDATE usuario SET cover = ? WHERE id = ?;", usuario.Cover, id)

	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}
