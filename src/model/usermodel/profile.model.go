package usermodel

import (
	"github.com/dekklabs/apirest/src/db"
	"github.com/dekklabs/apirest/src/entities"
)

//Profile obtiene datos del perfil de usuario
func Profile(idUser int) (user entities.Usuario, err error) {
	db, err := db.Conexion()

	if err != nil {
		return user, err
	}

	profile, err2 := db.Query("SELECT * FROM usuario WHERE id = ?;", idUser)

	if err2 != nil {
		return user, err
	}

	if profile.Next() {
		var id int64
		var nombre string
		var apellido string
		var username string
		var password string
		var email string
		var description string
		var image string
		var cover string
		var createdat string
		var updatedat string

		err = profile.Scan(&id, &nombre, &apellido, &username, &password, &email, &description, &image, &cover, &createdat, &updatedat)
		if err != nil {
			return user, err
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
			Created_at:  createdat,
			Updated_at:  updatedat,
		}
	} else {
		return user, err
	}

	return user, nil
}
