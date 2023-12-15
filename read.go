package objectdb

import (
	"fmt"
	"strings"

	"github.com/cdvelop/model"
)

func (c Connection) ReadAsyncDataDB(p model.ReadParams, callback func(r model.ReadResults, err string)) {

	callback(model.ReadResults{}, "ReadAsyncDataDB no implementado en paquete objectdb")
}

// from_tables ej: "users,products" or: public.reservation, public.patient"
// data ... map[string]string ej:{
// LIMIT: 10, 5, 100. note: Postgres y MySQL: "LIMIT 10", SQLite: "LIMIT 10 OFFSET 0" OR "" no limit
// ORDER_BY: name,phone,address
// SELECT: "name, phone, address" default *
// WHERE: "patient.id_patient = reservation.id_patient AND reservation.id_staff = '2'"
// ARGS: "1,4,33"
// }

func (c Connection) ReadSyncDataDB(p model.ReadParams, data ...map[string]string) (rowsMap []map[string]string, err string) {
	const this = "ReadSyncDataDB "
	// Verificar si queremos leer todos los objetos o solo un objeto específico
	var (
		// read_all           = true
		where_conditions   string
		order_by           string
		limit_clause       string
		args               []interface{}
		place_holder_index uint8
		choose             = "*"
		wheres_found       []map[string]string
	)

	if p.WHERE != nil && len(p.WHERE) != 0 {
		wheres_found = append(wheres_found, p.WHERE)
	}

	// búsqueda por multiples ids
	field_id := model.PREFIX_ID_NAME + p.FROM_TABLE
	for _, params := range data {
		for key, value := range params {
			if key == field_id {
				wheres_found = append(wheres_found, map[string]string{field_id: value})
			}
		}
	}

	if p.ID != "" { // búsqueda por un único id
		wheres_found = append(wheres_found, map[string]string{field_id: p.ID})
	}

	if p.SELECT != "" { //campos específicos a seleccionar
		choose = p.SELECT
	}

	if len(wheres_found) != 0 {
		var condition = " WHERE "

		for _, where := range wheres_found {

			place_holder_index++

			for key, value := range where {

				where_conditions += condition + key + " = " + c.PlaceHolders(place_holder_index)

				args = append(args, value)
			}

			if p.AND_CONDITION {
				condition = " AND "
			} else {
				condition = " OR "
			}
		}
	}

	if p.LIMIT != 0 { // Verificar si se proporciona un límite para la consulta
		place_holder_index++
		limit_clause = " LIMIT " + c.PlaceHolders(place_holder_index) // según db
		args = append(args, p.LIMIT)

	}

	if p.ORDER_BY != "" { // Verificar si se proporcionan nombres para ordenar
		var comma string
		order_by += ` ORDER BY `
		names_to_order := strings.Split(p.ORDER_BY, ",")
		for _, field_name := range names_to_order {

			order_by += comma + field_name
			comma = `,`
		}
	}

	// Construir la consulta SQL
	sql := fmt.Sprintf("SELECT %s FROM %s%s%s%s;", choose, p.FROM_TABLE, where_conditions, order_by, limit_clause)

	// fmt.Println("SQL READ: ", sql)
	// fmt.Println("ARGUMENTOS ", args)

	if len(wheres_found) != 1 {
		return c.QueryAll(sql, args...)
	}

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
