package villa

import(
    "fmt"
    "container/heap"
)

// PriorityQueue is an unbounded priority queue based on a priority heap.
//
// A less function, typed LessFunc, needs to be specified. This struct is a better
// encapsulation of the "container/heap" package provided by go standard library.
type PriorityQueue struct {
    list *LessAdapter
}

// NewPriorityQueue creates a PriorityQueue instance with a specified less function.
func NewPriorityQueue(less LessFunc) *PriorityQueue {
    return &PriorityQueue{NewArrayList().NewLessAdapter(less)}
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
