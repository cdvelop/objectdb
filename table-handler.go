package objectdb

import (
	"github.com/cdvelop/dbtools"
	"github.com/cdvelop/model"
)

func (c *Connection) CreateTablesInDB(tables []*model.Object, result func(error)) {

	for _, t := range tables {

		if t.Table == "" {
			result(model.Error("error nombre de tabla no definido en objeto:", t.ObjectName))
			return
		}

		if exist, err := c.TableExist(t.Table); !exist {
			// fmt.Println("TABLA ", t.Table, " ¡NO EXISTE! ", c.DataBasEngine())
			if err != nil {
				result(err)
				return
			}

			err := dbtools.CreateOneTABLE(c, t)
			if err != nil {
				result(model.Error("no se logro crear tabla:", t.Table, err))
				return
			}
		} else {
			// fmt.Println("TABLA ", t.Table, " ¡YA EXISTE!", c.DataBasEngine())
		}
	}

	result(nil)
}

func (c *Connection) TableExist(table_name string) (bool, error) {
	c.Open()
	defer c.Close()

	rows, err := c.Query(c.SQLTableExist(), table_name)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	data, err := dbtools.FetchOne(rows)
	if err != nil {
		return false, err
	}

	for _, v := range data {
		if v == "true" {
			return true, nil
		}

		if v == table_name {
			return true, nil
		}
	}

	// fmt.Println(c.SQLTableExist(), "RESULTADO CONSULTA:", data)

	return false, nil
}
