package objectdb

import (
	"log"
	"testing"

	"github.com/cdvelop/dbtools"
	"github.com/cdvelop/model"
)

func (c *Connection) cloneTest(tables []*model.Object, t *testing.T) {

	t.Run((`CLONE: Caso 1  se espera clonación de tablas correcta`), func(t *testing.T) {

		// clonamos las tablas
		for _, table := range tables {
			err := dbtools.ClonDATABLE(c, table)
			if err != nil {
				log.Fatalln("error al clonar tabla "+table.ObjectName+" ", err)
			}
		}

	})

}
