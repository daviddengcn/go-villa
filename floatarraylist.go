package villa

import(
    "fmt"
)

// FloatArrayList is a list of float64 values.
//
// Using FloatArrayList, the sort algorithm can be easily performed by calling the NewCmpAdapter method, which
// returns a new adapter struct that implements an extra Less() method and thus satisfied sort.Interface.
//    lst := villa.FloatArrayList()
//    adp := lst.NewCmpAdapter(
//        func (a, b float64) bool {
//            if a < b {
//                return -1
//            } else if a < b {
//                return 1
//            } // else if
//            return 0
//        })
//    sort.Sort(adp)
//    p, found := adp.BinarySearch(el)
type FloatArrayList struct {
    data []float64
}

// NewFloatArrayList creates a new FloatArrayList instance
func NewFloatArrayList() *FloatArrayList {
    return &FloatArrayList{}
}

// NewFloatArrayListCap creates a new FloatArrayList instance with an initialized length and capacity
func NewFloatArrayListCap(len, cap int) *FloatArrayList {
    return &FloatArrayList{data: make([]float64, len, cap)}
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

// FloatCmpAdapter is an adapter struct for an ArrayList with a cmp function.
type FloatCmpAdapter struct {
    *FloatArrayList
    cmp FloatCmpFunc
}

// The Less method in sort.Interface
func (adp *FloatCmpAdapter) Less(i, j int) bool {
    return adp.cmp(adp.data[i], adp.data[j]) < 0
}

// BinarySearch searchs a specified element e in a *sorted* list with binary search algorithm. If the list values are not sorted, the return values are undefined.
// If the element is found in the list, found equals true and pos is the index of the found element in the list.
// Otherwise found returns false and pos is the position where e is going to be inserted(and the resulting list is still in order)
func (adp *FloatCmpAdapter) BinarySearch(e float64) (pos int, found bool) {
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
func (lst *FloatArrayList) NewCmpAdapter(cmp FloatCmpFunc) *FloatCmpAdapter {
    return &FloatCmpAdapter{lst, cmp}
}

// String returns the internal data's string format as a result
func (lst *FloatArrayList) String() string {
    return fmt.Sprintf("%v", lst.data)
}
  