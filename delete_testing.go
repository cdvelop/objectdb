package objectdb

import (
	"log"
	"testing"
)

func (c Connection) deleteTest(t *testing.T) {

	for _, data := range dataTestCRUD {
		if data.Result { //solo los casos de éxito

			data.Data["id_usuario"] = data.IdRecovered

			t.Run(("DELETE: " + data.IdRecovered), func(t *testing.T) {

				object, exist := modelObjectForTest[data.Object]
				if !exist {
					log.Fatalf("objeto: %v no existe", data.Object)
				}

				// validar elemento aquí
				if mg, ok := object.ValidateData(false, data.Data); !ok {
					data.IdRecovered = mg
					return
				}

				mg, ok := c.DeleteObjects(defaulTableName, data.Data)
				if !ok {
					log.Fatalf("message: %v ok[%v]\n", mg, ok)
				}

				element_exists := c.ReadObject(defaulTableName, map[string]string{"id_" + defaulTableName: data.IdRecovered})

				if len(element_exists) != 0 {
					log.Fatalf("Error no se borro elemento:\n %v\n En base de datos: %v\n", defaulTableName, c.DataBasEngine())

				}
			})
		}
	}
}
