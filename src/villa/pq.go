package villa

import(
    "fmt"
    "container/heap"
)

type PriorityQueue struct {
    list *LessAdapter
}

func NewPriorityQueue(less LessFunc) *PriorityQueue {
    return &PriorityQueue{NewArrayList().NewLessAdapter(less)}
}

func (pq *PriorityQueue) Push(x interface{}) {
    heap.Push(pq.list, x)
}

func (pq *PriorityQueue) Pop() interface{} {
    return heap.Pop(pq.list)
}

func (pq *PriorityQueue) Remove(i int) interface{} {
    return heap.Remove(pq.list, i)
}

func (pq *PriorityQueue) Len() int {
    return pq.list.Len()
}

func (pq *PriorityQueue) String() string {
    return fmt.Sprintf("PriorityQueue(%d)", pq.list.Len())
}
