package objectdb

import (
	"log"
	"testing"
)

func (c Connection) createTest(t *testing.T) {

	for prueba, data := range dataTestCRUD {
		t.Run(("CREATE: " + prueba), func(t *testing.T) {

			object, exist := modelObjectForTest[data.Object]
			if !exist {
				log.Fatalf("objeto: %v no existe", data.Object)
			}

			// validar elemento aquí
			err := object.ValidateData(true, false, data.Data)
			if err != "" {
				if data.ExpectedError == "" {
					log.Fatalf("en objeto: [%v]\n=>la expectativa es: [%v]\n=>pero se obtuvo: [%v]\n%v", data.Object, data.ExpectedError, err, data.Object)
				}
				return
			} else {

				err = c.CreateObjectsInDB(data.Object, false, data.Data)
				if err != "" {
					if err != data.ExpectedError {
						log.Fatalf("en objeto: [%v]\n=>la expectativa es: [%v]\n=>pero se obtuvo: [%v]\n%v", data.Object, data.ExpectedError, err, data.Object)
					}

				}

			}

			// else {
			// 	// si esta ok ejecuto test de lectura
			// 	objRead := dataTestCRUD[prueba]
			// 	objRead.IdRecovered = data.Data["id_"+data.Object]
			// 	dataTestCRUD[prueba] = objRead
			// }

		})
	}
}
