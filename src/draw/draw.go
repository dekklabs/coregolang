package draw

import (
	"fmt"

	"github.com/dekklabs/apirest/src/db"
)

//Drawing muestra la gráfica en consola cuando se ejecuta el programa
func Drawing(port string) {

	_, err := db.Conexion()

	databaseStatus := "OFF"

	if err == nil {
		databaseStatus = "ON"
	}

	fmt.Println(`
~ Dekklabs
|--------------------------------|--------------
| Server running                 | Puerto: ` + port + `
|:----------------------------------------------
|                                               
| database                       |      ` + databaseStatus + `
|--------------------------------|--------------
	`)
}
