package objectdb

import (
	"fmt"
)

// Exists verifica si existe un que es y algo en la base de datos según sql
// ej: "la base de datos","tiendadb","sql para la consulta"
func (c Connection) Exists(textResponse, objectSelect, sql string) (ok bool) {
	c.Open()
	defer c.Close()

	var val string

	if val, ok = c.SelectValue(sql); ok && val == "1" {
		fmt.Printf(">>> ok %v %v existe\n", textResponse, objectSelect)
	} else {
		fmt.Printf("!!! %v %v no existe\n", textResponse, objectSelect)
	}
	return
}
