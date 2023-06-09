package objectdb

import (
	"fmt"
	"strings"
)

// CreateObjects
func (c Connection) CreateObjects(table_name string, all_data ...map[string]string) (message string, ok bool) {
	c.Open()
	defer c.Close()

	tx, err := c.Begin()
	if err != nil {
		filterMessageDBtoClient(err.Error(), table_name, &message)
		return
	}

	for i, data := range all_data {
		var columns, placeholders []string
		var values []interface{}

		var id string
		if ido, ok := data["id_"+table_name]; ok {
			id = ido //id objeto
		} else {
			//agregar id al objeto si este no existe
			id = c.GetNewID() //id nuevo
			data["id_"+table_name] = id
		}

		var index uint8
		for column, value := range data {
			index++
			columns = append(columns, column)
			placeholders = append(placeholders, c.PlaceHolders(index))
			values = append(values, value)
		}

		query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table_name, strings.Join(columns, ","), strings.Join(placeholders, ","))

		stmt, err := tx.Prepare(query)
		if err != nil {
			filterMessageDBtoClient(err.Error(), table_name, &message, data)
			tx.Rollback()
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(values...)
		if err != nil {
			filterMessageDBtoClient(err.Error(), table_name, &message, data)
			tx.Rollback()
			return
		}

		// si esta todo ok agregamos el id la data original
		all_data[i]["id_"+table_name] = id
	}

	if err := tx.Commit(); err != nil {
		filterMessageDBtoClient(err.Error(), table_name, &message)
		tx.Rollback()
		return
	}

	return "", true
}
