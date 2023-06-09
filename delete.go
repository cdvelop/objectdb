package objectdb

import (
	"fmt"
)

// DeleteObjects borra objetos de la base de datos según nombre de la tabla y ids.
func (c Connection) DeleteObjects(table_name string, all_data ...map[string]string) (message string, ok bool) {
	c.Open()
	defer c.Close()

	tx, err := c.Begin()
	if err != nil {
		filterMessageDBtoClient(err.Error(), table_name, &message)
		return
	}

	for _, data := range all_data {

		// borramos el objeto usando la clave primaria como condición
		query := fmt.Sprintf("DELETE FROM %s WHERE %s = %s", table_name, "id_"+table_name, c.PlaceHolders(1))
		stmt, err := tx.Prepare(query)
		if err != nil {
			filterMessageDBtoClient(err.Error(), table_name, &message, data)
			tx.Rollback()
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(data["id_"+table_name])
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

	return "Registro(s) Eliminado(s)", true
}
