package usermodel

import (
	"github.com/dekklabs/apirest/src/db"
	"github.com/dekklabs/apirest/src/entities"
)

//GetAvatar devuelve el avatar del usuario
func GetAvatar(id string) (usuario entities.Usuario, err error) {
	db, _ := db.Conexion()

	row, err := db.Query("SELECT image from usuario WHERE id = ?", id)
	if err != nil {
		return usuario, err
	}

	if row.Next() {
		var image string
		err = row.Scan(&image)

		if err != nil {
			return usuario, err
		}

		usuario = entities.Usuario{
			Image: image,
		}
	} else {
		return usuario, err
	}

	return usuario, nil
}
