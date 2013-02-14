package villa

import(
	"sort"
)

type sortI struct {
	l int
	less func(int, int) bool
	swap func(int, int)
}

func (s *sortI) Len() int {
	return s.l
}

func (s *sortI) Less(i, j int) bool {
	return s.less(i, j)
}

func (s *sortI) Swap(i, j int) {
	s.swap(i, j)
}

// SortF sorts the data defined by lenght, Less/Swap functions.
func SortF(Len int, Less func(int, int) bool, Swap func(int, int)) {
	sort.Sort(&sortI{l: Len, less: Less, swap: Swap})
}