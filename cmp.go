package villa

import "sort"

// CmpFunc is a function comparing two elements. The function returns a positive value if a > b, a negative value if a < b, and 0 otherwise.
//
// Sort, BinarySearch and Merge methods are defined.
// Usage:
//    s := []interface{}{...}
//    cmp := CmpFunc(func (a, b interface{}) int {
//            if a.(int) < b.(int) {
//                return -1
//            } else if a.(int) > b.(int) {
//                return 1
//            } // else if
//            return 0
//    })
//    cmp.Sort(s)
//    p, found := cmp.BinarySearch(s, e)
//
//    l := []interface{}{...}
//    cmp.Sort(l)
//    t := cmp.Merge(s, l)
type CmpFunc func(a, b interface{}) int

// Merge merges the current *sorted* elements with another *sorted* slice of elements.
// All elements should be sorted by the same comparator.
func (cmp CmpFunc) Merge(a, b []interface{}) []interface{} {
	na, nb := len(a), len(b)
	res := make([]interface{}, na+nb)
	for k, l, m := 0, 0, 0; l < na || m < nb; k++ {
		if m >= nb || l < na && cmp(a[l], b[m]) <= 0 {
			res[k] = a[l]
			l++
		} else {
			res[k] = b[m]
			m++
		} // else
	} // for l, m, k

	return res
}

// BinarySearch searchs a specified element e in a *sorted* list with binary search algorithm. If the list values are not sorted, the return values are undefined.
// If the element is found in the list, found equals true and pos is the index of the found element in the list.
// Otherwise found returns false and pos is the position where e is going to be inserted(and the resulting list is still in order)
func (cmp CmpFunc) BinarySearch(s []interface{}, e interface{}) (pos int, found bool) {
	l, r := 0, len(s)-1
	for l <= r {
		m := l + (r-l)/2
		c := cmp(e, s[m])
		if c == 0 {
			return m, true
		} // if

		if c < 0 {
			r = m - 1
		} else {
			l = m + 1
		} // else
	} // for
	return l, false
}

type sortList struct {
	Slice
	cmp CmpFunc
}

// The Len method in sort.Interface.
func (s *sortList) Len() int {
	return len(s.Slice)
}

// The Less method in sort.Interface
func (s *sortList) Less(i, j int) bool {
	return s.cmp(s.Slice[i], s.Slice[j]) < 0
}

// Sort calls the build-in sort.Sort to sort data in the slice.
func (cmp CmpFunc) Sort(s []interface{}) {
	sort.Sort(&sortList{Slice(s), cmp})
}

// IntCmpFunc is a function comparing two int elements. The function returns a positive value if a > b, a negative value if a < b, and 0 otherwise.
//
// Sort, BinarySearch and Merge methods are defined.
// Usage:
//    s := []int{}{...}
//    cmp := CmpFunc(func (a, b int) int {
//        if a < b {
//            return -1
//        } else if a > b {
//            return 1
//        } // else if
//        return 0
//    })
//    cmp.Sort(s)
//    p, found := cmp.BinarySearch(s, e)
//
//    l := []int{}{...}
//    cmp.Sort(l)
//    t := cmp.Merge(s, l)
type IntCmpFunc func(a, b int) int

// Merge merges the current *sorted* elements with another *sorted* slice of elements.
// All elements should be sorted by the same comparator.
func (cmp IntCmpFunc) Merge(a, b []int) []int {
	na, nb := len(a), len(b)
	res := make([]int, na+nb)
	for k, l, m := 0, 0, 0; l < na || m < nb; k++ {
		if m >= nb || l < na && cmp(a[l], b[m]) <= 0 {
			res[k] = a[l]
			l++
		} else {
			res[k] = b[m]
			m++
		} // else
	} // for l, m, k

	return res
}

// BinarySearch searchs a specified element e in a *sorted* list with binary search algorithm. If the list values are not sorted, the return values are undefined.
// If the element is found in the list, found equals true and pos is the index of the found element in the list.
// Otherwise found returns false and pos is the position where e is going to be inserted(and the resulting list is still in order)
func (cmp IntCmpFunc) BinarySearch(s []int, e int) (pos int, found bool) {
	l, r := 0, len(s)-1
	for l <= r {
		m := l + (r-l)/2
		c := cmp(e, s[m])
		if c == 0 {
			return m, true
		} // if

		if c < 0 {
			r = m - 1
		} else {
			l = m + 1
		} // else
	} // for
	return l, false
}

type intSortList struct {
	IntSlice
	cmp IntCmpFunc
}

// The Len method in sort.Interface.
func (s *intSortList) Len() int {
	return len(s.IntSlice)
}

// The Less method in sort.Interface
func (s *intSortList) Less(i, j int) bool {
	return s.cmp(s.IntSlice[i], s.IntSlice[j]) < 0
}

// Sort calls the build-in sort.Sort to sort data in the slice.
func (cmp IntCmpFunc) Sort(s []int) {
	sort.Sort(&intSortList{IntSlice(s), cmp})
}

