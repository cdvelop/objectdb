package objectdb

import (
	"fmt"

	"github.com/cdvelop/model"
	"github.com/cdvelop/strings"
)

func (c Connection) ReadDataDB(p struct {
	FROM_TABLE    string              //ej: "users,products" or: public.reservation, public.patient"
	SELECT        string              // "name, phone, address" default *
	WHERE         []map[string]string //ej: "patient.id_patient=reservation.id_patient, (OR) reservation.id_staff = '2'"
	AND_CONDITION bool                // OR default se agrega AND si es true
	ID            string              // unique search (usado en indexdb)
	ORDER_BY      string              // name,phone,address
	SORT_DESC     bool                //default ASC
	LIMIT         int                 // 10, 5, 100. note: Postgres y MySQL: "LIMIT 10", SQLite: "LIMIT 10 OFFSET 0" OR "" no limit
	RETURN_ANY    bool                // default string return []map[string]string, any = []map[string]interface{}
}, async_results func(r struct {
	String []map[string]string
	Any    []map[string]any
}, err string)) (sync_results []map[string]string, err string) {
	const this = "ReadDataDB "
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
		wheres_count       int64
	)

	if p.WHERE != nil && len(p.WHERE) != 0 {
		wheres_found = append(wheres_found, p.WHERE...)
	}

	// búsqueda por multiples ids
	field_id := model.PREFIX_ID_NAME + p.FROM_TABLE

	if p.ID != "" { // búsqueda por un único id
		wheres_found = append(wheres_found, map[string]string{field_id: p.ID})
	}

	if p.SELECT != "" { //campos específicos a seleccionar
		choose = p.SELECT
	}

	if len(wheres_found) != 0 {
		var condition = " WHERE "

		for _, where := range wheres_found {

			for key, value := range where {
				wheres_count++

				var where_value string

				if valueIsFieldName(key, value) || valueContainClauseAND(value) { // chequear valor si es de tipo nombre de campo de otra una tabla o contiene and
					// fmt.Println("KEY:", key, "VALUE:", value, " son de tipo nombre de campo")
					where_value = value

				} else {
					place_holder_index++

					where_value = c.PlaceHolders(place_holder_index)
					args = append(args, value)
				}

				where_conditions += condition + key + " = " + where_value

				condition = " AND "

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

	// fmt.Print(`

	// - SQL READ:  `+p.FROM_TABLE+`

	//    `+sql+`

	//    ARGUMENTOS:

	// `, args)
	// fmt.Println("wheres_count:", wheres_count)

	if wheres_count != 1 || (choose == "*" && wheres_count == 1) {
		// fmt.Println("QUERY ALL")
		return c.QueryAll(sql, args...)
	}

	// fmt.Println("QUERY ONE")
	rowMap, err := c.QueryOne(sql, args...)
	if err != "" {
		return nil, this + err
	}

	if rowMap != nil {
		sync_results = append(sync_results, rowMap)
	}

	return sync_results, ""
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
