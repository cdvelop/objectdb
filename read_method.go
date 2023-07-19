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
func (c Connection) QueryOne(sql string, args ...interface{}) (map[string]string, error) {
	c.Open()
	defer c.Close()

	rows, err := c.Query(sql, args...)
	if err != nil {
		return nil, err
	}

	return dbtools.FetchOne(rows)
}

// QueryAll .
func (c Connection) QueryAll(sql string, args ...interface{}) ([]map[string]string, error) {
	c.Open()
	defer c.Close()

	rows, err := c.Query(sql, args...)
	if err != nil {
		return nil, err
	}

	return dbtools.FetchAll(rows)
}
