package objectdb

type kv map[string]string

var dataTestCRUD map[string]dataModelDBTest

func (c *Connection) addataCrud() {

	dataTestCRUD = map[string]dataModelDBTest{

		"Luis campos correctos?": {
			defaulTableName,
			kv{"nombre": "Luis", "apellido": "de las carmenes", "genero": "V"},
			true, "", false},

		"Maria campos correctos?": {
			defaulTableName,
			kv{"nombre": "Maria", "apellido": "Ruiz", "genero": "D"},
			true, "", false},

		"Apellido en blanco Permitido?": {
			defaulTableName,
			kv{"nombre": "Arturo", "apellido": "", "genero": "V"},
			true, "", false},

		"Genero en blanco Permitido?": {
			defaulTableName,
			kv{"nombre": "Marta", "apellido": "", "genero": ""},
			false, "", false},

		"id + campos correctos?": {
			defaulTableName,
			kv{"id_" + defaulTableName: "123456", "nombre": "Juan", "apellido": "Soto", "genero": "V"},
			true, "", false},

		"genero H existe?": {
			defaulTableName,
			kv{"nombre": "Marco", "apellido": "de las carmenes", "genero": "H"},
			false, "", false},

		"apellido numérico, se requiere validación?": {
			defaulTableName,
			kv{"nombre": "Julia", "apellido": "2", "genero": "H"},
			false, "", false},

		"nombre corresponde a solo texto?": {
			defaulTableName,
			kv{"nombre": "mar1a", "apellido": "de las carmenes"},
			false, "", false},

		"todos los campos?": {
			defaulTableName,
			kv{"nombre": "Juana"},
			false, "", false},
	}

}
