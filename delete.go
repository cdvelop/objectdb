package objectdb

import (
	"fmt"

	"github.com/cdvelop/model"
)

// DeleteObjectsInDB borra objetos de la base de datos según nombre de la tabla y ids.
func (c Connection) DeleteObjectsInDB(table_name string, all_data ...map[string]string) (err string) {
	const this = "DeleteObjectsInDB error "
	c.Open()
	defer c.Close()

	tx, e := c.Begin()
	if e != nil {
		return this + filterMessageDBtoClient(e.Error(), table_name)
	}

	for _, data := range all_data {
		if data != nil {
			// borramos el objeto usando la clave primaria como condición
			query := fmt.Sprintf("DELETE FROM %s WHERE %s = %s", table_name, model.PREFIX_ID_NAME+table_name, c.PlaceHolders(1))
			stmt, e := tx.Prepare(query)
			if e != nil {
				tx.Rollback()
				return this + filterMessageDBtoClient(e.Error(), table_name, data)
			}
			defer stmt.Close()

			_, e = stmt.Exec(data[model.PREFIX_ID_NAME+table_name])
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
