package objectdb

import (
	"database/sql"
)

type databaseAdapter interface {
	dbEngineAdapter
	crudOrmMethods
}

type Connection struct {
	*sql.DB
	err error
	databaseAdapter
}

// DataBase Engine Adapter
type dbEngineAdapter interface {
	DataBasEngine() string    //motor base de datos Postgres,MySQL,SQLite3 etc
	DataBaseName() string     // nombre de la base de datos a conectar ej: "mydb"
	ConnectionString() string //cadena con formato de conexión base de datos dns
	GetNewID() string         // función que genera nuevo id exclusivo para la db elegida
}

type crudOrmMethods interface {
	//ej postgres:"$1", sqlite: "?"
	PlaceHolders(index ...uint8) string
}
