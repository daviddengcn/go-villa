package villa

import(
    "fmt"
    "container/heap"
)

// PriorityQueue is an unbounded priority queue based on a priority heap.
//
// A compare function, typed CmpFunc, needs to be specified. This struct is a better
// encapsulation of the "container/heap" package provided by go standard library.
// Usage:
//     pq := villa.NewPriorityQueue(
//        func (a, b interface{}) int {
//            if a.(int) < b.(int) {
//                return -1
//            } else if a.(int) < b.(int) {
//                return 1
//            } // else if
//            return 0
//        })
//     pq.Push(10)
//     pq.Push(20)
//     
//     vl := pq.Pop()
type PriorityQueue struct {
    list *SortList
}

// NewPriorityQueue creates a PriorityQueue instance with a specified compare function.
func NewPriorityQueue(cmp CmpFunc) *PriorityQueue {
    return &PriorityQueue{(&Slice{}).NewSortList(cmp)}
}

// NewPriorityQueue creates a PriorityQueue instance with a specified compare function and a capacity
func NewPriorityQueueCap(cmp CmpFunc, cap int) *PriorityQueue {
    lst := make(Slice, 0, cap)
    return &PriorityQueue{lst.NewSortList(cmp)}
}

// Push inserts the specified element into this priority queue.
func (pq *PriorityQueue) Push(x interface{}) {
    heap.Push(pq.list, x)
}

// Pop retrieves and removes the head of this queue, or returns nil if this queue is empty.
func (pq *PriorityQueue) Pop() interface{} {
    return heap.Pop(pq.list)
}

// Peek retrieves the head of this queue, or returns nil if this queue is empty.
func (pq *PriorityQueue) Peek() interface {} {
    if pq.list.Len() > 0 {
        return pq.list.Get(0)
    } // if
        
    return nil
}

// Remove removes the element at index i from the priority queue.
func (pq *PriorityQueue) Remove(i int) interface{} {
    return heap.Remove(pq.list, i)
}

// Len returns the number of elements in this queue.
func (pq *PriorityQueue) Len() int {
    return pq.list.Len()
}

// String returns a string with value of "PriorityQueue(Len())"
func (pq *PriorityQueue) String() string {
    return fmt.Sprintf("PriorityQueue(%d)", pq.list.Len())
}
