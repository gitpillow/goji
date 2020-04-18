package goji

import (
	"fmt"
	"reflect"
)

func StringSliceContains(strs []string, str string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}
	return false
}

func StringSliceFindFirst(strs []string, str string) (int, bool) {
	for i, s := range strs {
		if s == str {
			return i, true
		}
	}
	return -1, false
}

func SliceForEach(sli interface{}, f func(i int, item interface{}) error) error {
	t := reflect.TypeOf(sli)
	k := t.Kind()
	if k == reflect.Slice {
		v := reflect.ValueOf(sli)
		for i := 0; i < v.Len(); i++ {
			err := f(i, v.Index(i).Interface())
			if err != nil {
				return err
			}
		}
		return nil
	}
	return fmt.Errorf("for each on not slice value: %v -> %v", t.Name(), sli)
}
