package objectdb

import (
	"testing"
)

func (c Connection) deleteTest(t *testing.T) {

	for _, data := range dataTestCRUD {
		if data.ExpectedError == "" { //solo los casos de éxito

			t.Run(("DELETE: " + data.Data["nombre"]), func(t *testing.T) {

				object, exist := modelObjectForTest[data.Object]
				if !exist {
					t.Fatalf("objeto: %v no existe", data.Object)
					return
				}

				// validar elemento aquí
				err := object.ValidateData(false, true, data.Data)
				if err != "" {
					t.Fatal(err)
					return
				}

				err = c.DeleteObjectsInDB(defaulTableName, data.Data)
				if err != "" {
					if err != data.ExpectedError {
						t.Fatalf("en objeto: [%v]\n=>la expectativa es: [%v]\n=>pero se obtuvo: [%v]\n%v", data.Object, data.ExpectedError, err, data.Object)
						return
					}

				} else {

					element_exists, err := c.ReadSyncDataDB(defaulTableName, map[string]string{"id_" + defaulTableName: data.Data["id_"+defaulTableName]})
					if err != "" {
						t.Fatal("error en test de lectura ", err, data)
						return
					}

					if len(element_exists) != 0 {
						t.Fatalf("Error no se borro elemento:\n %v\n En base de datos: %v\n", defaulTableName, c.DataBasEngine())
						return
					}
				}

				// fmt.Println("DATA PARA NOTIFICACIÓN DE ELIMINACIÓN ", notify_data)
			})
		}
	}
}
