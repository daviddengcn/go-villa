package villa

import(
    "fmt"
)

// ComplexArrayList is a list of complex128 values.
//
// Using ComplexArrayList, the sort algorithm can be easily performed by calling the NewCmpAdapter method, which
// returns a new adapter struct that implements an extra Less() method and thus satisfied sort.Interface.
//    lst := villa.ComplexArrayList()
//    adp := lst.NewCmpAdapter(
//        func (a, b complex128) bool {
//            if cmplx.Abs(a) < cmplx.Abs(b) {
//                return -1
//            } else if cmplx.Abs(a) > cmplx.Abs(b) {
//                return 1
//            } // else if
//            return 0
//        })
//    sort.Sort(adp)
//    p, found := adp.BinarySearch(el)
type ComplexArrayList struct {
    data []complex128
}

// NewComplexArrayList creates a new ComplexArrayList instance
func NewComplexArrayList() *ComplexArrayList {
    return &ComplexArrayList{}
}

// NewComplexArrayListCap creates a new ComplexArrayList instance with an initialized length and capacity
func NewComplexArrayListCap(len, cap int) *ComplexArrayList {
    return &ComplexArrayList{data: make([]complex128, len, cap)}
}

// Add appends the specified element to the end of this list.
func (lst *ComplexArrayList) Add(e complex128) {
    lst.data = append(lst.data, e)
}

// AddSlice appends a slice of elements to the end of this list.
func (lst *ComplexArrayList) AddSlice(s []complex128) {
    lst.data = append(lst.data, s...)
}

// AddAll appends all elements of another ComplexArrayList to the end of this list.
func (lst *ComplexArrayList) AddAll(al *ComplexArrayList) {
    lst.data = append(lst.data, al.data...)
}

// Get returns the element at the specified position in this list.
func (lst *ComplexArrayList) Get(index int) complex128 {
    return lst.data[index]
}
// Set replaces the element at the specified position in this list with the specified element.
func (lst *ComplexArrayList) Set(index int, e complex128) {
    lst.data[index] = e
}
// Data returns the internal complex128 slice. If the array list performs structural modification, the returned
// slice could be no longer the one inside the array list.
func (lst *ComplexArrayList) Data() []complex128 {
    return lst.data
}

// The Swap method in sort.Interface.
func (lst *ComplexArrayList) Swap(i, j int) {
    lst.data[i], lst.data[j] = lst.data[j], lst.data[i]
}

// Insert inserts the specified element at the specified position in this list.
func (lst *ComplexArrayList) Insert(index int, e complex128) {
    lst.data = append(lst.data, 0)
    copy(lst.data[index + 1:], lst.data[index:])
    lst.data[index] = e
}

// The Push method
func (lst *ComplexArrayList) Push(e complex128) {
    lst.data = append(lst.data, e)
}

// The Pop method
func (lst *ComplexArrayList) Pop() (e complex128) {
    e = lst.data[len(lst.data) - 1]
    lst.data = lst.data[0:len(lst.data) - 1]
    return
}

// Remove removes the element at the specified position in this list.
func (lst *ComplexArrayList) Remove(index int) {
    lst.data = append(lst.data[0:index], lst.data[index + 1:]...)
}

// RemoveRange removes from this list all of the elements whose index is between from, inclusive, and to, exclusive.
func (lst *ComplexArrayList) RemoveRange(from, to int) {
    lst.data = append(lst.data[0:from], lst.data[to:]...)
}

// Clear removes all of the elements from this list.
func (lst *ComplexArrayList) Clear() {
    lst.data = lst.data[:0]
}

// Len returns the number of elements in this list.
//
// The Len method in sort.Interface.
func (lst *ComplexArrayList) Len() int {
    return len(lst.data)
}

// ComplexCmpAdapter is an adapter struct for an ArrayList with a cmp function.
type ComplexCmpAdapter struct {
    *ComplexArrayList
    cmp ComplexCmpFunc
}

// The Less method in sort.Interface
func (adp *ComplexCmpAdapter) Less(i, j int) bool {
    return adp.cmp(adp.data[i], adp.data[j]) < 0
}

// BinarySearch searchs a specified element e in a *sorted* list with binary search algorithm. If the list values are not sorted, the return values are undefined.
// If the element is found in the list, found equals true and pos is the index of the found element in the list.
// Otherwise found returns false and pos is the position where e is going to be inserted(and the resulting list is still in order)
func (adp *ComplexCmpAdapter) BinarySearch(e complex128) (pos int, found bool) {
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
func (lst *ComplexArrayList) NewCmpAdapter(cmp ComplexCmpFunc) *ComplexCmpAdapter {
    return &ComplexCmpAdapter{lst, cmp}
}

// String returns the internal data's string format as a result
func (lst *ComplexArrayList) String() string {
    return fmt.Sprintf("%v", lst.data)
}
   