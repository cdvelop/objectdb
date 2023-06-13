package objectdb

import (
	"fmt"
	"log"
	"strings"
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

// ReadObject ej: user,map[string]string{"id_user","111"}
func (c Connection) ReadObject(table_name string, where_fields map[string]string) map[string]string {
	// Construir la cláusula WHERE
	var whereConditions string
	for field, value := range where_fields {
		whereConditions += fmt.Sprintf("%s = '%s' AND ", field, value)
	}
	whereConditions = strings.TrimSuffix(whereConditions, " AND ")

	// Construir la consulta SQL
	sql := fmt.Sprintf("SELECT * FROM %s WHERE %s;", table_name, whereConditions)
	out, err := c.QueryOne(sql)
	if err != nil {
		log.Printf("Error al Obtener Última Data de la Tabla [%s] en la función ReadObject: %s\nSQL: [%s]\n", table_name, err.Error(), sql)
		return nil
	}

	return out
}

// SelectValue retorna valor de una consulta sql
func (c Connection) SelectValue(sql string) (out string, ok bool) {
	c.Open()
	defer c.Close()

	row := c.QueryRow(sql)
	err := row.Scan(&out)
	if err != nil {
		out = err.Error()
		return
	}
	ok = true
	return
}
