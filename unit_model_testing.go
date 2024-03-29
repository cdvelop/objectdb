package objectdb

import (
	"github.com/cdvelop/input"
	"github.com/cdvelop/model"
	"github.com/cdvelop/unixid"
)

//NOTA: nombre de tablas
//error: area.staff, area-staff
//correcto: areastaff, area_staff

// nombre id
//error: idusuario, id.usuario, id-usuario
//correcto: id_usuario

const defaulTableName = "usuario"

var (
	modelObjectForTest = map[string]*model.Object{

		defaulTableName: {
			ObjectName: defaulTableName,
			Table:      defaulTableName,
			Fields: []model.Field{
				{Name: "id_" + defaulTableName, Input: unixid.InputPK()},
				{Name: "nombre", Unique: true, Input: input.TextOnly()},
				{Name: "apellido", Input: input.Text(), SkipCompletionAllowed: true},
				{Name: "genero", Input: input.RadioGender()},
			},
		},

		"especialidad": {
			ObjectName: "especialidad",
			Table:      "especialidad",
			Fields: []model.Field{
				{Name: "id_especialidad"},
				{Name: "nombre_especialidad", Unique: true},
				{Name: "id_" + defaulTableName},
				{Name: "detalle"},
			},
		},

		"credentials": {
			ObjectName: "credentials",
			Table:      "credentials",
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
