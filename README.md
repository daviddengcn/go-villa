go-villa
========
Some helper types for go-lang. Current supporint: priority queue, slice wrapper, binary-search, merge-sort.

godoc Link: http://godoc.org/github.com/daviddengcn/go-villa

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
Slice is a wrapper for go slices. Implemented methods include: Add, AddSlice, Insert, Swap, Remove, RemoveRange, Fill, Clear

### Slice
Usage:
```go
type A struct {
    B, C int
}

var s Slice
s.Add(10, "20", 30)
s.InsertSlice(len(s), []A{A{50, 60}, A{70, 80}})
s.Insert(1, 40, 50)
s.Swap(1, len(s) - 1)
s.RemoveRange(1, 3)
s.Fill(0, len(s), 55)
s.Clear()
```

### StringSlice
StringSlice is a wrapper for []string. Using Add/Insert/InsertSlice methods, the values of other type can be converted and added to the string slice.
Usage:
```go
type A struct {
    B string
    C int
}

var s StringSlice
s.Add(10, "B", 30)
s.InsertSlice(len(s), []A{A{"E", 60}, A{"G", 80}})
s.Insert(1, "D", "E")
s.Swap(1, len(s) - 1)
s.RemoveRange(1, 3)
s.Fill(0, len(s), "EE")
s.Clear()
```

### IntSlice/FloatSlice/ComplexSlice
The following int can be replace with float or complex types(complex compare function needs rewriting).
Usage(of IntSlice):
```go
var s IntSlice
s.Add(10, 20, 30)
s.Insert(1, 40, 50)
s.Swap(1, len(s) - 1)
s.RemoveRange(1, 3)
s.Fill(0, len(s), 55)
s.Clear()
```

Comparator functions
--------------------
The common comparator function which compares elements and return the value <0, =0 or >0, if a < b, a==b, or a > b respectively.

Some algorithms that needs a comparator are defined in their methods, including sort(using build-in sort package algorithm), binary-search, and merge.
Cast your own comparator function to the corresponding comparator type to use them:
```go
func MyCmp(a, b int) int {
   ...
}

var s, l []int
cmp := IntCmpFunc(MyCmp)
cmp.Sort(s)
cmp.BinarySearch(s, e)

cmp.Sort(l)
cmp.Merge(s, l)
```

Two comparators are defined for natual orders of ints and floats.
```go
var IntValueCompare IntCmpFunc
var FloatValueCompare FloatCmpFunc
```

LICENSE
---------------
This library is distributed under BSD license.
