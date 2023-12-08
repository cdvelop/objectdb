package objectdb

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cdvelop/model"
)

func (c Connection) ReadAsyncDataDB(p model.ReadParams, callback func(r model.ReadResult)) {

	callback(model.ReadResult{Error: "error ReadAsyncDataDB no implementado en paquete objectdb"})
}

// from_tables ej: "users,products" or: public.reservation, public.patient"
// data ... map[string]string ej:{
// LIMIT: 10, 5, 100. note: Postgres y MySQL: "LIMIT 10", SQLite: "LIMIT 10 OFFSET 0" OR "" no limit
// ORDER_BY: name,phone,address
// SELECT: "name, phone, address" default *
// WHERE: "patient.id_patient = reservation.id_patient AND reservation.id_staff = '2'"
// ARGS: "1,4,33"
// }

func (c Connection) ReadSyncDataDB(from_tables string, data ...map[string]string) (out []map[string]string, err string) {
	const this = "ReadSyncDataDB error "
	// Verificar si queremos leer todos los objetos o solo un objeto específico
	var (
		// read_all           = true
		where_conditions   string
		order_by           string
		limit_clause       string
		args               []interface{}
		place_holder_index uint8
		choose             = "*"
		total_ids_found    int64
	)

	for i, params := range data {

		for key, value := range params {

			switch {

			case key == "id_"+from_tables:
				total_ids_found++
				place_holder_index++

				if i == 0 {

					where_conditions = " WHERE " + key + " = " + c.PlaceHolders(place_holder_index)
				} else {

					where_conditions += " OR " + key + " = " + c.PlaceHolders(place_holder_index)
				}

				args = append(args, value)

			case key == "LIMIT": // Verificar si se proporciona un límite para la consulta
				limit, e := strconv.Atoi(value)
				if e != nil {
					return nil, this + e.Error()
				}
				place_holder_index++
				limit_clause = " LIMIT " + c.PlaceHolders(place_holder_index) // según db
				args = append(args, limit)

			case key == "ORDER_BY": // Verificar si se proporcionan nombres para ordenar
				var comma string
				order_by += ` ORDER BY `
				names_to_order := strings.Split(value, ",")
				for _, field_name := range names_to_order {

					order_by += comma + field_name
					comma = `,`
				}

			case key == "SELECT": //campos específicos a seleccionar
				choose = value

			case key == "WHERE": //se envió una consulta con where
				where_conditions = " WHERE " + value

			case key == "ARGS": //se envió una consulta con where
				new_args := strings.Split(value, ",")
				args = append(args, new_args)

			}
		}
	}

	// Construir la consulta SQL
	sql := fmt.Sprintf("SELECT %s FROM %s%s%s%s;", choose, from_tables, where_conditions, order_by, limit_clause)

	// fmt.Println("SQL READ: ", sql)
	// fmt.Println("ARGUMENTOS ", args)

	if total_ids_found != 1 {
		return c.QueryAll(sql, args...)
	}

	// Ejecutar la consulta y obtener los resultados
	var rowsMap []map[string]string

	rowMap, err := c.QueryOne(sql, args...)
	if err != "" {
		return nil, this + err
	}

	if rowMap != nil {
		rowsMap = append(rowsMap, rowMap)
	}

	return rowsMap, ""
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
