package objectdb

import (
	"log"
	"strings"
	"testing"
)

func (c Connection) readTest(t *testing.T) {

	for _, data := range dataTestCRUD {

		if data.ExpectedError == "" { //solo los casos de éxito

			t.Run(("READ: "), func(t *testing.T) {
				out, err := c.ReadObjectsInDB(defaulTableName, map[string]string{"id_" + defaulTableName: data.IdRecovered})
				if err != nil {
					log.Fatalln("error en test de lectura ", err, data)
				}

				if len(out) == 0 {
					log.Fatalf("!!! READ data: [%v] resp\n", out)
				}

				// fmt.Println("=> DATA CAMBIADA?:", out)
				for _, o := range out {

					if !strings.Contains(o["apellido"], "NUEVO APELLIDO") {
						log.Fatalln("ERROR APELLIDO NUEVO NO CAMBIADO SALIDA:\n", out)
					}
				}

			})
		}
	}
}
