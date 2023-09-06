package objectdb

import (
	"fmt"
	"log"
	"testing"

	"github.com/cdvelop/gotools"
)

func (c *Connection) readTest(data_original []map[string]string, t *testing.T) {
	var total_creations = len(data_original)

	// fmt.Println("STORE DATA: ", data_original)

	t.Run((`READ: Caso 1 leemos la base de datos con la tabla usuario y la misma cantidad de creados. 
	se esperan idénticos resultados`), func(t *testing.T) {

		data_stored, err := c.ReadObjectsInDB(defaulTableName, map[string]string{"limit": fmt.Sprint(total_creations)})
		if err != nil {
			log.Fatalln("Caso 1 error en test de lectura ", err)
		}

		if !gotools.AreSliceMapsIdentical(data_original, data_stored) {
			fmt.Println("Caso 1 error la data deberían ser idéntica:")
			fmt.Println("data original: ", data_original)
			fmt.Println()
			fmt.Println("data almacenada: ", data_stored)
			log.Fatal()
		}
	})

	t.Run((`READ: Caso 2 consulta con orden por nombre Asc, se espera un resultado con limite 2`), func(t *testing.T) {
		data_stored, err := c.ReadObjectsInDB(defaulTableName, map[string]string{"ORDER_BY": "nombre", "LIMIT": "2"})
		if err != nil {
			log.Fatalln("Caso 2 error en test de lectura ", err)
		}

		// el primer elemento debe ser Arturo
		fist_name := data_stored[0]["nombre"]
		if fist_name != "Arturo" {
			log.Fatalln("Caso 2 se esperaba como primer nombre Arturo pero se obtuvo: ", fist_name)
		}

		if gotools.AreSliceMapsIdentical(data_original, data_stored) {
			log.Fatalln("Caso 2 error la data deberían ser diferente:")
			log.Fatalln("data original: ", data_original)
			log.Fatalln("data almacenada: ", data_stored)
		}
	})

	t.Run((`READ: Caso 3 consulta con limite 1 se espera solo 1 elemento`), func(t *testing.T) {
		data_stored, err := c.ReadObjectsInDB(defaulTableName, map[string]string{"LIMIT": "1"})
		if err != nil {
			log.Fatalln("Caso 3 error en test de lectura ", err)
		}

		if len(data_stored) != 1 {
			log.Fatalln("Caso 3 error se espera solo 1 elemento pero se obtuvo: ", len(data_stored))
		}
	})

	t.Run((`READ: Caso 4 consulta por campos específicos nombre y genero`), func(t *testing.T) {
		data_stored, err := c.ReadObjectsInDB(defaulTableName, map[string]string{"SELECT": "nombre, genero"})
		if err != nil {
			log.Fatalln("Caso 4 error en test de lectura ", err)
		}

		// fmt.Println("Caso 4 SELECCIÓN ESPECIFICA: ", data_stored)

		if len(data_stored) != total_creations {
			fmt.Printf("Caso 4 error se esperaban: %v resultados pero se obtuvieron: %v\n", total_creations, len(data_stored))
			fmt.Println("RESP: ", data_stored)
			log.Fatal()
		}

		for _, data := range data_stored {
			if len(data) != 2 {
				log.Fatalf("Caso 4 error se esperaban: %v elementos pero se obtuvieron: %v\n%v", 2, len(data), data)
			}
		}
	})

	t.Run((`READ: Caso 5 consulta con mapa nulo se espera todos los datos de la tabla`), func(t *testing.T) {
		data_stored, err := c.ReadObjectsInDB(defaulTableName, nil)
		if err != nil {
			log.Fatalln("Caso 5 error en test de lectura ", err)
		}

		if !gotools.AreSliceMapsIdentical(data_original, data_stored) {
			fmt.Println("Caso 5 error la data deberían ser idéntica:")
			fmt.Println("data original: ", data_original)
			fmt.Println()
			fmt.Println("data almacenada: ", data_stored)
			log.Fatal()
		}
	})

	t.Run((`READ: Caso 6 consulta con tabla que no existe se espera error`), func(t *testing.T) {
		data_stored, err := c.ReadObjectsInDB("tabla_X", nil)
		if err == nil {
			log.Fatalln("Caso 6 se esperaba error y se obtuvo data: ", data_stored)
		}
	})

	t.Run((`READ: Caso 7 consulta de datos con 3 ids se espera 3 resultados`), func(t *testing.T) {

		var data_query = []map[string]string{
			data_original[0], data_original[1], data_original[2],
		}

		data_stored, err := c.ReadObjectsInDB(defaulTableName, data_query...)
		if err != nil {
			log.Fatalln("no se esperaba error: ", err)
		}

		if len(data_stored) != 3 {
			log.Fatalf("error se esperaban: 3 elementos pero se obtuvieron: %v\n%v", len(data_stored), data_stored)
		}

		if !gotools.AreSliceMapsIdentical(data_query, data_stored) {
			fmt.Println("error la data deberían ser idéntica:")
			fmt.Println("data original: ", data_query)
			fmt.Println()
			fmt.Println("data almacenada: ", data_stored)
			log.Fatal()
		}
	})

	t.Run((`READ: Caso 8 consulta sin información se espera error`), func(t *testing.T) {
		data_stored, err := c.ReadObjectsInDB("", nil)
		if err == nil {
			log.Fatalln("se esperaba error y se obtuvo data: ", data_stored)
		}
	})

	// t.Run((`READ: Caso 9 consulta con sql y argumentos se espera ok`), func(t *testing.T) {

	// 		// el primer elemento debe ser Arturo
	// 			fist_name := data_stored[0]["nombre"]

	// 	query := map[string]string{

	// 		"args": fist_name,
	// 	}

	// 	data_stored, err := c.ReadObjectsInDB(defaulTableName, query)
	// 	if err == nil {
	// 		log.Printf("no se esperaba error: %v", err)
	// 	}

	// 	if len(data_stored) != total_creations {
	// 		log.Printf("-se esperaban: %v registros\n-pero se obtuvieron: %v\n-%v", total_creations, len(data_stored), data_stored)
	// 	}
	// })

}
