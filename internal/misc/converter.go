package misc

import (
	"fmt"
	"strconv"
)

var errConversionError = func(v interface{}) error {
	return fmt.Errorf("cannot convert value %v (type %T) to integer", v, v)
}

func ToInteger(v interface{}) (int, error) {
	i, err := strconv.Atoi(v.(string))
	if err != nil {
		return i, errConversionError(v)
	}

	return i, nil
}
