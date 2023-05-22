package objectdb

import (
	"testing"

	"github.com/cdvelop/dbtools"
)

func (c *Connection) TestCrudStart(t *testing.T, orm dbtools.OrmAdapter) {

	tables := c.addTestModelTablesDataBase()

	db := dbtools.NewOperationDB(c.DB, orm)

	if !db.CreateAllTABLES(tables...) {
		t.Fatal()
	}

	c.addataCrud()

	c.createTest(t)

	c.readTest(t)

	c.updateTest(t)

	for _, table := range tables {
		db.ClonDATABLE(table)
	}

	c.deleteTest(t)
}
