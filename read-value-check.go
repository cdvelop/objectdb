package objectdb

import (
	"github.com/cdvelop/strings"
)

func valueIsFieldName(key, val string) (name_type bool) {

	values_key := strings.Split(key, ".")

	if len(values_key) != 2 {
		return
	}

	values_val := strings.Split(val, ".")
	if len(values_val) != 2 {
		return
	}

	if values_key[1] == values_val[1] {
		return true
	}

	return
}

func valueContainClauseAND(value string) (with_and bool) {

	if strings.Contains(value, "AND") != 0 {
		return true
	}

	return
}
