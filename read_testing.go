package objectdb

import (
	"fmt"
	"log"
	"testing"

	"github.com/cdvelop/dbtools"
	"github.com/cdvelop/gotools"
	"github.com/cdvelop/model"
)

func (c *Connection) readTest(tables []*model.Object, t *testing.T) {
	var total_creations int
	var data_original []map[string]string

	for _, d := range dataTestCRUD {
		if d.ExpectedError == "" { // solo los casos de éxito
			total_creations++
			data_original = append(data_original, d.Data)
		}
	}

	// fmt.Println("STORE DATA: ", data_original)

	// clonamos las tablas
	for _, table := range tables {
		err := dbtools.ClonDATABLE(c, table)
		if err != nil {
			log.Fatalln("error al clonar tabla "+table.Name+" ", err)
		}
	}

	//Caso 1 leemos la base de datos con la tabla usuario y la misma cantidad de creados.
	// se esperan idénticos resultados
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

	//Caso 2 consulta con orden por nombre Asc, se espera un resultado con limite 2
	data_stored2, err := c.ReadObjectsInDB(defaulTableName, map[string]string{"order_by": "nombre", "limit": "2"})
	if err != nil {
		log.Fatalln("Caso 2 error en test de lectura ", err)
	}

	// el primer elemento debe ser Arturo
	fist_name := data_stored2[0]["nombre"]
	if fist_name != "Arturo" {
		log.Fatalln("Caso 2 se esperaba como primer nombre Arturo pero se obtuvo: ", fist_name)
	}

	if gotools.AreSliceMapsIdentical(data_original, data_stored2) {
		log.Fatalln("Caso 2 error la data deberían ser diferente:")
		log.Fatalln("data original: ", data_original)
		log.Fatalln("data almacenada: ", data_stored2)
	}

	//Caso 3 consulta con limite 1 se espera solo 1 elemento
	data_stored3, err := c.ReadObjectsInDB(defaulTableName, map[string]string{"limit": "1"})
	if err != nil {
		log.Fatalln("Caso 3 error en test de lectura ", err)
	}

	if len(data_stored3) != 1 {
		log.Fatalln("Caso 3 error se espera solo 1 elemento pero se obtuvo: ", len(data_stored3))
	}

	//Caso 4 consulta por campos específicos nombre y genero
	data_stored4, err := c.ReadObjectsInDB(defaulTableName, map[string]string{"choose": "nombre, genero"})
	if err != nil {
		log.Fatalln("Caso 4 error en test de lectura ", err)
	}

	// fmt.Println("Caso 4 SELECCIÓN ESPECIFICA: ", data_stored4)

	if len(data_stored4) != total_creations {
		log.Fatalf("Caso 4 error se esperaban: %v resultados pero se obtuvieron: %v", total_creations, len(data_stored4))
	}

	for _, data := range data_stored4 {
		if len(data) != 2 {
			log.Fatalf("Caso 4 error se esperaban: %v elementos pero se obtuvieron: %v\n%v", 2, len(data), data)
		}
	}

	// Caso 5 consulta con mapa nulo se espera todos los datos de la tabla
	data_stored5, err := c.ReadObjectsInDB(defaulTableName, nil)
	if err != nil {
		log.Fatalln("Caso 5 error en test de lectura ", err)
	}

	if !gotools.AreSliceMapsIdentical(data_original, data_stored5) {
		fmt.Println("Caso 5 error la data deberían ser idéntica:")
		fmt.Println("data original: ", data_original)
		fmt.Println()
		fmt.Println("data almacenada: ", data_stored5)
		log.Fatal()
	}

	//Caso 6 consulta con tabla que no existe se espera error
	data_stored6, err := c.ReadObjectsInDB("tabla_X", nil)
	if err == nil {
		log.Fatalln("Caso 6 se esperaba error y se obtuvo data: ", data_stored6)
	}

}
