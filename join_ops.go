package array

import (
	"errors"
	"reflect"
)

// Intersection takes two array/slices and finds the common values between the
// two arrays. The returned type will be of the first variable, and is
// gaurenteed to return no values if the two arrays have a dissimilar type.
func Intersection(a interface{}, b interface{}) (interface{}, error) {
	aVal := reflect.ValueOf(a)
	aType := aVal.Type()
	bType := reflect.TypeOf(b)

	if aType.Kind() != reflect.Slice || bType.Kind() != reflect.Slice {
		return reflect.Zero(aType),
			errors.New("First and second arguments must be slices")
	}

	newSlice := reflect.MakeSlice(aType, 0, 0)

	for i := 0; i < aVal.Len(); i++ {
		if In(aVal.Index(i).Interface(), b) {
			newSlice = reflect.Append(newSlice, aVal.Index(i))
		}
	}

	return newSlice.Interface(), nil
}

// Except takes two array/slices and finds values that are only in the first
// slice.
func Except(a interface{}, b interface{}) (interface{}, error) {
	aVal := reflect.ValueOf(a)
	aType := aVal.Type()
	bType := reflect.TypeOf(b)

	if aType.Kind() != reflect.Slice || bType.Kind() != reflect.Slice {
		return reflect.Zero(aType),
			errors.New("First and second arguments must be slices")
	}

	newSlice := reflect.MakeSlice(aType, 0, 0)

	for i := 0; i < aVal.Len(); i++ {
		if !In(aVal.Index(i).Interface(), b) {
			newSlice = reflect.Append(newSlice, aVal.Index(i))
		}
	}

	return newSlice.Interface(), nil
}
