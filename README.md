go-villa
========

Some helper structs for go-lang. Current supporint: priority queue, array list.

go.pkgdoc Link: http://go.pkgdoc.org/github.com/daviddengcn/go-villa

Priority Queues
---------------
A generic struct, named PriorityQueue, whose element is an interface{} and some structs whose element is a specific number type.

Using a priority queue requires a less function, with two elements to be compared as the arguments.

### PriorityQueue

It encapsulates the heap package using the ArrayList struct.

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
pq := villa.NewIntPriorityQueue(func(e1, e2 int) bool {
    func (a, b int) int {
        if a < b {
            return -1
        } else if a < b {
            return 1
        } // else if
        return 0
    })
pq.Push(10)
pq.Push(20)

vl := pq.Pop()
```

Array List
----------
Array list is a list of values(in the form of interface{} or some basic number type).

NOTE: these structs are not to replace the languange supported slice, but a higher level data-structures.

### ArrayList
Usage:
```go
lst := villa.NewArrayList()
lst.Add(10)
lst.Add(20)
lst.Insert(1, 30)
l := lst.Len()

sort.Sort(lst.NewCmpAdapter(
    func (a, b interface{}) int {
        if a.(int) < b.(int) {
            return -1
        } else if a.(int) < b.(int) {
            return 1
        } // else if
        return 0
    }))
```

### IntArrayList/FloatArrayList/ComplexArrayList
The following int can be replace with float or complex types(complex compare function needs rewriting).
```go
lst := villa.NewIntArrayList()
lst.Add(10)
lst.Add(20)
lst.Insert(1, 30)
l := lst.Len()

sort.Sort(lst.NewCmpAdapter(
    func (a, b int) int {
        if a < b {
            return -1
        } else if a < b {
            return 1
        } // else if
        return 0
    }))
```

Comparetor functions
--------------------
```go
// IntValueCompare compares the input int values a and b, returns -1 if a < b, 1 if a > b, and 0 otherwise.
// This is a natual IntCmpFunc.
func IntValueCompare(a, b int) int {
    if a < b {
        return -1
    } else if a > b {
        return 1
    } // else if
    
    return 0
}


// FloatValueCompare compares the input float64 values a and b, returns -1 if a < b, 1 if a > b, and 0 otherwise.
// This is a natual FloatCmpFunc.
func FloatValueCompare(a, b float64) int {
    if a < b {
        return -1
    } else if a > b {
        return 1
    } // else if
    
    return 0
}
```
