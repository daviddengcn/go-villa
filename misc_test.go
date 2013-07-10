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
