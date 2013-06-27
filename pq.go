package villa

import (
	"container/heap"
	"fmt"
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
	list *pqList
}

type pqList struct {
	Slice
	cmp CmpFunc
}

// The Push method in heap.Interface.
func (l *pqList) Push(e interface{}) {
	l.Add(e)
}

// The Pop method in heap.Interface.
func (l *pqList) Pop() interface{} {
	return l.Remove(len(l.Slice) - 1)
}

// The Len method in sort.Interface.
func (l *pqList) Len() int {
	return len(l.Slice)
}

// The Less method in sort.Interface
func (l *pqList) Less(i, j int) bool {
	return l.cmp(l.Slice[i], l.Slice[j]) <= 0
}

// NewPriorityQueue creates a PriorityQueue instance with a specified compare function.
func NewPriorityQueue(cmp CmpFunc) *PriorityQueue {
	return &PriorityQueue{&pqList{cmp: cmp}}
}

// NewPriorityQueue creates a PriorityQueue instance with a specified compare function and a capacity
func NewPriorityQueueCap(cmp CmpFunc, cap int) *PriorityQueue {
	return &PriorityQueue{&pqList{Slice: make(Slice, 0, cap), cmp: cmp}}
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
func (pq *PriorityQueue) Peek() interface{} {
	if pq.list.Len() > 0 {
		return pq.list.Slice[0]
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
