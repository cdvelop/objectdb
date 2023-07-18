package objectdb

import (
	"log"
	"testing"
)

func (c Connection) updateTest(t *testing.T) {

	for _, data := range dataTestCRUD {

		if data.ExpectedError == "" { //solo los casos que no contienen error

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
				err = c.UpdateObjects(defaulTableName, &data.Data)
				if err != nil {
					log.Fatalf("Error esperado: [%v] pero se obtuvo: [%v]\n%v", data.ExpectedError, err, data.Object)

				}
			})
		}
	}
}
