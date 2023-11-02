package objectdb

import (
	"fmt"
	"strings"

	"github.com/cdvelop/model"
)

// support: []map[string]string or map[string]string
func (c Connection) CreateObjectsInDB(table_name string, backup_required bool, items any) error {
	var all_data []map[string]string

	if data, ok := items.(map[string]string); ok {
		all_data = append(all_data, data)
	} else if data, ok := items.([]map[string]string); ok {
		all_data = data
	}

	if len(all_data) == 0 {
		return model.Error("error data ingresada para crear en tabla:", table_name, " incompatible, support only: []map[string]string or map[string]string")
	}

	c.Open()
	defer c.Close()

	tx, err := c.Begin()
	if err != nil {
		return filterMessageDBtoClient(err, table_name, items)
	}

	for i, data := range all_data {
		if data != nil {
			var columns, placeholders []string
			var values []interface{}

			var id string
			if ido, ok := data[model.PREFIX_ID_NAME+table_name]; ok {
				id = ido //id objeto
			} else {
				//agregar id al objeto si este no existe
				id = c.GetNewID() //id nuevo
				data[model.PREFIX_ID_NAME+table_name] = id
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
				tx.Rollback()
				return filterMessageDBtoClient(err, table_name, data)
			}
			defer stmt.Close()

			_, err = stmt.Exec(values...)
			if err != nil {
				tx.Rollback()
				return filterMessageDBtoClient(err, table_name, data)
			}

			// si esta todo ok agregamos el id la data original
			all_data[i]["id_"+table_name] = id
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return filterMessageDBtoClient(err, table_name)
	}

	return nil
}
