package objectdb

import (
	"log"
	"testing"
)

func (c Connection) readTest(t *testing.T) {

	for _, data := range dataTestCRUD {

		if data.Result { //solo los casos de éxito

			t.Run(("READ: "), func(t *testing.T) {
				out := c.ReadObject(defaulTableName, data.IdRecovered)
				if len(out) == 0 {
					log.Fatalf("!!! READ data: [%v] resp\n", out)
					t.Fail()
				}

			})
		}
	}
}
