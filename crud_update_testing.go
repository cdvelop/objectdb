package objectdb

import (
	"fmt"
	"log"
	"testing"
)

func (c Connection) updateTest(t *testing.T) {

	for index, data := range dataTestCRUD {

		if data.Result { //solo los casos de éxito
			lastName := fmt.Sprintf("apellido Actualizado %v", index)

			data.Data["id_usuario"] = data.IdRecovered
			data.Data["apellido"] = lastName

			t.Run(("UPDATE: " + data.IdRecovered), func(t *testing.T) {

				object, exist := modelObjectForTest[data.Object]
				if !exist {
					log.Fatalf("objeto: %v no existe", data.Object)
				}

				// validar elemento aquí
				if mg, ok := object.ValidateData(false, data.Data); !ok {
					data.IdRecovered = mg
					return
				}

				mg, ok := c.UpdateObjects(defaulTableName, data.Data)
				if !ok {
					log.Fatalf("%v message: %v data in: \n[%v]", mg, ok, data.Data)
				}
			})
		}
	}
}
