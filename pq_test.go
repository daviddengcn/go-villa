package villa

import(
    "testing"
//    "sort"
    "math/rand"
    "fmt"
)

func TestPriorityQueue(t *testing.T) {
    fmt.Println("== Begin TestPriorityQueue...");
    defer fmt.Println("== End TestPriorityQueue.");
    
    pq := NewPriorityQueue(intInterfaceCmpFunc)
    for i := 0; i < 1000; i ++ {
        pq.Push(rand.Int())
    } // for i
    
    AssertEquals(t, "pq.Len()", pq.Len(), 1000)

    peek := pq.Peek().(int)
    last := pq.Pop().(int)
    AssertEquals(t, "pg.Peek()", peek, last)
    for i := 1; i < 1000; i ++ {
        cur := pq.Pop().(int)
        if cur < last {
            t.Errorf("%d should be larger than %d", cur, last)
        } // if
        last = cur
    } // for i
    fmt.Println(pq)
}

func TestPriorityQueueCap(t *testing.T) {
    pq := NewPriorityQueueCap(intInterfaceCmpFunc, 10)
    AssertEquals(t, "pq.Len()", pq.Len(), 0)
}
