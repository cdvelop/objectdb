package objectdb

import (
	"fmt"

	"github.com/cdvelop/dbtools"
)

// UpdateObjectsInDB
func (c Connection) UpdateObjectsInDB(table_name string, backup_required bool, all_data ...map[string]string) (err string) {
	const this = "UpdateObjectsInDB error "
	c.Open()
	defer c.Close()

	tx, e := c.DB.Begin()
	if e != nil {
		return this + filterMessageDBtoClient(e.Error(), table_name)
	}

	for _, data := range all_data {
		if data != nil {
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

			stmt, e := tx.Prepare(query)
			if e != nil {
				tx.Rollback()
				return this + filterMessageDBtoClient(e.Error(), table_name, data)
			}
			defer stmt.Close()

			_, e = stmt.Exec(values...)
			if e != nil {
				tx.Rollback()
				return this + filterMessageDBtoClient(e.Error(), table_name, data)
			}
		}
	}

	if e := tx.Commit(); e != nil {
		tx.Rollback()
		return this + filterMessageDBtoClient(e.Error(), table_name)
	}

	return ""
}
