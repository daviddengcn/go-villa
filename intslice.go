package villa

import "fmt"

// IntSlice is wrapper to a slice of int.
//
// Using IntSlice, the sort/heap algorithm can be easily performed by calling the NewSortList method, which
// returns a new adapter type instance that implements some sort.Interface.
//    var s villa.IntSlice
//    s.Add(...)
//    sl := s.NewSortList(
//        func (a, b int) int {
//            if a < b {
//                return -1
//            } else if a < b {
//                return 1
//            } // else if
//            return 0
//        })
//    sort.Sort(sl) // sl(and s) is sorted.
//    p, found := sl.BinarySearch(el)
type IntSlice []int

// Add appends the specified element to the end of this slice.
func (s *IntSlice) Add(e... int) {
    *s = append(*s, e...)
}

// Insert inserts the specified element at the specified position in this slice.
// NOTE: the insertion algorithm is much better than the slice-trick in go community wiki
func (s *IntSlice) Insert(index int, e... int) {
    if cap(*s) >= len(*s) + len(e) {
        *s = (*s)[:len(*s) + len(e)]
    } else {
        *s = append(*s, e...)
    } // else
    copy((*s)[index + len(e):], (*s)[index:])
    copy((*s)[index:], e[:])
}

// The Swap method in sort.Interface.
func (s *IntSlice) Swap(i, j int) {
    (*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

// Remove removes the element at the specified position in this slice.
func (s *IntSlice) Remove(index int) int {
    e := (*s)[index]
    *s = append((*s)[0:index], (*s)[index + 1:]...)
    return e
}

// RemoveRange removes all of the elements whose index is between from, inclusive, and to, exclusive, from this slice.
func (s *IntSlice) RemoveRange(from, to int) {
    *s = append((*s)[0:from], (*s)[to:]...)
}

// Fill sets the elements between from, inclusive, and to, exclusive, to a speicified value.
func (s *IntSlice) Fill(from, to int, vl int) {
    for i := from; i < to; i ++ {
        (*s)[i] = vl
    } // for i
}

// Clear removes all of the elements from this slice.
func (s *IntSlice) Clear() {
    *s = (*s)[:0]
}

// String returns the internal data's string format as a result
func (s *IntSlice) String() string {
    return fmt.Sprintf("%v", *s)
}

// IntSortList is an adapter struct for an IntSlice which implements the sort interface and related functions using a comparator.
type IntSortList struct {
    *IntSlice
    cmp IntCmpFunc
}

// The Len method in sort.Interface.
func (s *IntSortList) Len() int {
    return len(*s.IntSlice)
}

// The Less method in sort.Interface
func (s *IntSortList) Less(i, j int) bool {
    return s.cmp((*s.IntSlice)[i], (*s.IntSlice)[j]) < 0
}

// BinarySearch searchs a specified element e in a *sorted* list with binary search algorithm. If the list values are not sorted, the return values are undefined.
// If the element is found in the list, found equals true and pos is the index of the found element in the list.
// Otherwise found returns false and pos is the position where e is going to be inserted(and the resulting list is still in order)
func (s *IntSortList) BinarySearch(e int) (pos int, found bool) {
    l, r := 0, len(*s.IntSlice) - 1
    for l <= r {
        m := l +(r - l)/2
        c := s.cmp(e, (*s.IntSlice)[m])
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

// NewSortList returns an adapter type instance that implenents sort.Interface. A compare function (IntCmpFunc) is needed to define the order of elements.
func (s *IntSlice) NewSortList(cmp IntCmpFunc) *IntSortList {
    return &IntSortList{s, cmp}
}
