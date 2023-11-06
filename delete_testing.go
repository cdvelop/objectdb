package objectdb

import (
	"log"
	"testing"
)

func (c Connection) deleteTest(t *testing.T) {

	for _, data := range dataTestCRUD {
		if data.ExpectedError == "" { //solo los casos de éxito

			t.Run(("DELETE: " + data.Data["nombre"]), func(t *testing.T) {

				object, exist := modelObjectForTest[data.Object]
				if !exist {
					log.Fatalf("objeto: %v no existe", data.Object)
				}

				// validar elemento aquí
				err := object.ValidateData(false, true, data.Data)
				if err != nil {
					log.Fatal(err)
				}

				err = c.DeleteObjectsInDB(defaulTableName, data.Data)
				if err != nil {
					if err.Error() != data.ExpectedError {
						log.Fatalf("en objeto: [%v]\n=>la expectativa es: [%v]\n=>pero se obtuvo: [%v]\n%v", data.Object, data.ExpectedError, err, data.Object)
					}

				} else {

					element_exists, err := c.ReadObjectsInDB(defaulTableName, map[string]string{"id_" + defaulTableName: data.Data["id_"+defaulTableName]})
					if err != nil {
						log.Fatalln("error en test de lectura ", err, data)
					}

					if len(element_exists) != 0 {
						log.Fatalf("Error no se borro elemento:\n %v\n En base de datos: %v\n", defaulTableName, c.DataBasEngine())

					}
				}

				// fmt.Println("DATA PARA NOTIFICACIÓN DE ELIMINACIÓN ", notify_data)
			})
		}
	}
}
