package villa

import (
	"testing"
)

func TestStrSet(t *testing.T) {
	defer __(o_(t))

	var set StrSet

	set.Put("hello", "david")
	t.Logf("%v", set)
	AssertEquals(t, "set.In(hello)", set.In("hello"), true)
	AssertEquals(t, "set.In(david)", set.In("david"), true)
	AssertEquals(t, "set.In(villa)", set.In("villa"), false)

	AssertEquals(t, "set.Equals(david, hello)", set.Equals(NewStrSet("david", "hello")), true)
	AssertEquals(t, "set.Equals(david, hello)", set.Equals(NewStrSet("hello")), false)

	set.Delete("david")
	AssertEquals(t, "set.In(david)", set.In("david"), false)
	AssertEquals(t, "set.Equals(hello)", set.Equals(NewStrSet("hello")), true)

	AssertEquals(t, "set.Elements().Equals(hello)", set.Elements().Equals(StringSlice{"hello"}), true)
}

func TestStrSet_nil(t *testing.T) {
	var s, ss StrSet
	AssertEquals(t, "nil.In(david)", s.In("david"), false)
	AssertEquals(t, "nil.Equals(nil)", s.Equals(ss), true)
	AssertStringEquals(t, "nil.Elements()", s.Elements(), StringSlice{})
	s.Delete("david")
}
