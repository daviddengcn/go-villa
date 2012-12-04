go-villa
========

Some helper structs for go-lang

ArrayList
---------
ArrayList is a list of values(in the form of interface{}).

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

PriorityQueue
-------------
A ready-to-use priority queue struct given a less function(comparator). It encapsulates the heap package using the ArrayList struct.

Usage:
```go
    pq := villa.NewPriorityQueue(func(e1, e2 interface{}) bool {
        return e1.(int32) < e2.(int32)
    })
    pq.Push(10)
    pq.Push(20)

    vl := pq.Pop()
```