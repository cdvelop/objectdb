package objectdb

import (
	"database/sql"

	"github.com/cdvelop/dbtools"
	"github.com/cdvelop/unixid"
)

type databaseAdapter interface {
	dbEngineAdapter
	dbtools.OrmAdapter
}

type Connection struct {
	*sql.DB
	*unixid.UnixID
	err error
	databaseAdapter
}

// DataBase Engine Adapter
type dbEngineAdapter interface {
	DataBasEngine() string    //motor base de datos Postgres,MySQL,SQLite3 etc
	DataBaseName() string     // nombre de la base de datos a conectar ej: "mydb"
	ConnectionString() string //cadena con formato de conexión base de datos dns
}
