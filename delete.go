package objectdb

// DeleteObjectsInDB borra objetos de la base de datos según nombre de la tabla y ids.
func (c Connection) DeleteObjectsInDB(table_name string, backup_required bool, all_data ...map[string]string) (err string) {
	const this = "DeleteObjectsInDB "
	c.Open()
	defer c.Close()

	tx, e := c.Begin()
	if e != nil {
		return this + filterMessageDBtoClient(e.Error(), table_name)
	}

	for _, data := range all_data {
		if data != nil {
			// Construir la query dinámicamente
			query, values := c.buildDeleteQuery(table_name, data)

			// Preparar y ejecutar la sentencia SQL
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

// buildDeleteQuery construye la query DELETE de forma dinámica
func (c Connection) buildDeleteQuery(table_name string, where map[string]string) (query string, args []interface{}) {

	var (
		condition          string
		place_holder_index uint8
	)

	query = "DELETE FROM " + table_name + " WHERE "

	for key, value := range where {

		place_holder_index++
		query += condition + key + " = " + c.PlaceHolders(place_holder_index)
		args = append(args, value)
		condition = " AND "
	}

	query += ";"

	// fmt.Println("*** QUERY DELETE:", query)

	return query, args
}
