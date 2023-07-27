package objectdb

import (
	"fmt"

	"github.com/cdvelop/dbtools"
	"github.com/cdvelop/model"
)

func (c *Connection) CreateTablesInDB(tables ...*model.Object) error {

	for _, t := range tables {
		if exist, err := c.TableExist(t.Name); !exist {

			if err != nil {
				return err
			}

			err := dbtools.CreateOneTABLE(c, t)
			if err != nil {
				return fmt.Errorf("no se logro crear tabla: %v\n%v", t.Name, err)
			}
		}
	}

	return nil
}

func (c Connection) TableExist(table_name string) (bool, error) {
	rows, err := c.Query(c.SQLTableExist(), table_name)
	if err != nil {
		return false, err
	}

	defer rows.Close()

	return rows.Next(), nil
}
