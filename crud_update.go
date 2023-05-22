package objectdb

import (
	"fmt"

	"github.com/cdvelop/dbtools"
)

// UpdateObjects
func (c Connection) UpdateObjects(table_name string, all_data ...map[string]string) (message string, ok bool) {

	tx, err := c.DB.Begin()
	if err != nil {
		c.filterMessageDBtoClient(err.Error(), table_name, &message)
		return
	}
	defer tx.Rollback()

	for _, data := range all_data {
		sql := fmt.Sprintf("UPDATE %s SET ", table_name)
		values := make([]interface{}, 0)
		var id_pk string
		var id_val string
		var index uint8
		for field, value := range data {
			if _, pk := dbtools.IdpkTABLA(field, table_name); !pk {
				index++

				sql += fmt.Sprintf("%s = %s, ", field, c.PlaceHolders(index))
				values = append(values, value)

			} else {
				id_pk = field
				id_val = value
			}

		}
		index++

		sql = sql[:len(sql)-2] + " WHERE " + id_pk + " = " + c.PlaceHolders(index)

		values = append(values, data[id_val])

		_, err := tx.Exec(sql, values...)
		if err != nil {
			c.filterMessageDBtoClient(err.Error(), table_name, &message, data)
			return
		}
	}

	if err := tx.Commit(); err != nil {
		c.filterMessageDBtoClient(err.Error(), table_name, &message)
		return
	}

	return "", true
}
