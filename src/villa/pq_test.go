package villa

import(
    "testing"
//    "sort"
    "math/rand"
    "fmt"
)

func TestPriorityQueue(t *testing.T) {
    pq := NewPriorityQueue(func(e1, e2 interface{}) bool {
        return e1.(int32) < e2.(int32)
    })
    for i := 0; i < 1000; i ++ {
        pq.Push(rand.Int31())
    } // for i
    
    AssertEquals(t, "pq.Len()", pq.Len(), 1000)

    last := pq.Pop().(int32)
    for i := 1; i < 1000; i ++ {
        cur := pq.Pop().(int32)
        if cur < last {
            t.Errorf("%d should be larger than %d", cur, last)
        } // if
        last = cur
    } // for i
    fmt.Println(pq)
}

