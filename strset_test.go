package villa

import (
	"testing"

	"github.com/golangplus/testing/assert"
)

func TestStrSet(t *testing.T) {
	defer __(o_(t))

	var set StrSet

	set.Put("hello", "david")
	t.Logf("%v", set)
	assert.Equal(t, "set.In(hello)", set.In("hello"), true)
	assert.Equal(t, "set.In(david)", set.In("david"), true)
	assert.Equal(t, "set.In(villa)", set.In("villa"), false)

	assert.Equal(t, "set.Equals(david, hello)", set.Equals(NewStrSet("david", "hello")), true)
	assert.Equal(t, "set.Equals(david, hello)", set.Equals(NewStrSet("hello")), false)

	set.Delete("david")
	assert.Equal(t, "set.In(david)", set.In("david"), false)
	assert.Equal(t, "set.Equals(hello)", set.Equals(NewStrSet("hello")), true)

	assert.Equal(t, "set.Elements().Equals(hello)", set.Elements().Equals(StringSlice{"hello"}), true)
}

func TestStrSet_nil(t *testing.T) {
	var s, ss StrSet
	assert.Equal(t, "nil.In(david)", s.In("david"), false)
	assert.Equal(t, "nil.Equals(nil)", s.Equals(ss), true)
	assert.StringEqual(t, "nil.Elements()", s.Elements(), StringSlice{})
	s.Delete("david")
}
