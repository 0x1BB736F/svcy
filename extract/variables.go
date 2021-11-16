package extract

import "reflect"

var assignMap = map[reflect.Kind](func(string, reflect.Value) error){
	reflect.Int:   assignInt,
	reflect.Int64: assignInt,
	reflect.Int32: assignInt,
	reflect.Int16: assignInt,
	reflect.Int8:  assignInt,

	reflect.Uint:   assignUint,
	reflect.Uint64: assignUint,
	reflect.Uint32: assignUint,
	reflect.Uint16: assignUint,
	reflect.Uint8:  assignUint,

	reflect.Bool: assignBool,

	reflect.Float32: assignFloat,
	reflect.Float64: assignFloat,
}
