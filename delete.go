package objectdb

import (
	"fmt"
)

// DeleteObjects borra objetos de la base de datos según nombre de la tabla y ids.
func (c Connection) DeleteObjects(table_name string, all_data ...*map[string]string) error {
	c.Open()
	defer c.Close()

	tx, err := c.Begin()
	if err != nil {
		return filterMessageDBtoClient(err, table_name)
	}

	for _, data := range all_data {
		if data != nil {
			// borramos el objeto usando la clave primaria como condición
			query := fmt.Sprintf("DELETE FROM %s WHERE %s = %s", table_name, "id_"+table_name, c.PlaceHolders(1))
			stmt, err := tx.Prepare(query)
			if err != nil {
				tx.Rollback()
				return filterMessageDBtoClient(err, table_name, data)
			}
			defer stmt.Close()

			_, err = stmt.Exec((*data)["id_"+table_name])
			if err != nil {
				tx.Rollback()
				return filterMessageDBtoClient(err, table_name, data)
			}
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return filterMessageDBtoClient(err, table_name)
	}

	return nil
}
