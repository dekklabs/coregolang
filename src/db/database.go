package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//Conexion conexion con la base de datos
func Conexion() (db *sql.DB, err error) {

	// dbDriver := "mysql"
	// dbUser := "root"
	// dbPass := "123456"
	// dbName := "apirest"

	// db, _ = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	// err = db.Ping()

	// if err != nil {
	// 	fmt.Println("Error al conectar a la base de datos")
	// }

	dbDriver := "mysql"
	dbUser := "uwrafiap_gouserapi"
	dbPass := "zdnN99A5fQ0_"
	dbName := "uwrafiap_goapirest"
	dbHost := "tcp(50.31.174.103:3306)"

	// db, _ = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbHost+"/"+dbName)
	db, _ = sql.Open(dbDriver, fmt.Sprintf("%s:%s@%s/%s", dbUser, dbPass, dbHost, dbName))
	err = db.Ping()

	if err != nil {
		fmt.Println("Error al conectar a la base de datos")
	}

	return
}
