package objectdb

import (
	"fmt"
	"strings"
)

func filterMessageDBtoClient(err error, table_name string, items ...any) error {

	if strings.Contains(err.Error(), `delete`) && strings.Contains(err.Error(), "viola la llave") {
		return fmt.Errorf("error No se puede eliminar. Datos comprometidos en otra tabla")
	}

	if strings.Contains(err.Error(), `llave duplicada`) {
		if len(items) >= 1 {
			if data, ok := items[0].(map[string]string); ok {
				fieldError, valueError := findFieldWithError(err.Error(), data)
				if fieldError != "" && valueError != "" {
					return fmt.Errorf("error en el campo %v valor %v no se puede repetir", fieldError, valueError)
				}
			}
		}
	}

	return err
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
