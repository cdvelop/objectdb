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

	fmt.Println("STORE DATA: ", data_original)

	// clonamos las tablas
	for _, table := range tables {
		err := dbtools.ClonDATABLE(c, table)
		if err != nil {
			log.Fatalln("error al clonar tabla "+table.Name+" ", err)
		}
	}

	//1- leemos la base de datos con la tabla usuario y la misma cantidad de creados.
	// se esperan idénticos resultados
	data_stored, err := c.ReadObjectsInDB(defaulTableName, map[string]string{"limit": fmt.Sprint(total_creations)})
	if err != nil {
		log.Fatalln("error en test de lectura ", err)
	}

	if !gotools.AreSliceMapsIdentical(data_original, data_stored) {
		fmt.Println("error la data deberían ser idéntica:")
		fmt.Println("data original: ", data_original)
		fmt.Println()
		fmt.Println("data almacenada: ", data_stored)
		log.Fatal()
	}

	//2- consulta con orden por nombre se espera un resultado diferente
	data_stored2, err := c.ReadObjectsInDB(defaulTableName, map[string]string{"order_by": "nombre", "limit": fmt.Sprint(total_creations)})
	if err != nil {
		log.Fatalln("error en test de lectura ", err)
	}

	// el primer elemento debe ser Arturo
	fist_name := data_original[0]["nombre"]
	if fist_name != "Arturo" {
		log.Fatalln(" se esperaba como primer nombre Arturo pero se obtuvo: ", fist_name)
	}

	if gotools.AreSliceMapsIdentical(data_original, data_stored2) {
		log.Fatalln("error la data deberían ser diferente:")
		log.Fatalln("data original: ", data_original)
		log.Fatalln("data almacenada: ", data_stored2)
	}

	//3- consulta con limite 2 se espera solo dos elementos
	data_stored, err = c.ReadObjectsInDB(defaulTableName, map[string]string{"limit": fmt.Sprint(total_creations)})
	if err != nil {
		log.Fatalln("error en test de lectura ", err)
	}
	fmt.Println("DATA ORDENADA POR NOMBRE: ", data_stored)

	// var testData = map[string]struct {
	// 	data   kv
	// 	expect int
	// }{
	// 	"consulta por todos los elemento que fueron cambiados en la tabla " + defaulTableName: {data: kv{"apellido": "NUEVO APELLIDO"}, expect: total_creations},
	// }

	// fmt.Println("CREACIONES: ", total_creations)

	// for testName, d := range testData {
	// 	t.Run(testName, func(t *testing.T) {

	// 		out, err := c.ReadObjectsInDB(defaulTableName, d.data)
	// 		if err != nil {
	// 			log.Fatalln("error en test de lectura ", err, d)
	// 		}

	// 		if len(out) != d.expect {
	// 			log.Fatalf("Para entrada '%v', se esperaba '%v' pero se obtuvo '%v'", d.data, d.expect, len(out))
	// 		}

	// 		fmt.Println(d)

	// 	})
	// }

}

// for _, data := range dataTestCRUD {

// 	if data.ExpectedError == "" { //solo los casos de éxito

// 		t.Run(("READ: "), func(t *testing.T) {
// 			out, err := c.ReadObjectsInDB(defaulTableName, map[string]string{"id_" + defaulTableName: data.IdRecovered})
// 			if err != nil {
// 				log.Fatalln("error en test de lectura ", err, data)
// 			}

// 			if len(out) == 0 {
// 				log.Fatalf("!!! READ data: [%v] resp\n", out)
// 			}

// 			// fmt.Println("=> DATA CAMBIADA?:", out)
// 			for _, o := range out {

// 				if !strings.Contains(o["apellido"], "NUEVO APELLIDO") {
// 					log.Fatalln("ERROR APELLIDO NUEVO NO CAMBIADO SALIDA:\n", out)
// 				}
// 			}

// 		})
// 	}
// }
