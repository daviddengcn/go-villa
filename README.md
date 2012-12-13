go-villa
========

Some helper types for go-lang. Current supporint: priority queue, slice wrapper.

go.pkgdoc Link: http://go.pkgdoc.org/github.com/daviddengcn/go-villa

Priority Queues
---------------
A generic struct, named PriorityQueue, whose element is an interface{} and some structs whose element is a specific number type.

Using a priority queue requires a less function, with two elements to be compared as the arguments.

### PriorityQueue

It encapsulates the heap package using the Slice struct.

Usage:
```go
pq := villa.NewPriorityQueue(
    func (a, b interface{}) int {
        if a.(int) < b.(int) {
            return -1
        } else if a.(int) < b.(int) {
            return 1
        } // else if
        return 0
    })
pq.Push(10)
pq.Push(20)

vl := pq.Pop()
```

### IntPriorityQueue

It rewrites the algorithm in heap package. Integers are internally stored in an int slice.
Usage:
```go
pq := villa.NewIntPriorityQueue(IntValueCompare)
pq.Push(10)
pq.Push(20)

vl := pq.Pop()
```

Slice Wrappers
--------------
Slice is a warpper for go slices.

### Slice
Usage:
```go
var s Slice
s.Add(10, 20)
s.Insert(1, 30)
l := len(s)

sl := s.NewSortList(
    func (a, b interface{}) int {
        if a.(int) < b.(int) {
            return -1
        } else if a.(int) < b.(int) {
            return 1
        } // else if
        return 0
    })
sort.Sort(sl)
p, found := sl.BinarySearch(20)
```

### IntSlice/FloatSlice/ComplexSlice
The following int can be replace with float or complex types(complex compare function needs rewriting).
```go
var s IntSlice
s.Add(10, 20)
s.Insert(1, 30)
l := len(s)

sl := lst.NewSortList(villa.IntValueCompare)
sort.Sort(sl)
p, found := sl.BinarySearch(20)
```

Comparator functions
--------------------
```go
// IntValueCompare compares the input int values a and b, returns -1 if a < b, 1 if a > b, and 0 otherwise.
// This is a natural IntCmpFunc.
func IntValueCompare(a, b int) int {
    if a < b {
        return -1
    } else if a > b {
        return 1
    } // else if
    
    return 0
}


// FloatValueCompare compares the input float64 values a and b, returns -1 if a < b, 1 if a > b, and 0 otherwise.
// This is a natural FloatCmpFunc.
func FloatValueCompare(a, b float64) int {
    if a < b {
        return -1
    } else if a > b {
        return 1
    } // else if
    
    return 0
}
```
