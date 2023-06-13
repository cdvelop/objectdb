package objectdb

import (
	"log"
	"strings"
	"testing"
)

func (c Connection) readTest(t *testing.T) {

	for _, data := range dataTestCRUD {

		if data.Result { //solo los casos de éxito

			t.Run(("READ: "), func(t *testing.T) {
				out := c.ReadObject(defaulTableName, map[string]string{"id_" + defaulTableName: data.IdRecovered})
				if len(out) == 0 {
					log.Fatalf("!!! READ data: [%v] resp\n", out)
					t.Fail()
				}

				// fmt.Println("=> DATA CAMBIADA?:", out)
				if !strings.Contains(out["apellido"], "NUEVO APELLIDO") {
					log.Fatalln("ERROR APELLIDO NUEVO NO CAMBIADO SALIDA:\n", out)
				}

			})
		}
	}
}
