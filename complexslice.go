package villa

import "fmt"

// ComplexSlice is wrapper to a slice of complex128.
type ComplexSlice []complex128

// Add appends the specified element to the end of this slice.
func (s *ComplexSlice) Add(e... complex128) {
    *s = append(*s, e...)
}

// Insert inserts the specified element at the specified position in this slice.
// NOTE: the insertion algorithm is much better than the slice-trick in go community wiki
func (s *ComplexSlice) Insert(index int, e... complex128) {
    if cap(*s) >= len(*s) + len(e) {
        *s = (*s)[:len(*s) + len(e)]
    } else {
        *s = append(*s, e...)
    } // else
    copy((*s)[index + len(e):], (*s)[index:])
    copy((*s)[index:], e[:])
}

// The Swap method in sort.Interface.
func (s ComplexSlice) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

// Remove removes the element at the specified position in this slice.
func (s *ComplexSlice) Remove(index int) complex128 {
    e := (*s)[index]
    *s = append((*s)[0:index], (*s)[index + 1:]...)
    return e
}

// RemoveRange removes all of the elements whose index is between from, inclusive, and to, exclusive, from this slice.
func (s *ComplexSlice) RemoveRange(from, to int) {
    *s = append((*s)[0:from], (*s)[to:]...)
}

// Fill sets the elements between from, inclusive, and to, exclusive, to a speicified value.
func (s ComplexSlice) Fill(from, to int, vl complex128) {
    for i := from; i < to; i ++ {
        s[i] = vl
    } // for i
}

// Clear removes all of the elements from this slice.
func (s *ComplexSlice) Clear() {
    *s = (*s)[:0]
}

// String returns the internal data's string format as a result
func (s *ComplexSlice) String() string {
    return fmt.Sprintf("%v", *s)
}
