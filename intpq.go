package villa

import(
    "fmt"
)

// IntPriorityQueue is an unbounded priority queue based on a priority heap.
//
// A compare function, typed IntCmpFunc, needs to be specified.
// Usage:
//     pq := villa.NewIntPriorityQueue(
//         func (a, b int) int {
//             if a < b {
//                 return -1
//             } else if a < b {
//                 return 1
//             } // else if
//             return 0
//         })
//     pq.Push(10)
//     pq.Push(20)
//     
//     vl := pq.Pop()
type IntPriorityQueue struct {
    list []int
    cmp IntCmpFunc
}

// NewIntPriorityQueue creates a IntPriorityQueue instance with a specified campare function.
func NewIntPriorityQueue(cmp IntCmpFunc) *IntPriorityQueue {
    return &IntPriorityQueue{nil, cmp}
}

// NewIntPriorityQueueCap creates a IntPriorityQueue instance with a specified compare function and a capacity
func NewIntPriorityQueueCap(cmp IntCmpFunc, cap int) *IntPriorityQueue {
    return &IntPriorityQueue{make([] int, 0, cap), cmp}
}

func (pq *IntPriorityQueue)intUp(j int) {
    for {
        i := (j - 1) / 2 // parent
        if i == j || pq.cmp(pq.list[i], pq.list[j]) < 0 {
            break
        } // if
        pq.list[i], pq.list[j] = pq.list[j], pq.list[i]
        j = i
    } // for true
}

func (pq *IntPriorityQueue)intDown(i, n int) {
    for {
        j1 := 2*i + 1
        if j1 >= n {
            break
        } // if
        j := j1 // left child
        if j2 := j1 + 1; j2 < n && pq.cmp(pq.list[j1], pq.list[j2]) >= 0 {
            j = j2 // = 2*i + 2  // right child
        } // if
        if pq.cmp(pq.list[i], pq.list[j]) < 0 {
            break
        } // if
        pq.list[i], pq.list[j] = pq.list[j], pq.list[i]
        i = j
    } // for
}

// Push inserts the specified element into this priority queue.
func (pq *IntPriorityQueue) Push(x int) {
    pq.list = append(pq.list, x)
    pq.intUp(len(pq.list) - 1)
}

// Pop retrieves and removes the head of this queue, or returns nil if this queue is empty.
func (pq *IntPriorityQueue) Pop() int {
    n := len(pq.list) - 1
    res := pq.list[0]
    pq.list[0] = pq.list[n]
    pq.intDown(0, n)
    pq.list = pq.list[:n]
    return res
}

// Peek retrieves the head of this queue, or returns 0 if this queue is empty.
func (pq *IntPriorityQueue) Peek() int {
    if len(pq.list) > 0 {
        return pq.list[0]
    } // if
        
    return 0
}

// Remove removes the element at index i from the priority queue.
func (pq *IntPriorityQueue) Remove(i int) int {
    n := len(pq.list) - 1
    res := pq.list[i]
    if n != i {
        pq.list[i] = pq.list[n]
        pq.intDown(i, n)
        pq.intUp(i)
    } // if
    pq.list = pq.list[:n]
    return res
}

// Len returns the number of elements in this queue.
func (pq *IntPriorityQueue) Len() int {
    return len(pq.list)
}

// String returns a string with value of "IntPriorityQueue(#)" where # is the number of elements in the queue
func (pq *IntPriorityQueue) String() string {
    return fmt.Sprintf("IntPriorityQueue(%d)", len(pq.list))
}
 