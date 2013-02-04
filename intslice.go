package villa

//import "fmt"

/*
IntSlice is a wrapper to a slice of int.

You can either create an IntSlice instance directly, or converting the type when necessary.

Usage 1:
    var s IntSlice
    s.Add(10, 20, 30)
    s.Insert(1, 40, 50)
    s.Swap(1, len(s) - 1)
    s.RemoveRange(1, 3)
    s.Fill(0, len(s), 55)
    s.Clear()
    
Usage 2:
    var s []int
    s = append(s, 10, 20, 30)
    (*IntSlice)(&s).Insert(1, 40, 50)
    IntSlice(s).Swap(1, len(s) - 1)
    (*IntSlice)(&s).RemoveRange(1, 3)
    IntSlice(s).Fill(0, len(s), 55)
    s = s[:0]
*/
type IntSlice []int

// Add appends the specified element to the end of this slice.
func (s *IntSlice) Add(e... int) *IntSlice {
    *s = append(*s, e...)
	return s	
}

// Insert inserts the specified element at the specified position.
// NOTE: the insertion algorithm is much better than the slice-trick in go community wiki
func (s *IntSlice) Insert(index int, e... int) {
    if cap(*s) >= len(*s) + len(e) {
        *s = (*s)[:len(*s) + len(e)]
    } else {
        *s = append(*s, e...)
    } // else
    copy((*s)[index + len(e):], (*s)[index:])
    copy((*s)[index:], e[:])
}

// The Swap method in sort.Interface.
func (s IntSlice) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

// Remove removes the element at the specified position in this slice.
func (s *IntSlice) Remove(index int) int {
    e := (*s)[index]
    *s = append((*s)[0:index], (*s)[index + 1:]...)
    return e
}

// RemoveRange removes all of the elements whose index is between from, inclusive, and to, exclusive, from this slice.
func (s *IntSlice) RemoveRange(from, to int) {
    *s = append((*s)[0:from], (*s)[to:]...)
}

// Fill sets the elements between from, inclusive, and to, exclusive, to a speicified value.
func (s IntSlice) Fill(from, to int, vl int) {
    for i := from; i < to; i ++ {
        s[i] = vl
    } // for i
}

// Clear removes all of the elements from this slice.
func (s *IntSlice) Clear() {
    *s = (*s)[:0]
}

// Equals returns true if a given slice has the same contents with the slice
func (s IntSlice) Equals(t []int) bool {
    if len(s) != len(t) {
        return false
    } // if
    
    for i := range(s) {
        if s[i] != t[i] {
            return false
        } // if
    } // for i
    
    return true
}
