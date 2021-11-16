package extract

import (
	"reflect"
)

// TODO: add documentation

func extractRequest(v interface{}, getter StringGetter, tag string) error {

	if reflect.TypeOf(v).Kind() != reflect.Ptr {
		return ErrNotPointer
	}

	value := reflect.Indirect(reflect.ValueOf(v)).Interface()

	structValue := reflect.ValueOf(v).Elem()
	structType := reflect.TypeOf(value)

	var field reflect.StructField

	var tagValue string
	var gotValue string

	var err error

	for i := 0; i < structType.NumField(); i++ {
		field = structType.Field(i)

		tagValue = field.Tag.Get(tag)

		gotValue = getter.Get(tagValue)

		if gotValue == "" {
			continue
		}

		fn, ok := assignMap[field.Type.Kind()]
		if !ok {
			continue
		}

		err = fn(gotValue, structValue.Field(i))
		if err != nil {
			return err
		}

	}

	return nil
}
