package villa

import "fmt"

// ComplexSlice is wrapper to a slice of complex128.
//
// Using ComplexSlice, the sort/heap algorithm can be easily performed by calling the NewSortList method, which
// returns a new adapter type instance that implements some sort.Interface.
//    var s villa.ComplexSlice
//    s.Add(...)
//    sl := s.NewSortList(
//        func (a, b complex128) int {
//            if a < b {
//                return -1
//            } else if a < b {
//                return 1
//            } // else if
//            return 0
//        })
//    sort.Sort(sl) // sl(and s) is sorted.
//    p, found := sl.BinarySearch(el)
type ComplexSlice []complex128

// Add appends the specified element to the end of this slice.
func (s *ComplexSlice) Add(e... complex128) {
    *s = append(*s, e...)
}

// Insert inserts the specified element at the specified position in this slice.
func (s *ComplexSlice) Insert(index int, e... complex128) {
    *s = append(*s, e...)
    copy((*s)[index + len(e):], (*s)[index:])
    copy((*s)[index:], e[:])
}

// The Swap method in sort.Interface.
func (s *ComplexSlice) Swap(i, j int) {
    (*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

// Remove removes the element at the specified position in this slice.
func (s *ComplexSlice) Remove(index int) complex128 {
    e := (*s)[index]
    *s = append((*s)[0:index], (*s)[index + 1:]...)
    return e
}

// RemoveRange removes all of the elements whose index is between from, inclusive, and to, exclusive, from this slice.
func (s *ComplexSlice) RemoveRange(from, to int) {
    *s = append((*s)[0:from], (*s)[to:]...)
}

// Clear removes all of the elements from this slice.
func (s *ComplexSlice) Clear() {
    *s = (*s)[:0]
}

// String returns the internal data's string format as a result
func (s *ComplexSlice) String() string {
    return fmt.Sprintf("%v", *s)
}

// ComplexSortList is an adapter struct for an ComplexSlice which implements the sort interface and related functions using a comparator.
type ComplexSortList struct {
    *ComplexSlice
    cmp ComplexCmpFunc
}

// The Len method in sort.Interface.
func (s *ComplexSortList) Len() int {
    return len(*s.ComplexSlice)
}

// The Less method in sort.Interface
func (s *ComplexSortList) Less(i, j int) bool {
    return s.cmp((*s.ComplexSlice)[i], (*s.ComplexSlice)[j]) < 0
}

// BinarySearch searchs a specified element e in a *sorted* list with binary search algorithm. If the list values are not sorted, the return values are undefined.
// If the element is found in the list, found equals true and pos is the index of the found element in the list.
// Otherwise found returns false and pos is the position where e is going to be inserted(and the resulting list is still in order)
func (s *ComplexSortList) BinarySearch(e complex128) (pos int, found bool) {
    l, r := 0, len(*s.ComplexSlice) - 1
    for l <= r {
        m := (l + r) / 2
        c := s.cmp(e, (*s.ComplexSlice)[m])
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

// NewSortList returns an adapter type instance that implenents sort.Interface. A compare function (ComplexCmpFunc) is needed to define the order of elements.
func (s *ComplexSlice) NewSortList(cmp ComplexCmpFunc) *ComplexSortList {
    return &ComplexSortList{s, cmp}
}
