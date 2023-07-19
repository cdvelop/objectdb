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

			//removemos nombre ya que es un campo único en el modelo y no se puede actualizar
			delete(data.Data, "nombre")

			t.Run(("UPDATE: " + data.IdRecovered), func(t *testing.T) {

				object, exist := modelObjectForTest[data.Object]
				if !exist {
					log.Fatalf("objeto: %v no existe", data.Object)
				}

				// fmt.Println("DATA A ACTUALIZAR: ", data.Data)

				// validar elemento aquí
				err := object.ValidateData(false, true, &data.Data)
				if err != nil {
					log.Fatalln(err)
				}
				// fmt.Println("=> DATA A ACTUALIZAR: ", data.Data)
				err = c.UpdateObjectsInDB(defaulTableName, data.Data)
				if err != nil {
					if err.Error() != data.ExpectedError {
						log.Fatalf("en objeto: [%v]\n=>la expectativa es: [%v]\n=>pero se obtuvo: [%v]\n%v", data.Object, data.ExpectedError, err, data.Object)
					}
				}

			})
		}
	}
}
