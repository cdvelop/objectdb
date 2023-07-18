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
			err := object.ValidateData(true, false, &data.Data)
			if err != nil {
				data.IdRecovered = err.Error()
				return
			}

			err = c.CreateObjects(data.Object, data.Data)
			if err.Error() != data.ExpectedError {
				log.Fatalf("Error esperado: [%v] pero se obtuvo: [%v]\n%v", data.ExpectedError, err, data.Object)
			} else {
				// si esta ok ejecuto test de lectura
				objRead := dataTestCRUD[prueba]
				objRead.IdRecovered = data.Data["id_"+data.Object]
				dataTestCRUD[prueba] = objRead
			}

		})
	}
}
