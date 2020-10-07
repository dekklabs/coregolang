package utils

import (
	"github.com/dekklabs/apirest/src/db"
)

//VerifyExistsEmail verifica si un email existe en la base de datos
func VerifyExistsEmail(emailBody string) (encontrado bool, err error) {
	db, _ := db.Conexion()

	row := db.QueryRow("select * from usuario where email = ? LIMIT 1", emailBody)

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

	err = row.Scan(
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
		return encontrado, err
	}

	encontrado = true
	return encontrado, nil
}
