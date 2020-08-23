package db

import "database/sql"

//Conexion conexion con la base de datos
func Conexion() (db *sql.DB, err error) {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "123456"
	dbName := "usuario"

	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err)
	}
	return
}
