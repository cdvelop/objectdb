package objectdb

type kv map[string]string

var dataTestCRUD map[string]dataModelDBTest

func (c *Connection) addataCrud() {

	dataTestCRUD = map[string]dataModelDBTest{

		"Luis campos correctos?": {
			defaulTableName,
			kv{"nombre": "Luis", "apellido": "de las carmenes", "genero": "V"},
			""},

		"Maria campos correctos?": {
			defaulTableName,
			kv{"nombre": "Maria", "apellido": "Ruiz", "genero": "D"},
			""},

		"Apellido en blanco Permitido?": {
			defaulTableName,
			kv{"nombre": "Arturo", "apellido": "", "genero": "V"},
			""},

		"Genero en blanco Permitido?": {
			defaulTableName,
			kv{"nombre": "Marta", "apellido": "", "genero": ""},
			"ERROR"},

		"id + campos correctos?": {
			defaulTableName,
			kv{"id_" + defaulTableName: "123456", "nombre": "Juan", "apellido": "Soto", "genero": "V"},
			""},

		"genero H existe?": {
			defaulTableName,
			kv{"nombre": "Marco", "apellido": "de las carmenes", "genero": "H"},
			"ERROR"},

		"apellido numérico, se requiere validación?": {
			defaulTableName,
			kv{"nombre": "Julia", "apellido": "2", "genero": "H"},
			"ERROR"},

		"nombre corresponde a solo texto?": {
			defaulTableName,
			kv{"nombre": "mar1a", "apellido": "de las carmenes", "genero": "D"},
			"ERROR"},

		"todos los campos?": {
			defaulTableName,
			kv{"nombre": "Juana"},
			"ERROR"},
	}

}
