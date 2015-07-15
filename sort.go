package villa

import (
	"github.com/golangplus/sort"
)

// Deprecated. Use sortp.SortF in "github.com/golangplus/sort" instead.
func SortF(Len int, Less func(int, int) bool, Swap func(int, int)) {
	sortp.SortF(Len, Less, Swap)
}
