package objectdb

import (
	"fmt"
	"log"
	"testing"
)

func (c Connection) createTest(t *testing.T) {

	for prueba, data := range dataTestCRUD {
		t.Run(("CREATE: " + prueba + " " + fmt.Sprint(data.Result)), func(t *testing.T) {

			object, exist := modelObjectForTest[data.Object]
			if !exist {
				log.Fatalf("objeto: %v no existe", data.Object)
			}

			// validar elemento aquí
			if mg, ok := object.ValidateData(true, data.Data); !ok {
				data.IdRecovered = mg
				return
			}

			mg, ok := c.CreateObjects(data.Object, data.Data)
			if ok != data.Result {
				log.Fatalf("Error: [%v]\n", mg)
			} else {
				// si esta ok ejecuto test de lectura
				objRead := dataTestCRUD[prueba]
				objRead.IdRecovered = data.Data["id_"+data.Object]
				dataTestCRUD[prueba] = objRead
			}

		})
	}
}
