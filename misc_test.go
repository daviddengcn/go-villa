package villa

import (
	"errors"
	"testing"

	"github.com/golangplus/testing/assert"
)

func TestNestedError(t *testing.T) {
	defer __(o_(t))

	orgErr := errors.New("Original Error")
	ne := NestErrorf(NestErrorf(orgErr, "Level %d", 1), "Nested again")
	assert.Equal(t, "ne.Error", ne.Error(), "Nested again: Level 1: Original Error")
	assert.Equal(t, "ne.Deepest", ne.Deepest(), orgErr)
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
