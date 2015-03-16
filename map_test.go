package villa

import (
	"testing"
)

func TestMapKeyAsStrings(t *testing.T) {
	m := map[string]int{
		"abc": 1,
		"def": 2,
	}
	AssertStrSetEquals(t, "keys", NewStrSet(MapKeyAsStrings(m)...), NewStrSet("abc", "def"))
}
