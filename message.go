package objectdb

import (
	"strings"
)

func filterMessageDBtoClient(err_in string, table_name string, items ...any) (err_out string) {

	if strings.Contains(err_in, `delete`) && strings.Contains(err_in, "viola la llave") {
		return "no se puede eliminar. Datos comprometidos en otra tabla"
	}

	if strings.Contains(err_in, `llave duplicada`) {
		if len(items) >= 1 {
			if data, ok := items[0].(map[string]string); ok {
				fieldError, valueError := findFieldWithError(err_in, data)
				if fieldError != "" && valueError != "" {
					return "el campo " + fieldError + " valor " + valueError + " no se puede repetir"
				}
			}
		}
	}

	return err_in
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
