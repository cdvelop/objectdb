package objectdb

import (
	"testing"

	"github.com/cdvelop/dbtools"
)

func (c *Connection) TestCrudStart(t *testing.T) {

	tables := c.addTestModelTablesDataBase()

	// eliminar tablas y data anterior
	for _, t := range tables {
		dbtools.DeleteTABLE(c, t.Name)
	}

	// crear tablas nuevas
	if !dbtools.CreateAllTABLES(c, tables...) {
		t.Fatal()
	}

	c.addataCrud()

	c.createTest(t)

	for _, table := range tables {
		dbtools.ClonDATABLE(c, table)
	}

	c.updateTest(t)

	c.readTest(t)

	c.deleteTest(t)
}
