package villa

import "fmt"

// IntSlice is wrapper to a slice of int.
//
// Using IntSlice, the sort/heap algorithm can be easily performed by calling the NewSortList method, which
// returns a new adapter type instance that implements some sort.Interface and heap.Interface.
//    var s villa.IntSlice
//    s.Add(...)
//    s := s.NewSortAdapter(
//        func (a, b int) int {
//            if a < b {
//                return -1
//            } else if a < b {
//                return 1
//            } // else if
//            return 0
//        })
//    sort.Sort(s) // s(and s) is sorted.
//    p, found := s.BinarySearch(el)
type IntSlice []int

// Add appends the specified element to the end of this slice.
func (s *IntSlice) Add(e... int) {
    *s = append(*s, e...)
}

// Insert inserts the specified element at the specified position in this slice.
func (s *IntSlice) Insert(index int, e... int) {
    *s = append(*s, e...)
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

// Get returns the i-th element in the slice. This is implemented since [] operator is not embedded
func (s *IntSortList) Get(i int) int {
    return (*s.IntSlice)[i]
}

// BinarySearch searchs a specified element e in a *sorted* list with binary search algorithm. If the list values are not sorted, the return values are undefined.
// If the element is found in the list, found equals true and pos is the index of the found element in the list.
// Otherwise found returns false and pos is the position where e is going to be inserted(and the resulting list is still in order)
func (s *IntSortList) BinarySearch(e int) (pos int, found bool) {
    l, r := 0, len(*s.IntSlice) - 1
    for l <= r {
        m := (l + r) / 2
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

// NewSortList returns an adapter type instance that implenents sort.Interface and heap.Interface. A compare function (IntCmpFunc) is needed to define the order of elements.
func (s *IntSlice) NewSortList(cmp IntCmpFunc) *IntSortList {
    return &IntSortList{s, cmp}
}
