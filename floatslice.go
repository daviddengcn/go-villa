package villa

import "fmt"

// FloatSlice is wrapper to a slice of float64.
//
// Using FloatSlice, the sort/heap algorithm can be easily performed by calling the NewSortList method, which
// returns a new adapter type instance that implements some sort.Interface.
//    var s villa.FloatSlice
//    s.Add(...)
//    sl := s.NewSortList(
//        func (a, b float64) int {
//            if a < b {
//                return -1
//            } else if a < b {
//                return 1
//            } // else if
//            return 0
//        })
//    sort.Sort(sl) // sl(and s) is sorted.
//    p, found := sl.BinarySearch(el)
type FloatSlice []float64

// Add appends the specified element to the end of this slice.
func (s *FloatSlice) Add(e... float64) {
    *s = append(*s, e...)
}

// Insert inserts the specified element at the specified position in this slice.
// NOTE: the insertion algorithm is much better than the slice-trick in go community wiki
func (s *FloatSlice) Insert(index int, e... float64) {
    if cap(*s) >= len(*s) + len(e) {
        *s = (*s)[:len(*s) + len(e)]
    } else {
        *s = append(*s, e...)
    } // else
    copy((*s)[index + len(e):], (*s)[index:])
    copy((*s)[index:], e[:])
}

// The Swap method in sort.Interface.
func (s *FloatSlice) Swap(i, j int) {
    (*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

// Remove removes the element at the specified position in this slice.
func (s *FloatSlice) Remove(index int) float64 {
    e := (*s)[index]
    *s = append((*s)[0:index], (*s)[index + 1:]...)
    return e
}

// RemoveRange removes all of the elements whose index is between from, inclusive, and to, exclusive, from this slice.
func (s *FloatSlice) RemoveRange(from, to int) {
    *s = append((*s)[0:from], (*s)[to:]...)
}

// Fill sets the elements between from, inclusive, and to, exclusive, to a speicified value.
func (s *FloatSlice) Fill(from, to int, vl float64) {
    for i := from; i < to; i ++ {
        (*s)[i] = vl
    } // for i
}

// Clear removes all of the elements from this slice.
func (s *FloatSlice) Clear() {
    *s = (*s)[:0]
}

// String returns the internal data's string format as a result
func (s *FloatSlice) String() string {
    return fmt.Sprintf("%v", *s)
}

// FloatSortList is an adapter struct for an FloatSlice which implements the sort interface and related functions using a comparator.
type FloatSortList struct {
    *FloatSlice
    cmp FloatCmpFunc
}

// The Len method in sort.Interface.
func (s *FloatSortList) Len() int {
    return len(*s.FloatSlice)
}

// The Less method in sort.Interface
func (s *FloatSortList) Less(i, j int) bool {
    return s.cmp((*s.FloatSlice)[i], (*s.FloatSlice)[j]) < 0
}

// BinarySearch searchs a specified element e in a *sorted* list with binary search algorithm. If the list values are not sorted, the return values are undefined.
// If the element is found in the list, found equals true and pos is the index of the found element in the list.
// Otherwise found returns false and pos is the position where e is going to be inserted(and the resulting list is still in order)
func (s *FloatSortList) BinarySearch(e float64) (pos int, found bool) {
    l, r := 0, len(*s.FloatSlice) - 1
    for l <= r {
        m := l +(r - l)/2
        c := s.cmp(e, (*s.FloatSlice)[m])
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

// NewSortList returns an adapter type instance that implenents sort.Interface. A compare function (FloatCmpFunc) is needed to define the order of elements.
func (s *FloatSlice) NewSortList(cmp FloatCmpFunc) *FloatSortList {
    return &FloatSortList{s, cmp}
}
