package objectdb

import (
	"fmt"
	"log"
)

// table name ej: users,products
// limit ej 10, 5, 100. 0 no limit etc. note: Postgres y MySQL: "LIMIT 10", SQLite: "LIMIT 10 OFFSET 0" OR "" no limit
// names_to order ej: name,phone,address
func (c Connection) ReadAllObjects(table_name string, limit uint8, names_to_order ...string) (out []map[string]string) {

	var order_by string

	if len(names_to_order) != 0 {

		for i, field_name := range names_to_order {

			if i == 0 {
				order_by = ` ORDER BY ` + field_name + ` ASC`
			} else {
				order_by += `, ORDER BY ` + field_name + ` ASC`
			}
		}
	}

	var limitation string
	if limit != 0 {
		limitation = fmt.Sprintf(` LIMIT %v`, limit)
	}

	rowsMap, err := c.QueryAll(`SELECT * FROM ` + table_name + order_by + limitation + `;`)
	if err != nil {
		log.Printf("Error al Obtener Todos Los Datos Tabla [%v] func ReadAllObjects %v\n", table_name, err.Error())
	}

	return rowsMap
}

// ReadObject
func (c Connection) ReadObject(table_name, id string) map[string]string {
	sql := fmt.Sprintf("SELECT * FROM %v WHERE id_%v ='%v';", table_name, table_name, id)
	out, err := c.QueryOne(sql)
	if err != nil {
		log.Printf("Error al Obtener Ultima Data Tabla [%v] func ReadObject %v\nSQL: [%v]\n", table_name, err.Error(), sql)
		return nil
	}

	return out
}

// SelectValue retorna valor de una consulta sql
func (c Connection) SelectValue(sql string) (out string, ok bool) {
	row := c.QueryRow(sql)
	err := row.Scan(&out)
	if err != nil {
		out = err.Error()
		return
	}
	ok = true
	return
}
