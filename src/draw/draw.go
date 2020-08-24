package draw

import (
	"fmt"
)

//Drawing muestra la gráfica en consola cuando se ejecuta el programa
func Drawing(port string) {
	fmt.Println(`
~ Dekklabs
|--------------------------------|--------------|
| Servidor Corriendo Hot Reload  | Puerto: ` + port + ` |
|:-------------------------------|:------------:|
|  DDDDD     DDDDD   D    D   D    D            |
|  D    D    D       D  D     D  D              |
|  D     D   DDDDD   DD       DD                |
|  D    D    D       D  D     D  D              |
|  DDDDD     DDDDD   D    D   D    D            |
|--------------------------------|--------------|
	`)
}
