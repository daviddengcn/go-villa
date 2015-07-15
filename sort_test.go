package villa

import (
	"testing"

	"github.com/golangplus/testing/assert"
)

func TestSort(t *testing.T) {
	defer __(o_(t))

	ints := []int{3, 4, 1, 7, 0}
	SortF(len(ints), func(i, j int) bool {
		return ints[i] < ints[j]
	}, func(i, j int) {
		ints[i], ints[j] = ints[j], ints[i]
	})
	assert.StringEqual(t, "ints", ints, []int{0, 1, 3, 4, 7})

	ints = []int{3, 4, 1, 7, 0}
	SortF(len(ints), func(i, j int) bool {
		return ints[i] > ints[j]
	}, func(i, j int) {
		ints[i], ints[j] = ints[j], ints[i]
	})
	assert.StringEqual(t, "ints", ints, []int{7, 4, 3, 1, 0})
}
