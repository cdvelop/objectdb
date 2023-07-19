package objectdb

import (
	"log"
	"testing"
)

func (c Connection) deleteTest(t *testing.T) {

	for _, data := range dataTestCRUD {
		if data.ExpectedError == "" { //solo los casos de éxito

			data.Data["id_usuario"] = data.IdRecovered

			t.Run(("DELETE: " + data.IdRecovered), func(t *testing.T) {

				object, exist := modelObjectForTest[data.Object]
				if !exist {
					log.Fatalf("objeto: %v no existe", data.Object)
				}

				// validar elemento aquí
				err := object.ValidateData(false, true, &data.Data)
				if err != nil {
					data.IdRecovered = err.Error()
					return
				}

				err = c.DeleteObjects(defaulTableName, data.Data)
				if err != nil {
					if err.Error() != data.ExpectedError {
						log.Fatalf("en objeto: [%v]\n=>la expectativa es: [%v]\n=>pero se obtuvo: [%v]\n%v", data.Object, data.ExpectedError, err, data.Object)
					}

				} else {

					element_exists := c.ReadObject(defaulTableName, map[string]string{"id_" + defaulTableName: data.IdRecovered})

					if len(element_exists) != 0 {
						log.Fatalf("Error no se borro elemento:\n %v\n En base de datos: %v\n", defaulTableName, c.DataBasEngine())

					}
				}
			})
		}
	}
}
