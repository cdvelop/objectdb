package objectdb

import (
	"github.com/cdvelop/dbtools"
)

// QueryWithoutANSWER SinResultado ejecuta sql en bd con sin respuesta de mas de 1 operación
// recibe sql y mensaje a mostrar en consola
func (c Connection) QueryWithoutANSWER(sql string) (err string) {
	c.Open()
	defer c.Close()

	_, er := c.Exec(sql)
	if er != nil {
		return "QueryWithoutANSWER " + er.Error()
	}

	return
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
