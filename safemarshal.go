package safemarshal

import (
	"reflect"
)

// Check returns true if the type of v can be safely JSON marshalled.
func Check(v any) bool {
	return !unsafe(reflect.TypeOf(v), nil)
}

func unsafe(t reflect.Type, visited []reflect.Type) bool {
	for i := range visited {
		if t == visited[i] {
			return false
		}
	}

	copied := make([]reflect.Type, len(visited)+1)
	copy(copied, visited)
	copied[len(visited)] = t

	switch t.Kind() {
	case reflect.String:
		fallthrough
	case reflect.Bool:
		fallthrough
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fallthrough
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fallthrough
	case reflect.Float32, reflect.Float64:
		return false
	case reflect.Array, reflect.Slice:
		fallthrough
	case reflect.Pointer:
		return unsafe(t.Elem(), copied)
	case reflect.Map:
		if unsafe(t.Key(), copied) {
			return true
		}
		return unsafe(t.Elem(), copied)
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			if unsafe(t.Field(i).Type, copied) {
				return true
			}
		}
		return false
	default:
		return true
	}
}
