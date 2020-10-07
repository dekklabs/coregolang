package usermodel

import (
	"github.com/dekklabs/apirest/src/db"
	"github.com/dekklabs/apirest/src/entities"
)

//GetCover obtiene el cover de la base de datos
func GetCover(id string) (usuario entities.Usuario, err error) {
	db, _ := db.Conexion()

	row, err := db.Query("SELECT cover from usuario WHERE id = ?;", id)

	if err != nil {
		return usuario, err
	}

	if row.Next() {
		var cover string
		err = row.Scan(&cover)

		if err != nil {
			return usuario, err
		}

		usuario = entities.Usuario{
			Cover: cover,
		}
	} else {
		return usuario, err
	}

	return usuario, nil
}
