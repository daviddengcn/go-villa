go-villa
========

Some helper structs for go-lang. Current supporint: priority queue, array list.

Priority queues
---------------
A generic struct, named PriorityQueue, whose element is an interface{} and some structs whose element is a specific number type.

Using a priority queue requires a less function, with two elements to be compared as the arguments.

### PriorityQueue

It encapsulates the heap package using the ArrayList struct.

Usage:
```go
    pq := villa.NewPriorityQueue(func(e1, e2 interface{}) bool {
        return e1.(int32) < e2.(int32)
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
        return e1 < e2
    })
    pq.Push(10)
    pq.Push(20)

    vl := pq.Pop()
```

Array list
---------
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

    sort.Sort(lst.NewLessAdapter(func(e1, e2 interface{}) bool {
        return e1.(int) < e2.(int)
    }))
```

### IntArrayList/FloatArrayList/ComplexArrayList
The following int can be replace with float or complex types.
```go
    lst := villa.NewIntArrayList()
    lst.Add(10)
    lst.Add(20)
    lst.Insert(1, 30)
    l := lst.Len()

    sort.Sort(lst.NewLessAdapter(func(e1, e2 int) bool {
        return e1 < e2
    }))
```
