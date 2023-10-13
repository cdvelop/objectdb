package objectdb

import (
	"fmt"

	"github.com/cdvelop/dbtools"
	"github.com/cdvelop/model"
)

func (c *Connection) CreateTablesInDB(tables []*model.Object, action model.Subsequently) error {

	for _, t := range tables {

		if t.Table == "" {
			return model.Error("error nombre de tabla no definido en objeto:", t.Name)
		}

		if exist, err := c.TableExist(t.Table); !exist {

			if err != nil {
				return err
			}

			err := dbtools.CreateOneTABLE(c, t)
			if err != nil {
				return fmt.Errorf("no se logro crear tabla: %v\n%v", t.Table, err)
			}
		}
	}

	return nil
}

func (c *Connection) TableExist(table_name string) (bool, error) {
	c.Open()
	defer c.Close()

	rows, err := c.Query(c.SQLTableExist(), table_name)
	if err != nil {
		return false, err
	}

	defer rows.Close()

	return rows.Next(), nil
}
