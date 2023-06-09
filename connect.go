package objectdb

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/cdvelop/dbtools"
)

func (c *Connection) Open() {
	c.Set(c.databaseAdapter)
}

// obtener conexión
func Get(dba databaseAdapter) *Connection {

	c := Connection{
		UnixID:          dbtools.NewUnixIdHandler(),
		databaseAdapter: dba,
	}

	c.Set(dba)

	err := c.Ping()
	if err != nil {
		log.Fatalf("¡Error ping: %v!", err)
	}

	fmt.Printf("*** Conexión DB: %v Establecida, Engine: %v ***\n", dba.DataBaseName(), c.DataBasEngine())
	defer c.Close()

	return &c
}

// setear conexión base de datos cerrar después de usar
func (c *Connection) Set(dba databaseAdapter) {
	// db, err = sql.Open(dns.DataBasEngine, fmt.Sprintf("%v://%v:%v@%v:%v/%v?sslmode=disable", dns.DataBasEngine, dns.UserDatabase, dns.PasswordDatabase, dns.IPLocalServer, dns.DataBasePORT, dns.DataBaseName))

	c.DB, c.err = sql.Open(dba.DataBasEngine(), dba.ConnectionString())
	if c.err != nil {
		log.Fatalf("¡Error al abrir conexión db %v!", c.err)
	}

	// alexedwards.net/blog/configuring-sqldb - odb.SetMaxOpenConns(25) - odb.SetMaxIdleConns(25) - odb.SetConnMaxLifetime(5 * time.Minute)

}
