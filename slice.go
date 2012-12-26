package villa

import "fmt"

// Slice is wrapper to a slice of interface{}.
//
// Using Slice, the sort/heap algorithm can be easily performed by calling the NewSortList method, which
// returns a new adapter type instance that implements some sort.Interface and heap.Interface.
//    var s villa.Slice
//    s.Add(...)
//    sl := s.NewSortList(
//        func (a, b interface{}) int {
//            if a.(int) < b.(int) {
//                return -1
//            } else if a.(int) < b.(int) {
//                return 1
//            } // else if
//            return 0
//        })
//    sort.Sort(sl) // sl(and s) is sorted.
//    p, found := sl.BinarySearch(el)
type Slice []interface{}

// Add appends the specified element to the end of this slice.
func (s *Slice) Add(e... interface{}) {
    *s = append(*s, e...)
}

// Insert inserts the specified element at the specified position in this slice.
// NOTE: the insertion algorithm is much better than the slice-trick in go community wiki
func (s *Slice) Insert(index int, e... interface{}) {
    if cap(*s) >= len(*s) + len(e) {
        *s = (*s)[:len(*s) + len(e)]
    } else {
        *s = append(*s, e...)
    } // else
    copy((*s)[index + len(e):], (*s)[index:])
    copy((*s)[index:], e[:])
}

// The Swap method in sort.Interface.
func (s Slice) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

// Remove removes the element at the specified position in this slice.
// The hole in the original slice is filled with a nil value.
func (s *Slice) Remove(index int) interface{} {
    e := (*s)[index]
    copy((*s)[index:], (*s)[index + 1:])
    (*s)[len(*s) - 1] = nil
    *s = (*s)[:len(*s) - 1]
    return e
}

// RemoveRange removes all of the elements whose index is between from, inclusive, and to, exclusive, from this slice.
// The holes in the original slice are filled with nil values.
func (s *Slice) RemoveRange(from, to int) {
    if to <= from {
        return
    } // if
    
    copy((*s)[from:], (*s)[to:])
    n := len(*s);  l := n - to + from
    for i := l; i < n; i ++ {
        (*s)[i] = nil
    } // for i
    *s = (*s)[:l]
}

// Fill sets the elements between from, inclusive, and to, exclusive, to a speicified value.
func (s Slice) Fill(from, to int, vl interface{}) {
    for i := from; i < to; i ++ {
        s[i] = vl
    } // for i
}

// Clear sets the slice to nil.
func (s *Slice) Clear() {
    *s = (*s)[:0]
}

// String returns the internal data's string format as a result
func (s *Slice) String() string {
    return fmt.Sprintf("%v", *s)
}

// SortList is an adapter struct for a Slice which implements the sort interface and related functions with a comparator.
type SortList struct {
    *Slice
    cmp CmpFunc
}

// The Push method in heap.Interface.
func (s *SortList) Push(e interface{}) {
    *s.Slice = append(*s.Slice, e)
}

// The Pop method in heap.Interface.
func (s *SortList) Pop() interface{} {
    return s.Remove(len(*s.Slice) - 1)
}

// The Len method in sort.Interface.
func (s *SortList) Len() int {
    return len(*s.Slice)
}

// The Less method in sort.Interface
func (s *SortList) Less(i, j int) bool {
    return s.cmp((*s.Slice)[i], (*s.Slice)[j]) < 0
}

// Get returns the i-th element in the slice. This is implemented since [] operator is not embedded
func (s *SortList) Get(i int) interface{} {
    return (*s.Slice)[i]
}

// BinarySearch searchs a specified element e in a *sorted* list with binary search algorithm. If the list values are not sorted, the return values are undefined.
// If the element is found in the list, found equals true and pos is the index of the found element in the list.
// Otherwise found returns false and pos is the position where e is going to be inserted(and the resulting list is still in order)
func (s *SortList) BinarySearch(e interface{}) (pos int, found bool) {
    l, r := 0, len(*s.Slice) - 1
    for l <= r {
        m := l +(r - l)/2
        c := s.cmp(e, (*s.Slice)[m])
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

// NewSortList returns an adapter type instance that implenents sort.Interface and heap.Interface. A compare function (CmpFunc) is needed to define the order of elements.
func (s *Slice) NewSortList(cmp CmpFunc) *SortList {
    return &SortList{s, cmp}
}
