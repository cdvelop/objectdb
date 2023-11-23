package objectdb

import (
	"fmt"
	"log"

	"github.com/cdvelop/dbtools"
)

// QueryWithoutANSWER SinResultado ejecuta sql en bd con sin respuesta de mas de 1 operación
// recibe sql y mensaje a mostrar en consola
func (c Connection) QueryWithoutANSWER(sql, mje string) bool {
	c.Open()
	defer c.Close()

	_, err := c.Exec(sql)
	if err != nil {
		log.Printf("%v %v", err, sql)
		return false
	}
	if mje != "" {
		fmt.Println(mje)
	}

	return true
}

// QueryOne .
// https://my.oschina.net/nowayout/blog/139398
func (c Connection) QueryOne(sql string, args ...interface{}) (out map[string]string, err string) {
	c.Open()
	defer c.Close()

	rows, e := c.Query(sql, args...)
	if e != nil {
		return nil, "QueryOne error " + e.Error()
	}

	return dbtools.FetchOne(rows)
}

// QueryAll .
func (c Connection) QueryAll(sql string, args ...interface{}) (out []map[string]string, err string) {
	c.Open()
	defer c.Close()

	rows, e := c.Query(sql, args...)
	if e != nil {
		return nil, "QueryAll error " + e.Error()
	}

	return dbtools.FetchAll(rows)
}
