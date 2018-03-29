package array

import (
	"reflect"
)

// getABValType is a helper that saves LOCs
func getABValType(a, b interface{}) (reflect.Value, reflect.Type, reflect.Value, reflect.Type) {
	aVal := reflect.ValueOf(a)
	aType := aVal.Type()
	bVal := reflect.ValueOf(b)
	bType := bVal.Type()
	return aVal, aType, bVal, bType
}

// Intersection takes two array/slices and finds the common values between the
// two arrays. The returned type will be of the first variable, and is
// gaurenteed to return no values if the two arrays have a dissimilar type.
func Intersection(a, b interface{}) interface{} {
	aVal, aType, _, bType := getABValType(a, b)

	if aType.Kind() != reflect.Slice || bType.Kind() != reflect.Slice {
		return reflect.Zero(aType)
	}

	newSlice := reflect.MakeSlice(aType, 0, 0)

	for i := 0; i < aVal.Len(); i++ {
		if In(aVal.Index(i).Interface(), b) {
			newSlice = reflect.Append(newSlice, aVal.Index(i))
		}
	}

	return newSlice.Interface()
}

// Except takes two array/slices and finds values that are only in the first
// slice.
func Except(a, b interface{}) interface{} {
	aVal, aType, _, bType := getABValType(a, b)

	if aType.Kind() != reflect.Slice || bType.Kind() != reflect.Slice {
		return reflect.Zero(aType)
	}

	newSlice := reflect.MakeSlice(aType, 0, 0)

	for i := 0; i < aVal.Len(); i++ {
		if !In(aVal.Index(i).Interface(), b) {
			newSlice = reflect.Append(newSlice, aVal.Index(i))
		}
	}

	return newSlice.Interface()
}
