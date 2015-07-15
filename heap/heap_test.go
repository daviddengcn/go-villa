package heap

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/golangplus/testing/assert"
)

type IntHeap []int

func (h *IntHeap) Pop() int {
	PopToLast(sort.IntSlice(*h))
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]

	return res
}

func (h *IntHeap) Push(x int) {
	*h = append(*h, x)
	PushLast(sort.IntSlice(*h))
}

type Data struct {
	Value    string
	Priority int
}

func TestIntHeap(t *testing.T) {
	var h IntHeap

	for i := 0; i < 1000; i++ {
		h.Push(rand.Int())
	}

	assert.Equal(t, "len(h)", len(h), 1000)

	peek := h[0]
	last := h.Pop()
	assert.Equal(t, "h.Peek()", peek, last)
	//    fmt.Println(h)
	for i := 1; i < 1000; i++ {
		cur := h.Pop()
		if cur < last {
			t.Errorf("%d should be larger than %d", cur, last)
		}
		last = cur
	}
}

type DataHeap []Data

func (h DataHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h DataHeap) Len() int {
	return len(h)
}

func (h DataHeap) Less(i, j int) bool {
	return h[i].Priority < h[j].Priority
}

func (h *DataHeap) Pop() Data {
	PopToLast(h)
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]

	return res
}

func (h *DataHeap) Push(x Data) {
	*h = append(*h, x)
	PushLast(h)
}

func TestDataHeap(t *testing.T) {
	var h DataHeap

	for i := 0; i < 1000; i++ {
		h.Push(Data{"A", rand.Int()})
	}

	assert.Equal(t, "len(h)", len(h), 1000)

	peek := h[0]
	last := h.Pop()
	assert.Equal(t, "h.Peek()", peek, last)

	for i := 1; i < 1000; i++ {
		cur := h.Pop()
		if cur.Priority < last.Priority {
			t.Errorf("%v should be larger than %v", cur, last)
		}
		last = cur
	}
}
