package villa

import(
	"testing"
)

func TestSort(t *testing.T) {
    defer __(o_())
    
	ints := []int{3, 4, 1, 7, 0}
	SortF(len(ints), func(i, j int) bool {
		return ints[i] < ints[j]
	}, func (i, j int) {
		ints[i], ints[j] = ints[j], ints[i]
	})
	AssertStringEquals(t, "ints", ints, []int{0, 1, 3, 4, 7})

	ints = []int{3, 4, 1, 7, 0}
	SortF(len(ints), func(i, j int) bool {
		return ints[i] > ints[j]
	}, func (i, j int) {
		ints[i], ints[j] = ints[j], ints[i]
	})
	AssertStringEquals(t, "ints", ints, []int{7, 4, 3, 1, 0})
}