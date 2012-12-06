package villa

import(
    "fmt"
)

// IntArrayList is a list of int values.
//
// Using IntArrayList, the sort algorithm can be easily performed by calling the NewLessAdapter method, which
// returns a new adapter struct that implements an extra Less() method and thus satisfied sort.Interface.
//    lst := villa.IntArrayList()
//    sort.Sort(lst.NewLessAdapter(
//        func (a, b int) bool {
//            return a < b
//        }))
type IntArrayList struct {
    data []int
}

// NewIntArrayList creates a new IntArrayList instance
func NewIntArrayList() *IntArrayList {
    return &IntArrayList{}
}

// Add appends the specified element to the end of this list.
func (lst *IntArrayList) Add(e int) {
    lst.data = append(lst.data, e)
}

// AddSlice appends a slice of elements to the end of this list.
func (lst *IntArrayList) AddSlice(s []int) {
    lst.data = append(lst.data, s...)
}

// AddAll appends all elements of another IntArrayList to the end of this list.
func (lst *IntArrayList) AddAll(al *IntArrayList) {
    lst.data = append(lst.data, al.data...)
}

// Get returns the element at the specified position in this list.
func (lst *IntArrayList) Get(index int) int {
    return lst.data[index]
}
// Set replaces the element at the specified position in this list with the specified element.
func (lst *IntArrayList) Set(index int, e int) {
    lst.data[index] = e
}
// Data returns the internal int slice. If the array list performs structural modification, the returned
// slice could be no longer the one inside the array list.
func (lst *IntArrayList) Data() []int {
    return lst.data
}

// The Swap method in sort.Interface.
func (lst *IntArrayList) Swap(i, j int) {
    lst.data[i], lst.data[j] = lst.data[j], lst.data[i]
}

// Insert inserts the specified element at the specified position in this list.
func (lst *IntArrayList) Insert(index int, e int) {
    lst.data = append(lst.data, 0)
    copy(lst.data[index + 1:], lst.data[index:])
    lst.data[index] = e
}

// The Push method in heap.Interface.
func (lst *IntArrayList) Push(e int) {
    lst.data = append(lst.data, e)
}

// The Pop method in heap.Interface.
func (lst *IntArrayList) Pop() (e int) {
    e = lst.data[len(lst.data) - 1]
    lst.data = lst.data[0:len(lst.data) - 1]
    return
}

// Remove removes the element at the specified position in this list.
func (lst *IntArrayList) Remove(index int) {
    lst.data = append(lst.data[0:index], lst.data[index + 1:]...)
}

// RemoveRange removes from this list all of the elements whose index is between from, inclusive, and to, exclusive.
func (lst *IntArrayList) RemoveRange(from, to int) {
    lst.data = append(lst.data[0:from], lst.data[to:]...)
}

// Clear removes all of the elements from this list.
func (lst *IntArrayList) Clear() {
    lst.data = lst.data[:0]
}

// Len returns the number of elements in this list.
//
// The Len method in sort.Interface.
func (lst *IntArrayList) Len() int {
    return len(lst.data)
}

// LessFunc is the function compares two elements.
type IntLessFunc func(int, int) bool

// LessAdapter is an adapter struct for an ArrayList with a less function.
type IntLessAdapter struct {
    *IntArrayList
    less IntLessFunc
}

// The Less method in sort.Interface
func (adp *IntLessAdapter) Less(i, j int) bool {
    return adp.less(adp.data[i], adp.data[j])
}

// NewLessAdapter returns an adapter instance that implenents sort.Interface.Less function.
func (lst *IntArrayList) NewLessAdapter(less IntLessFunc) *IntLessAdapter {
    return &IntLessAdapter{lst, less}
}

// String returns the internal data's string format as a result
func (lst *IntArrayList) String() string {
    return fmt.Sprintf("%v", lst.data)
}
 