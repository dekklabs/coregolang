package utils

import (
	"github.com/dekklabs/apirest/src/db"
	"github.com/dekklabs/apirest/src/entities"
)

//VerifyExistsUser verofoca si un usuario existe
func VerifyExistsUser(usernameVerify string) (user entities.Usuario, encontrado bool, err error) {
	db, _ := db.Conexion()

	rows, err := db.Query("select * from user where username = ? limit 1", usernameVerify)

	if err != nil {
		encontrado = false
		return user, encontrado, err
	}

	if rows.Next() {
		var id int64
		var nombre string
		var apellido string
		var username string
		var password string
		var email string
		var description string
		var image string
		var cover string
		var created_at string
		var updated_at string

		err = rows.Scan(
			&id,
			&nombre,
			&apellido,
			&username,
			&password,
			&email,
			&description,
			&image,
			&cover,
			&created_at,
			&updated_at,
		)

		if err != nil {
			encontrado = false
			return user, encontrado, err
		}

		user = entities.Usuario{
			ID:          id,
			Nombre:      nombre,
			Apellido:    apellido,
			Username:    username,
			Password:    password,
			Email:       email,
			Description: description,
			Image:       image,
			Cover:       cover,
			Created_at:  created_at,
			Updated_at:  updated_at,
		}
	} else {
		encontrado = false
		return user, encontrado, nil
	}

	encontrado = true
	return user, encontrado, nil
}
