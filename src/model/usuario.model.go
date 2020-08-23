package model

import (
	"github.com/dekklabs/apirest/src/db"
	"github.com/dekklabs/apirest/src/entities"
)

//RegistroUsuario registro de usuario
func RegistroUsuario(usuario *entities.Usuario) (err error) {
	db, _ := db.Conexion()

	results, err := db.Exec("insert into usuario(nombre, apellido, edad) values(?, ?, ?);", usuario.Nombre, usuario.Apellido, usuario.Edad)

	if err != nil {
		return err
	}

	usuario.Id, _ = results.LastInsertId()
	return nil
}
