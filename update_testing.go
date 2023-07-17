package objectdb

import (
	"log"
	"testing"
)

func (c Connection) updateTest(t *testing.T) {

	for _, data := range dataTestCRUD {

		if data.Result { //solo los casos de éxito

			data.Data["id_usuario"] = data.IdRecovered
			data.Data["apellido"] = "NUEVO APELLIDO"

			t.Run(("UPDATE: " + data.IdRecovered), func(t *testing.T) {

				object, exist := modelObjectForTest[data.Object]
				if !exist {
					log.Fatalf("objeto: %v no existe", data.Object)
				}

				// validar elemento aquí
				err := object.ValidateData(false, true, &data.Data)
				if err != nil {
					log.Fatalln(err)
				}
				// fmt.Println("=> DATA A ACTUALIZAR: ", data.Data)
				mg, ok := c.UpdateObjects(defaulTableName, data.Data)
				if !ok {
					log.Fatalf("%v message: %v data in: \n[%v]", mg, ok, data.Data)
				}
			})
		}
	}
}
