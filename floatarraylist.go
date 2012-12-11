package villa

import(
    "fmt"
)

// FloatArrayList is a list of float64 values.
//
// Using FloatArrayList, the sort algorithm can be easily performed by calling the NewLessAdapter method, which
// returns a new adapter struct that implements an extra Less() method and thus satisfied sort.Interface.
//    lst := villa.FloatArrayList()
//    sort.Sort(lst.NewLessAdapter(
//        func (a, b float64) bool {
//            return a < b
//        }))
type FloatArrayList struct {
    data []float64
}

// NewFloatArrayList creates a new FloatArrayList instance
func NewFloatArrayList() *FloatArrayList {
    return &FloatArrayList{}
}

// Add appends the specified element to the end of this list.
func (lst *FloatArrayList) Add(e float64) {
    lst.data = append(lst.data, e)
}

// AddSlice appends a slice of elements to the end of this list.
func (lst *FloatArrayList) AddSlice(s []float64) {
    lst.data = append(lst.data, s...)
}

// AddAll appends all elements of another FloatArrayList to the end of this list.
func (lst *FloatArrayList) AddAll(al *FloatArrayList) {
    lst.data = append(lst.data, al.data...)
}

// Get returns the element at the specified position in this list.
func (lst *FloatArrayList) Get(index int) float64 {
    return lst.data[index]
}
// Set replaces the element at the specified position in this list with the specified element.
func (lst *FloatArrayList) Set(index int, e float64) {
    lst.data[index] = e
}
// Data returns the internal float64 slice. If the array list performs structural modification, the returned
// slice could be no longer the one inside the array list.
func (lst *FloatArrayList) Data() []float64 {
    return lst.data
}

// The Swap method in sort.Interface.
func (lst *FloatArrayList) Swap(i, j int) {
    lst.data[i], lst.data[j] = lst.data[j], lst.data[i]
}

// Insert inserts the specified element at the specified position in this list.
func (lst *FloatArrayList) Insert(index int, e float64) {
    lst.data = append(lst.data, 0)
    copy(lst.data[index + 1:], lst.data[index:])
    lst.data[index] = e
}

// The Push method
func (lst *FloatArrayList) Push(e float64) {
    lst.data = append(lst.data, e)
}

// The Pop method
func (lst *FloatArrayList) Pop() (e float64) {
    e = lst.data[len(lst.data) - 1]
    lst.data = lst.data[0:len(lst.data) - 1]
    return
}

// Remove removes the element at the specified position in this list.
func (lst *FloatArrayList) Remove(index int) {
    lst.data = append(lst.data[0:index], lst.data[index + 1:]...)
}

// RemoveRange removes from this list all of the elements whose index is between from, inclusive, and to, exclusive.
func (lst *FloatArrayList) RemoveRange(from, to int) {
    lst.data = append(lst.data[0:from], lst.data[to:]...)
}

// Clear removes all of the elements from this list.
func (lst *FloatArrayList) Clear() {
    lst.data = lst.data[:0]
}

// Len returns the number of elements in this list.
//
// The Len method in sort.Interface.
func (lst *FloatArrayList) Len() int {
    return len(lst.data)
}

// FloatLessFunc is the function compares two elements.
type FloatLessFunc func(float64, float64) bool

// LessAdapter is an adapter struct for an ArrayList with a less function.
type FloatLessAdapter struct {
    *FloatArrayList
    less FloatLessFunc
}

// The Less method in sort.Interface
func (adp *FloatLessAdapter) Less(i, j int) bool {
    return adp.less(adp.data[i], adp.data[j])
}

// NewLessAdapter returns an adapter instance that implenents sort.Interface.Less function.
func (lst *FloatArrayList) NewLessAdapter(less FloatLessFunc) *FloatLessAdapter {
    return &FloatLessAdapter{lst, less}
}

// String returns the internal data's string format as a result
func (lst *FloatArrayList) String() string {
    return fmt.Sprintf("%v", lst.data)
}
  