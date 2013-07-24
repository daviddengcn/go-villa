package villa

import (
	"errors"
	"testing"
)

func TestNestedError(t *testing.T) {
	defer __(o_(t))

	orgErr := errors.New("Original Error")
	ne := NestErrorf(NestErrorf(orgErr, "Level %d", 1), "Nested again")
	AssertEquals(t, "ne.Error", ne.Error(), "Nested again: Level 1: Original Error")
	AssertEquals(t, "ne.Deepest", ne.Deepest(), orgErr)
}

/*
	AssertEquals shows error message when act and exp don't equal
*/
func AssertEquals(t *testing.T, name string, act, exp interface{}) {
	if act != exp {
		t.Errorf("%s is expected to be %v, but got %v", name, exp, act)
	}
}

/*
	AssertEquals shows error message when strings forms of act and exp don't
	equal
*/
func AssertStringEquals(t *testing.T, name string, act, exp interface{}) {
	if fmt.Sprintf("%v", act) != fmt.Sprintf("%v", exp) {
		t.Errorf("%s is expected to be %v, but got %v", name, exp, act)
	} // if
}

/*
	AssertStrSetEquals shows error message when act and exp are equal string
	sets.
*/
func AssertStrSetEquals(t *testing.T, name string, act, exp StrSet) {
	if !act.Equals(exp) {
		t.Errorf("%s is expected to be %v, but got %v", name, exp, act)
	}
}
