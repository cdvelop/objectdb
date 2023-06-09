package objectdb

import (
	"fmt"
	"strings"
)

func filterMessageDBtoClient(dbErr string, table_name string, out_message *string, all_data ...map[string]string) {

	if strings.Contains(dbErr, `delete`) && strings.Contains(dbErr, "viola la llave") {
		*out_message = "¡Error No se puede eliminar!. Datos comprometidos en otra tabla"
		return
	}

	if strings.Contains(dbErr, `llave duplicada`) {
		for _, data := range all_data {
			fieldError, valueError := findFieldWithError(dbErr, data)
			if fieldError != "" && valueError != "" {
				*out_message = fmt.Sprintf("¡Error en el campo %v valor %v no se puede repetir!", fieldError, valueError)
				return
			}
		}
	}

	*out_message = dbErr
}

func findFieldWithError(dbErr string, object map[string]string) (fieldError, valueError string) {
	for key, value := range object {
		if strings.Contains(dbErr, key) {
			fieldError = key
			valueError = value
			break
		}
	}
	return
}
