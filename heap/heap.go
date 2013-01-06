/*
This heap package implements a very similar function to the build-in heap package except the elements are not necessarily interface{}, but can be any type.

The trick is using the last element as the in/out place. Push/Pop/Remove are replaced with PushLast/PopToLast/RemoveToLast, respectively. An heap with int value can be easily implemented as follow:

    type IntHeap []int
    func (h *IntHeap) Pop() int {
        heap.PopToLast(sort.IntSlice(*h))
        res := (*h)[len(*h) - 1]
        *h = (*h)[:len(*h) - 1]
        
        return res
    }
    
    func (h *IntHeap) Push(x int) {
        *h = append(*h, x)
        heap.PushLast(sort.IntSlice(*h))
    }
*/
package heap

import "sort"

// A none-empty heap must be initialized before any of the heap operations
// can be used. Init is idempotent with respect to the heap invariants
// and may be called whenever the heap invariants may have been invalidated.
// Its complexity is O(n) where n = h.Len().
//
func Init(h sort.Interface) {
    // heapify
    n := h.Len()
    for i := n/2 - 1; i >= 0; i-- {
        down(h, i, n)
    } // for i
}

// Push pushes the last element of the heap, which is considered not as the part of the heap, onto the heap. The complexity is
// O(log(n)) where n = h.Len().
//
// NOTE You need to append the element to be pushed as the last element before calling to this method.
func PushLast(h sort.Interface) {
    up(h, h.Len()-1)
}

// Pop removes the minimum element (according to Less) from the heap
// and place it at the last element of the heap. The complexity is O(log(n)) where n = h.Len().
// Same as Remove(h, 0).
//
// NOTE You need to remove the last element after calling to this method.
func PopToLast(h sort.Interface) {
    n := h.Len() - 1
    h.Swap(0, n)
    down(h, 0, n)
}

// Remove removes the element at index i from the heap and place it at the last element of the heap.
// The complexity is O(log(n)) where n = h.Len().
//
// NOTE You need to remove the last element after calling to this method.
func RemoveToLast(h sort.Interface, i int) {
    n := h.Len() - 1
    if n != i {
        h.Swap(i, n)
        down(h, i, n)
        up(h, i)
    } //  if
}

func up(h sort.Interface, j int) {
    for {
        i := (j - 1) / 2 // parent
        if i == j || h.Less(i, j) {
            break
        } // if
        h.Swap(i, j)
        j = i
    } // for
}

func down(h sort.Interface, i, n int) {
    for {
        j1 := 2*i + 1
        if j1 >= n {
            break
        } // if
        j := j1 // left child
        if j2 := j1 + 1; j2 < n && !h.Less(j1, j2) {
            j = j2 // = 2*i + 2  // right child
        } // if
        if h.Less(i, j) {
            break
        } // if
        h.Swap(i, j)
        i = j
    } // for
}
