package objectdb

import (
	"fmt"
	"strconv"
	"strings"
)

// table name ej: users,products
// limit: ej 10, 5, 100. note: Postgres y MySQL: "LIMIT 10", SQLite: "LIMIT 10 OFFSET 0" OR "" no limit
// order_by: ej: name,phone,address
// choose:"name, phone, address" default *
func (c Connection) ReadObjectsInDB(table_name string, params map[string]string) ([]map[string]string, error) {
	// Verificar si queremos leer todos los objetos o solo un objeto específico
	var (
		read_all           = true
		where_conditions   string
		order_by           string
		limit_clause       string
		args               []interface{}
		place_holder_index uint8
		choose             = "*"
	)

	for key, value := range params {

		switch {
		case key == "id_"+table_name:
			read_all = false
			place_holder_index++
			where_conditions = " WHERE " + key + " = " + c.PlaceHolders(place_holder_index)
			args = append(args, value)

		case key == "limit": // Verificar si se proporciona un límite para la consulta
			limit, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			place_holder_index++
			limit_clause = " LIMIT " + c.PlaceHolders(place_holder_index) // según db
			args = append(args, limit)

		case key == "order_by": // Verificar si se proporcionan nombres para ordenar

			names_to_order := strings.Split(value, ",")
			for i, field_name := range names_to_order {

				if i == 0 {
					order_by = ` ORDER BY ` + field_name + ` ASC`
				} else {
					order_by += `, ORDER BY ` + field_name + ` ASC`
				}
			}

		case key == "choose": //campos específicos a seleccionar
			choose = value

		}
	}

	// Construir la consulta SQL
	sql := fmt.Sprintf("SELECT %s FROM %s%s%s%s;", choose, table_name, where_conditions, order_by, limit_clause)

	// fmt.Println("SQL READ: ", sql)

	// Ejecutar la consulta y obtener los resultados
	var rowsMap []map[string]string
	var err error

	if read_all {
		return c.QueryAll(sql, args...)
	}

	rowMap, err := c.QueryOne(sql, args...)
	if err != nil {
		return nil, err
	}

	if rowMap != nil {
		rowsMap = append(rowsMap, rowMap)
	}

	return rowsMap, nil
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
