package villa

import(
    "fmt"
)

// ArrayList is a list of values(in the form of interface{}).
//
// Using ArrayList, the sort/heap algorithm can be easily performed by calling the NewCmpAdapter method, which
// returns a new adapter struct that implements an extra Less() method and thus satisfied sort.Interface and 
// heap.Interface.
//    lst := villa.NewArrayList()
//    sort.Sort(lst.NewCmpAdapter(
//        func (a, b interface{}) int {
//            if a.(int) < b.(int) {
//                return -1
//            } else if a.(int) < b.(int) {
//                return 1
//            } // else if
//            return 0
//        }))
type ArrayList struct {
    data []interface{}
}

// NewArrayList creates a new ArrayList instance
func NewArrayList() *ArrayList {
    return &ArrayList{}
}
// NewArrayListCap creates a new ArrayList instance with an initialized length and capacity
func NewArrayListCap(len, cap int) *ArrayList {
    return &ArrayList{data: make([]interface{}, len, cap)}
}

// Add appends the specified element to the end of this list.
func (lst *ArrayList) Add(e interface{}) {
    lst.data = append(lst.data, e)
}

// AddSlice appends a slice of elements to the end of this list.
func (lst *ArrayList) AddSlice(s []interface{}) {
    lst.data = append(lst.data, s...)
}

// AddAll appends all elements of another ArrayList to the end of this list.
func (lst *ArrayList) AddAll(al *ArrayList) {
    lst.data = append(lst.data, al.data...)
}

// Get returns the element at the specified position in this list.
func (lst *ArrayList) Get(index int) interface{} {
    return lst.data[index]
}
// Set replaces the element at the specified position in this list with the specified element.
func (lst *ArrayList) Set(index int, e interface{}) {
    lst.data[index] = e
}
// Data returns the internal interface{} slice. If the array list performs structural modification, the returned
// slice could be no longer the one inside the array list.
func (lst *ArrayList) Data() []interface{} {
    return lst.data
}

// The Swap method in sort.Interface.
func (lst *ArrayList) Swap(i, j int) {
    lst.data[i], lst.data[j] = lst.data[j], lst.data[i]
}

// Insert inserts the specified element at the specified position in this list.
func (lst *ArrayList) Insert(index int, e interface{}) {
    lst.data = append(lst.data, nil)
    copy(lst.data[index + 1:], lst.data[index:])
    lst.data[index] = e
}

// The Push method in heap.Interface.
func (lst *ArrayList) Push(e interface{}) {
    lst.data = append(lst.data, e)
}

// The Pop method in heap.Interface.
func (lst *ArrayList) Pop() (e interface{}) {
    e = lst.data[len(lst.data) - 1]
    lst.data = lst.data[0:len(lst.data) - 1]
    return
}

// Remove removes the element at the specified position in this list.
func (lst *ArrayList) Remove(index int) {
    lst.data = append(lst.data[0:index], lst.data[index + 1:]...)
}

// RemoveRange removes from this list all of the elements whose index is between from, inclusive, and to, exclusive.
func (lst *ArrayList) RemoveRange(from, to int) {
    lst.data = append(lst.data[0:from], lst.data[to:]...)
}

// Clear removes all of the elements from this list.
func (lst *ArrayList) Clear() {
    lst.data = lst.data[0:0]
}

// Len returns the number of elements in this list.
//
// The Len method in sort.Interface.
func (lst *ArrayList) Len() int {
    return len(lst.data)
}

// CmpAdapter is an adapter struct for an ArrayList with a less function.
type CmpAdapter struct {
    *ArrayList
    cmp CmpFunc
}

// The Less method in sort.Interface
func (adp *CmpAdapter) Less(i, j int) bool {
    return adp.cmp(adp.data[i], adp.data[j]) < 0
}

// BinarySearch searchs a specified element e in a *sorted* list with binary search algorithm. If the list values are not sorted, the return values are undefined.
// If the element is found in the list, found equals true and pos is the index of the found element in the list.
// Otherwise found returns false and pos is the position where e is going to be inserted(and the resulting list is still in order)
func (adp *CmpAdapter) BinarySearch(e interface{}) (pos int, found bool) {
    l, r := 0, len(adp.data) - 1
    for l <= r {
        m := (l + r) / 2
        c := adp.cmp(e, adp.data[m])
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

// NewCmpAdapter returns an adapter instance that implenents sort.Interface.Less function.
func (lst *ArrayList) NewCmpAdapter(cmp CmpFunc) *CmpAdapter {
    return &CmpAdapter{lst, cmp}
}

// String returns the internal data's string format as a result
func (lst *ArrayList) String() string {
    return fmt.Sprintf("%v", lst.data)
}
