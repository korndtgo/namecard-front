package covstr2any

import (
	"reflect"
	"strconv"
)

// if have ptr you can use only one level
// support types [string, float64, int, bool]
func ConvertStringToAnyType(str string, targetType reflect.Type) (reflect.Value, error) {
	var isPtr bool
	if targetType.Kind() == reflect.Ptr {
		isPtr = true
		targetType = targetType.Elem()
	}

	var value reflect.Value
	// add type for support
	switch targetType.Kind() {
	case reflect.String:
		tmp := str
		value = reflect.ValueOf(&tmp)
	case reflect.Float64:
		tmp, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return reflect.Value{}, err
		}
		value = reflect.ValueOf(&tmp)
	case reflect.Int:
		tmp, err := strconv.Atoi(str)
		if err != nil {
			return reflect.Value{}, err
		}
		value = reflect.ValueOf(&tmp)
	case reflect.Bool:
		tmp, err := strconv.ParseBool(str)
		if err != nil {
			return reflect.Value{}, err
		}
		value = reflect.ValueOf(&tmp)
	}

	if !isPtr {
		value = value.Elem()
	}
	return value, nil
}
