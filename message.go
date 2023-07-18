package objectdb

import (
	"fmt"
	"strings"
)

func filterMessageDBtoClient(err error, table_name string, all_data ...*map[string]string) error {

	if strings.Contains(err.Error(), `delete`) && strings.Contains(err.Error(), "viola la llave") {
		return fmt.Errorf("¡Error No se puede eliminar!. Datos comprometidos en otra tabla")
	}

	if strings.Contains(err.Error(), `llave duplicada`) {
		for _, data := range all_data {
			fieldError, valueError := findFieldWithError(err.Error(), data)
			if fieldError != "" && valueError != "" {
				return fmt.Errorf("¡Error en el campo %v valor %v no se puede repetir!", fieldError, valueError)
			}
		}
	}

	return err
}

func findFieldWithError(dbErr string, object *map[string]string) (fieldError, valueError string) {
	if object != nil {
		for key, value := range *object {
			if strings.Contains(dbErr, key) {
				fieldError = key
				valueError = value
				break
			}
		}
	}
	return
}
