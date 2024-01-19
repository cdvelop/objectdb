package objectdb

import (
	"testing"

	"github.com/cdvelop/dbtools"
)

func (c *Connection) TestCrudStart(t *testing.T) {

	tables := c.addTestModelTablesDataBase()

	// eliminar tablas y data anterior
	for _, t := range tables {
		dbtools.DeleteTABLE(c, t.ObjectName)
	}

	// crear tablas nuevas
	err := dbtools.CreateTablesInDB(c, tables...)
	if err != "" {
		t.Fatal(err)
		return
	}

	// for _, table := range tables {
	// 	ok, err := c.TableExist(table.Table)
	// 	fmt.Println("TABLA", table.Table, "exist:", ok, " error?", err)
	// }

	c.addataCrud()

	c.createTest(t)

	var default_data_tests []map[string]string

	for _, d := range dataTestCRUD {
		if d.ExpectedError == "" && d.Object == defaulTableName { // solo los casos de éxito
			default_data_tests = append(default_data_tests, d.Data)
		}
	}

	c.cloneTest(tables, t)

	c.readTest(default_data_tests, t)

	c.updateTest(t)

	c.deleteTest(t)
}
