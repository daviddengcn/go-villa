package villa

import "fmt"

// ArrayList is a list of values(in the form of interface{}).
//
// Using ArrayList, the sort/heap algorithm can be easily performed by calling the NewCmpAdapter method, which
// returns a new adapter struct that implements an extra Less() method and thus satisfied sort.Interface and 
// heap.Interface.
//    lst := villa.NewArrayList()
//    sa := lst.NewSortAdapter(
//        func (a, b interface{}) int {
//            if a.(int) < b.(int) {
//                return -1
//            } else if a.(int) < b.(int) {
//                return 1
//            } // else if
//            return 0
//        })
//    sort.Sort(sa) // sa(and lst) is sorted.
//    p, found := sa.BinarySearch(el)
type ArrayList []interface{}

// Add appends the specified element to the end of this list.
func (lst *ArrayList) Add(e... interface{}) {
    *lst = append(*lst, e...)
}

// Insert inserts the specified element at the specified position in this list.
func (lst *ArrayList) Insert(index int, e... interface{}) {
   *lst = append(*lst, e...)
   copy((*lst)[index + len(e):], (*lst)[index:])
   copy((*lst)[index:], e[:])
}

// The Swap method in sort.Interface.
func (lst *ArrayList) Swap(i, j int) {
    (*lst)[i], (*lst)[j] = (*lst)[j], (*lst)[i]
}

// Remove removes the element at the specified position in this list.
func (lst *ArrayList) Remove(index int) interface{} {
    e := (*lst)[index]
    *lst = append((*lst)[0:index], (*lst)[index + 1:]...)
    return e
}

// RemoveRange removes from this list all of the elements whose index is between from, inclusive, and to, exclusive.
func (lst *ArrayList) RemoveRange(from, to int) {
    *lst = append((*lst)[0:from], (*lst)[to:]...)
}

// Clear removes all of the elements from this list.
func (lst *ArrayList) Clear() {
    *lst = (*lst)[:0]
}

// String returns the internal data's string format as a result
func (lst *ArrayList) String() string {
    return fmt.Sprintf("%v", *lst)
}

// SortAdapter is an adapter struct for an ArrayList which implements the sort interface and related functions using a comparator.
type SortAdapter struct {
    *ArrayList
    cmp CmpFunc
}

// The Push method in heap.Interface.
func (sa *SortAdapter) Push(e interface{}) {
    *sa.ArrayList = append(*sa.ArrayList, e)
}

// The Pop method in heap.Interface.
func (sa *SortAdapter) Pop() interface{} {
    return sa.Remove(len(*sa.ArrayList) - 1)
}

// Len returns the number of elements in this list.
//
// The Len method in sort.Interface.
func (sa *SortAdapter) Len() int {
    return len(*sa.ArrayList)
}

// The Less method in sort.Interface
func (sa *SortAdapter) Less(i, j int) bool {
    return sa.cmp((*sa.ArrayList)[i], (*sa.ArrayList)[j]) < 0
}

// Get returns the i-th element in the list. This is implemented since [] operator is not embedded
func (sa *SortAdapter) Get(i int) interface{} {
    return (*sa.ArrayList)[i]
}

// BinarySearch searchs a specified element e in a *sorted* list with binary search algorithm. If the list values are not sorted, the return values are undefined.
// If the element is found in the list, found equals true and pos is the index of the found element in the list.
// Otherwise found returns false and pos is the position where e is going to be inserted(and the resulting list is still in order)
func (sa *SortAdapter) BinarySearch(e interface{}) (pos int, found bool) {
    l, r := 0, len(*sa.ArrayList) - 1
    for l <= r {
        m := (l + r) / 2
        c := sa.cmp(e, (*sa.ArrayList)[m])
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

// NewSortAdapter returns an adapter instance that implenents sort.Interface.Less function.
func (lst *ArrayList) NewSortAdapter(cmp CmpFunc) *SortAdapter {
    return &SortAdapter{lst, cmp}
}
