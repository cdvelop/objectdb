package objectdb

import (
	"fmt"
	"strings"
)

func (c Connection) filterMessageDBtoClient(dbErr string, table_name string, out_message *string, all_data ...map[string]string) (message, wrong_field string) {

	if strings.Contains(dbErr, `delete`) && strings.Contains(dbErr, "viola la llave") {
		message = "¡Error No se puede eliminar!. Datos comprometidos en otra tabla"
		return
	}

	if strings.Contains(dbErr, `llave duplicada`) {
		for _, data := range all_data {
			fieldError, valueError := c.findFieldWithError(dbErr, data)
			if fieldError != "" && valueError != "" {
				message = fmt.Sprintf("¡Error en el campo %v valor %v no se puede repetir!", fieldError, valueError)
				wrong_field = fieldError
				return
			}
		}
	}

	message = dbErr
	// log.Println(dbErr)

	return
}

func (c Connection) findFieldWithError(dbErr string, object map[string]string) (fieldError, valueError string) {
	for key, value := range object {
		if strings.Contains(dbErr, key) {
			fieldError = key
			valueError = value
			break
		}
	}
	return
}
