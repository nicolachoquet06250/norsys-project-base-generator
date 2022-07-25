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

func ArrayFilter[T any](t []T, cb func(e T, i int) bool) (result []T) {
	for i, e := range t {
		if cb(e, i) {
			result = append(result, e)
		}
	}

	return result
}

func ArrayReverse[T any](input []T) (result []T) {
	keys := make([]int, len(input))

	i := 0
	for k := range input {
		keys[i] = k
		i++
	}

	max := len(keys) - 1
	cmp := 0

	for range keys {
		result = append(result, input[max-cmp])
		cmp++
	}

	return result
}
