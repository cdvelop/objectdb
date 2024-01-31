package objectdb

import (
	"strings"
	"testing"
)

func (c Connection) updateTest(t *testing.T) {

	for _, data := range dataTestCRUD {

		if data.ExpectedError == "" { //solo los casos que no contienen error

			// fmt.Println("ID PARA ACTUALIZAR: ", data.Data["id_"+defaulTableName])

			data.Data["apellido"] = "NUEVO APELLIDO"

			name := data.Data["nombre"]
			//removemos nombre ya que es un campo único en el modelo y no se puede actualizar
			delete(data.Data, "nombre")

			t.Run(("UPDATE: apellido de: " + name), func(t *testing.T) {

				object, exist := modelObjectForTest[data.Object]
				if !exist {
					t.Fatalf("objeto: %v no existe", data.Object)
					return
				}

				// fmt.Println("DATA A ACTUALIZAR: ", data.Data)

				// validar elemento aquí
				err := object.ValidateData(false, true, data.Data)
				if err != "" {
					t.Fatal(err)
					return
				}
				// fmt.Println("=> DATA A ACTUALIZAR: ", data.Data)
				err = c.UpdateObjectsInDB(defaulTableName, false, data.Data)
				if err != "" {
					if err != data.ExpectedError {
						t.Fatalf("en objeto: [%v]\n=>la expectativa es: [%v]\n=>pero se obtuvo: [%v]\n%v", data.Object, data.ExpectedError, err, data.Object)
						return
					}
				}

				// fmt.Println("DATA PARA NOTIFICACIÓN DE ACTUALIZACIÓN ", notify_data)

			})
		}
	}

	// chequear si se realizo la actualización
	for _, data := range dataTestCRUD {

		if data.ExpectedError == "" { //solo los casos de éxito

			t.Run(("UPDATE READ CHECK: "), func(t *testing.T) {

				out, err := c.ReadDataDB(struct {
					FROM_TABLE    string
					SELECT        string
					WHERE         []map[string]string
					AND_CONDITION bool
					ID            string
					ORDER_BY      string
					SORT_DESC     bool
					LIMIT         int
					RETURN_ANY    bool
				}{
					FROM_TABLE: defaulTableName,
					WHERE:      []map[string]string{{"id_" + defaulTableName: data.Data["id_"+defaulTableName]}},
				}, nil)
				if err != "" {
					t.Fatal("error en test de lectura ", err, data)
					return
				}

				if len(out) == 0 {
					t.Fatalf("!!! READ data: [%v] resp\n", out)
					return
				}

				// fmt.Println("=> DATA CAMBIADA?:", out)
				for _, o := range out {
					if !strings.Contains(o["apellido"], "NUEVO APELLIDO") {
						t.Fatal("ERROR APELLIDO NUEVO NO CAMBIADO SALIDA:\n", out)
						return
					}
				}
			})
		}
	}
}
