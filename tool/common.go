package tool

import (
	"reflect"
)

// 主要用于判断interface
func IsNil(i interface{}) bool {
	if i == nil {
		return true
	} else {
		vi := reflect.ValueOf(i)
		if vi.Kind() == reflect.Ptr {
			return vi.IsNil()
		}
	}
	return false
}
