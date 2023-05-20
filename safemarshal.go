package safemarshal

import (
	"reflect"
)

func Check(v any) bool {
	return !unsafe(reflect.TypeOf(v))
}

// Methods applicable only to some types, depending on Kind.
// The methods allowed for each kind are:
//
//	Int*, Uint*, Float*, Complex*: Bits
//	Array: Elem, Len
//	Chan: ChanDir, Elem
//	Func: In, NumIn, Out, NumOut, IsVariadic.
//	Map: Key, Elem
//	Pointer: Elem
//	Slice: Elem
//	Struct: Field, FieldByIndex, FieldByName, FieldByNameFunc, NumField

func unsafe(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.String:
		fallthrough
	case reflect.Bool:
		fallthrough
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		fallthrough
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		fallthrough
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		return false
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		fallthrough
	case reflect.Pointer:
		return unsafe(t.Elem())
	case reflect.Map:
		if unsafe(t.Key()) {
			return true
		}
		return unsafe(t.Elem())
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			if unsafe(t.Field(i).Type) {
				return true
			}
		}
		return false
	default:
		return true
	}
}
