package villa

import (
	"math/rand"
	"testing"

	"github.com/golangplus/testing/assert"
)

func TestIntPriorityQueue(t *testing.T) {
	defer __(o_(t))

	pq := NewIntPriorityQueue(IntValueCompare)
	for i := 0; i < 1000; i++ {
		pq.Push(rand.Int())
	} // for i

	assert.Equal(t, "pq.Len()", pq.Len(), 1000)

	peek := pq.Peek()
	last := pq.Pop()
	assert.Equal(t, "pg.Peek()", peek, last)
	for i := 1; i < 1000; i++ {
		cur := pq.Pop()
		if cur < last {
			t.Errorf("%d should be larger than %d", cur, last)
		} // if
		last = cur
	} // for i
	t.Logf("%v", pq)
}

func TestIntPriorityQueueCap(t *testing.T) {
	pq := NewIntPriorityQueueCap(IntValueCompare, 10)
	assert.Equal(t, "pq.Len()", pq.Len(), 0)
}
