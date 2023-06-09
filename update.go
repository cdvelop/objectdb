package objectdb

import (
	"fmt"

	"github.com/cdvelop/dbtools"
)

// UpdateObjects
func (c Connection) UpdateObjects(table_name string, all_data ...map[string]string) (message string, ok bool) {
	c.Open()
	defer c.Close()

	tx, err := c.DB.Begin()
	if err != nil {
		filterMessageDBtoClient(err.Error(), table_name, &message)
		return
	}

	for _, data := range all_data {
		query := fmt.Sprintf("UPDATE %s SET ", table_name)
		values := make([]interface{}, 0)
		var field_pk string
		var id_value string
		var index uint8
		for field, value := range data {
			if _, pk := dbtools.IdpkTABLA(field, table_name); !pk {
				index++

				query += fmt.Sprintf("%s = %s, ", field, c.PlaceHolders(index))
				values = append(values, value)

			} else {
				field_pk = field
				id_value = value
			}

		}

		index++

		query = query[:len(query)-2] + " WHERE " + field_pk + " = " + c.PlaceHolders(index)

		values = append(values, id_value)

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
	}

	if err := tx.Commit(); err != nil {
		filterMessageDBtoClient(err.Error(), table_name, &message)
		tx.Rollback()
		return
	}

	return "Actualización Exitosa", true
}
