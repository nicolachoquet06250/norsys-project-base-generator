package helpers

import (
	"fmt"
	"reflect"
)

func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

func ArrayPop[T any](s *[]T) (v *T, err error) {
	if len(*s) == 0 {
		var s T
		return &s, fmt.Errorf("can't remove last element of array")
	}
	ep := len(*s) - 1
	e := (*s)[ep]
	*s = (*s)[:ep]
	return &e, nil
}
