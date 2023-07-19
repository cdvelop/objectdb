package objectdb

import (
	"github.com/cdvelop/input"
	"github.com/cdvelop/model"
)

//NOTA: nombre de tablas
//error: area.staff, area-staff
//correcto: areastaff, area_staff

// nombre id
//error: idusuario, id.usuario, id-usuario
//correcto: id_usuario

const defaulTableName = "usuario"

type genero struct{}

func (genero) SourceData() map[string]string {
	return map[string]string{"D": "Dama", "V": "Varón"}
}

var (
	modelObjectForTest = map[string]*model.Object{

		defaulTableName: {
			Name: defaulTableName,
			Fields: []model.Field{
				{Name: "id_" + defaulTableName, Input: input.Number()},
				{Name: "nombre", Unique: true, Input: input.Text()},
				{Name: "apellido", Input: input.Text(), SkipCompletionAllowed: true},
				{Name: "genero", Input: input.Radio(genero{})},
			},
		},

		"especialidad": {
			Name: "especialidad",
			Fields: []model.Field{
				{Name: "id_especialidad"},
				{Name: "nombre_especialidad", Unique: true},
				{Name: "id_" + defaulTableName},
				{Name: "detalle"},
			},
		},

		"credentials": {
			Name: "credentials",
			Fields: []model.Field{
				{Name: "id_" + "credentials"},
				{Name: "id_" + defaulTableName},
				{Name: "id_especialidad"},
			},
		},
	}
)

func (c *Connection) addTestModelTablesDataBase() []*model.Object {
	return []*model.Object{
		modelObjectForTest[defaulTableName],
		modelObjectForTest["especialidad"],
		modelObjectForTest["credentials"],
	}
}
