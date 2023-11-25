package objectdb

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/cdvelop/timeserver"
	"github.com/cdvelop/unixid"
)

func (c *Connection) Open() *sql.DB {
	c.Set(c.databaseAdapter)
	return c.DB
}

// obtener conexión
func Get(dba databaseAdapter) *Connection {

	uid, err := unixid.NewHandler(timeserver.Add(), &sync.Mutex{}, nil)
	if err != "" {
		log.Fatal(err)
	}

	c := Connection{
		UnixID:          uid,
		databaseAdapter: dba,
	}

	c.Set(dba)

	e := c.Ping()
	if e != nil {
		log.Fatalf("¡Error ping: %v!", e)
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
