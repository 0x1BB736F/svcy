package extract

import (
	"reflect"
	"strconv"
)

func assignInt(v string, field reflect.Value) error {
	intValue, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return err
	}

	field.SetInt(intValue)
	return nil
}

func assignUint(v string, field reflect.Value) error {
	uintValue, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return err
	}

	field.SetUint(uintValue)
	return nil
}

func assignBool(v string, field reflect.Value) error {
	boolValue, err := strconv.ParseBool(v)
	if err != nil {
		return err
	}

	field.SetBool(boolValue)
	return nil
}

func assignFloat(v string, field reflect.Value) error {
	floatValue, err := strconv.ParseFloat(v, field.Type().Bits())
	if err != nil {
		return err
	}

	field.SetFloat(floatValue)
	return nil
}

func assignString(v string, field reflect.Value) error {
	field.SetString(v)
	return nil
}
