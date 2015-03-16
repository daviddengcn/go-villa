package villa

import (
	"reflect"
)

// MapKeyAsStrings returns the keys of a map as a string slice.
// if m is not a map, it panics.
// Keys are extracted by reflection which doesn't guarentee the efficiency.
// So this function is mainly used for debugging or error message showing.
func MapKeyAsStrings(m interface{}) []string {
	keys := reflect.ValueOf(m).MapKeys()
	res := make([]string, len(keys))
	for i, key := range keys {
		res[i] = key.String()
	}
	return res
}
