package objectdb

import (
	"fmt"
	"log"
	"testing"

	"github.com/cdvelop/maps"
	"github.com/cdvelop/model"
)

func (c *Connection) readTest(data_original []map[string]string, t *testing.T) {
	var total_creations = len(data_original)

	// fmt.Println("STORE DATA: ", data_original)

	t.Run((`READ: Caso 1 leemos la base de datos con la tabla usuario y la misma cantidad de creados. 
	se esperan idénticos resultados`), func(t *testing.T) {

		data_stored, err := c.ReadSyncDataDB(model.ReadParams{
			FROM_TABLE: defaulTableName, LIMIT: total_creations})
		if err != "" {
			t.Fatal("Caso 1 error en test de lectura ", err)
			return
		}

		if !maps.AreSliceMapsIdentical(data_original, data_stored) {
			fmt.Println("Caso 1 error la data deberían ser idéntica:")
			fmt.Println("data original: ", data_original)
			fmt.Println()
			fmt.Println("data almacenada: ", data_stored)
			t.Fatal()
			return
		}
	})

	t.Run((`READ: Caso 2 consulta con orden por nombre Asc, se espera un resultado con limite 2`), func(t *testing.T) {
		data_stored, err := c.ReadSyncDataDB(
			model.ReadParams{
				FROM_TABLE: defaulTableName,
				ORDER_BY:   "nombre",
				LIMIT:      2,
			})
		if err != "" {
			t.Fatal("Caso 2 error en test de lectura ", err)
			return
		}

		// el primer elemento debe ser Arturo
		fist_name := data_stored[0]["nombre"]
		if fist_name != "Arturo" {
			t.Fatal("Caso 2 se esperaba como primer nombre Arturo pero se obtuvo: ", fist_name)
		}

		if maps.AreSliceMapsIdentical(data_original, data_stored) {
			t.Fatal("Caso 2 error la data deberían ser diferente:")
			t.Fatal("data original: ", data_original)
			t.Fatal("data almacenada: ", data_stored)
			return
		}
	})

	t.Run((`READ: Caso 3 consulta con limite 1 se espera solo 1 elemento`), func(t *testing.T) {
		data_stored, err := c.ReadSyncDataDB(model.ReadParams{
			FROM_TABLE: defaulTableName,
			LIMIT:      1,
		}, map[string]string{"LIMIT": "1"})
		if err != "" {
			t.Fatal("Caso 3 error en test de lectura ", err)
			return
		}

		if len(data_stored) != 1 {
			t.Fatal("Caso 3 error se espera solo 1 elemento pero se obtuvo: ", len(data_stored))
			return
		}
	})

	t.Run((`READ: Caso 4 consulta por campos específicos nombre y genero`), func(t *testing.T) {
		data_stored, err := c.ReadSyncDataDB(model.ReadParams{
			FROM_TABLE: defaulTableName,
			SELECT:     "nombre, genero",
		})
		if err != "" {
			t.Fatal("Caso 4 error en test de lectura ", err)
			return
		}

		// fmt.Println("Caso 4 SELECCIÓN ESPECIFICA: ", data_stored)

		if len(data_stored) != total_creations {
			fmt.Printf("Caso 4 error se esperaban: %v resultados pero se obtuvieron: %v\n", total_creations, len(data_stored))
			fmt.Println("RESP: ", data_stored)
			log.Fatal()
			return
		}

		for _, data := range data_stored {
			if len(data) != 2 {
				t.Fatalf("Caso 4 error se esperaban: %v elementos pero se obtuvieron: %v\n%v", 2, len(data), data)
				return
			}
		}
	})

	t.Run((`READ: Caso 5 consulta con mapa nulo se espera todos los datos de la tabla`), func(t *testing.T) {
		data_stored, err := c.ReadSyncDataDB(model.ReadParams{
			FROM_TABLE: defaulTableName})
		if err != "" {
			t.Fatal("Caso 5 error en test de lectura ", err)
			return
		}

		if !maps.AreSliceMapsIdentical(data_original, data_stored) {
			fmt.Println("Caso 5 error la data deberían ser idéntica:")
			fmt.Println("data original: ", data_original)
			fmt.Println()
			fmt.Println("data almacenada: ", data_stored)
			log.Fatal()
		}
	})

	t.Run((`READ: Caso 6 consulta con tabla que no existe se espera error`), func(t *testing.T) {
		data_stored, err := c.ReadSyncDataDB(
			model.ReadParams{
				FROM_TABLE: "tabla_X"})
		if err == "" {
			t.Fatal("Caso 6 se esperaba error y se obtuvo data: ", data_stored)
			return
		}
	})

	t.Run((`READ: Caso 7 consulta de datos con 3 ids se espera 3 resultados`), func(t *testing.T) {

		var data_query = []map[string]string{
			data_original[0], data_original[1], data_original[2],
		}

		data_stored, err := c.ReadSyncDataDB(model.ReadParams{
			FROM_TABLE: defaulTableName}, data_query...)
		if err != "" {
			t.Fatal("no se esperaba error: ", err)
			return
		}

		if len(data_stored) != 3 {
			fmt.Printf("error se esperaban: 3 elementos pero se obtuvieron: %v\n\n", len(data_stored))

			for _, v := range data_stored {
				fmt.Println(v)
				fmt.Println()
			}
			log.Fatal()
		}

		if !maps.AreSliceMapsIdentical(data_query, data_stored) {
			fmt.Println("error la data deberían ser idéntica:")
			fmt.Println("data original: ", data_query)
			fmt.Println()
			fmt.Println("data almacenada: ", data_stored)
			log.Fatal()
		}
	})

	t.Run((`READ: Caso 8 consulta sin información se espera error`), func(t *testing.T) {
		data_stored, err := c.ReadSyncDataDB(model.ReadParams{}, nil)
		if err == "" {
			t.Fatal("se esperaba error y se obtuvo data: ", data_stored)
			return
		}
	})

	// t.Run((`READ: Caso 9 consulta con sql y argumentos se espera ok`), func(t *testing.T) {

	// 		// el primer elemento debe ser Arturo
	// 			fist_name := data_stored[0]["nombre"]

	// 	query := map[string]string{

	// 		"args": fist_name,
	// 	}

	// 	data_stored, err := c.ReadSyncDataDB(defaulTableName, query)
	// 	if err == nil {
	// 		log.Printf("no se esperaba error: %v", err)
	// 	}

	// 	if len(data_stored) != total_creations {
	// 		log.Printf("-se esperaban: %v registros\n-pero se obtuvieron: %v\n-%v", total_creations, len(data_stored), data_stored)
	// 	}
	// })

}
