package objectdb

import (
	"github.com/cdvelop/dbtools"
	"github.com/cdvelop/model"
)

func (c *Connection) CreateTablesInDB(tables []*model.Object, result func(err string)) {
	const this = "CreateTablesInDB error "
	for _, t := range tables {

		if t.Table == "" {
			result(this + "error nombre de tabla no definido en objeto: " + t.ObjectName)
			return
		}

		if exist, err := c.TableExist(t.Table); !exist {
			// fmt.Println("TABLA ", t.Table, " ¡NO EXISTE! ", c.DataBasEngine())
			if err != "" {
				result(this + err)
				return
			}

			err := dbtools.CreateOneTABLE(c, t)
			if err != "" {
				result(this + "no se logro crear tabla: " + t.Table + " " + err)
				return
			}
		}
	}

	result("")
}

func (c *Connection) TableExist(table_name string) (exist bool, err string) {
	const this = "TableExist error "
	c.Open()
	defer c.Close()

	rows, e := c.Query(c.SQLTableExist(), table_name)
	if e != nil {
		return false, e.Error()
	}
	defer rows.Close()

	data, err := dbtools.FetchOne(rows)
	if err != "" {
		return false, this + err
	}

	for _, v := range data {
		if v == "true" {
			return true, ""
		}

		if v == table_name {
			return true, ""
		}
	}

	// fmt.Println(c.SQLTableExist(), "RESULTADO CONSULTA:", data)

	return false, ""
}