// FloatCmpFunc is a function comparing two float elements. The function returns a positive value if a > b, a negative value if a < b, and 0 otherwise.
//
// Sort, BinarySearch and Merge methods are defined.
// Usage:
//    s := []float64{}{...}
//    cmp := CmpFunc(func (a, b float64) int {
//        if a < b {
//            return -1
//        } else if a > b {
//            return 1
//        } // else if
//        return 0
//    })
//    cmp.Sort(s)
//    p, found := cmp.BinarySearch(s, e)
//
//    l := []float64{}{...}
//    cmp.Sort(l)
//    t := cmp.Merge(s, l)
type FloatCmpFunc func(a, b float64) int

// Merge merges the current *sorted* elements with another *sorted* slice of elements.
// All elements should be sorted by the same comparator.
func (cmp FloatCmpFunc) Merge(a, b []float64) []float64 {
	na, nb := len(a), len(b)
	res := make([]float64, na+nb)
	for k, l, m := 0, 0, 0; l < na || m < nb; k++ {
		if m >= nb || l < na && cmp(a[l], b[m]) <= 0 {
			res[k] = a[l]
			l++
		} else {
			res[k] = b[m]
			m++
		} // else
	} // for l, m, k

	return res
}

// BinarySearch searchs a specified element e in a *sorted* list with binary search algorithm. If the list values are not sorted, the return values are undefined.
// If the element is found in the list, found equals true and pos is the index of the found element in the list.
// Otherwise found returns false and pos is the position where e is going to be inserted(and the resulting list is still in order)
func (cmp FloatCmpFunc) BinarySearch(s []float64, e float64) (pos int, found bool) {
	l, r := 0, len(s)-1
	for l <= r {
		m := l + (r-l)/2
		c := cmp(e, s[m])
		if c == 0 {
			return m, true
		} // if

		if c < 0 {
			r = m - 1
		} else {
			l = m + 1
		} // else
	} // for
	return l, false
}

type floatSortList struct {
	FloatSlice
	cmp FloatCmpFunc
}

// The Len method in sort.Interface.
func (s *floatSortList) Len() int {
	return len(s.FloatSlice)
}

// The Less method in sort.Interface
func (s *floatSortList) Less(i, j int) bool {
	return s.cmp(s.FloatSlice[i], s.FloatSlice[j]) < 0
}

// Sort calls the build-in sort.Sort to sort data in the slice.
func (cmp FloatCmpFunc) Sort(s []float64) {
	sort.Sort(&floatSortList{FloatSlice(s), cmp})
}

// ComplexCmpFunc is a function comparing two complex128 elements. The function returns a positive value if a > b, a negative value if a < b, and 0 otherwise.
//
// Sort, BinarySearch and Merge methods are defined.
// Usage:
//    s := []complex128{}{...}
//    cmp := CmpFunc(func (a, b complex128) int {
//        if cmplx.Abs(a) < cmplx.Abs(b) {
//            return -1
//        } else if cmplx.Abs(a) > cmplx.Abs(b) {
//            return 1
//        } // else if
//        return 0
//    })
//    cmp.Sort(s)
//    p, found := cmp.BinarySearch(s, e)
//
//    l := []complex128{}{...}
//    cmp.Sort(l)
//    t := cmp.Merge(s, l)
type ComplexCmpFunc func(a, b complex128) int

func (cmp ComplexCmpFunc) Merge(a, b []complex128) []complex128 {
	na, nb := len(a), len(b)
	res := make([]complex128, na+nb)
	for k, l, m := 0, 0, 0; l < na || m < nb; k++ {
		if m >= nb || l < na && cmp(a[l], b[m]) <= 0 {
			res[k] = a[l]
			l++
		} else {
			res[k] = b[m]
			m++
		} // else
	} // for l, m, k

	return res
}

// BinarySearch searchs a specified element e in a *sorted* list with binary search algorithm. If the list values are not sorted, the return values are undefined.
// If the element is found in the list, found equals true and pos is the index of the found element in the list.
// Otherwise found returns false and pos is the position where e is going to be inserted(and the resulting list is still in order)
func (cmp ComplexCmpFunc) BinarySearch(s []complex128, e complex128) (pos int, found bool) {
	l, r := 0, len(s)-1
	for l <= r {
		m := l + (r-l)/2
		c := cmp(e, s[m])
		if c == 0 {
			return m, true
		} // if

		if c < 0 {
			r = m - 1
		} else {
			l = m + 1
		} // else
	} // for
	return l, false
}

type complexSortList struct {
	ComplexSlice
	cmp ComplexCmpFunc
}

// The Len method in sort.Interface.
func (s *complexSortList) Len() int {
	return len(s.ComplexSlice)
}

// The Less method in sort.Interface
func (s *complexSortList) Less(i, j int) bool {
	return s.cmp(s.ComplexSlice[i], s.ComplexSlice[j]) < 0
}

// Sort calls the build-in sort.Sort to sort data in the slice.
func (cmp ComplexCmpFunc) Sort(s []complex128) {
	sort.Sort(&complexSortList{ComplexSlice(s), cmp})
}

var (
	// IntValueCompare compares the input int values a and b, returns -1 if a < b, 1 if a > b, and 0 otherwise.
	// This is a natural IntCmpFunc.
	IntValueCompare = IntCmpFunc(func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		} // else if

		return 0
	})

	// FloatValueCompare compares the input float64 values a and b, returns -1 if a < b, 1 if a > b, and 0 otherwise.
	// This is a natural FloatCmpFunc.
	FloatValueCompare = FloatCmpFunc(func(a, b float64) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		} // else if

		return 0
	})
)
