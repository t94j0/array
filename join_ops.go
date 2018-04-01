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
	aVal, aType, bVal, bType := getABValType(a, b)

	itemType := bType.Elem()
	if aType.Kind() != reflect.Slice || bType.Kind() != reflect.Slice {
		return reflect.Zero(itemType)
	}

	newSlice := reflect.MakeSlice(aType, 0, 0)

	aLen := aVal.Len()
	bLen := bVal.Len()
	itemMap := reflect.MapOf(itemType, reflect.TypeOf(1))
	newMap := reflect.MakeMapWithSize(itemMap, bLen)

	// Create map with 'b'
	for i := 0; i < bLen; i++ {
		sliceItem := bVal.Index(i)
		newMap.SetMapIndex(sliceItem, reflect.ValueOf(1))
	}

	// Check 'b' map with 'a' values
	for i := 0; i < aLen; i++ {
		sliceItem := aVal.Index(i)

		// TODO: There should be a better way to do this
		if newMap.MapIndex(sliceItem).Kind().String() != reflect.Invalid.String() {
			newSlice = reflect.Append(newSlice, sliceItem)
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
