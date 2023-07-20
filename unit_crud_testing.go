package objectdb

import (
	"log"
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
	err := dbtools.CreateAllTABLES(c, tables...)
	if err != nil {
		log.Fatalln(err)
	}

	c.addataCrud()

	c.createTest(t)

	c.readTest(tables, t)

	// c.updateTest(t)

	// c.deleteTest(t)
}
