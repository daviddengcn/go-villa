package villa

//import "fmt"

// FloatSlice is wrapper to a slice of float64.
//
// See examples of IntSlice for reference.
type FloatSlice []float64

// Add appends the specified element to the end of this slice.
func (s *FloatSlice) Add(e... float64) *FloatSlice {
    *s = append(*s, e...)
	return s
}

// Insert inserts the specified element at the specified position.
// NOTE: the insertion algorithm is much better than the slice-trick in go community wiki
func (s *FloatSlice) Insert(index int, e... float64) {
    if cap(*s) >= len(*s) + len(e) {
        *s = (*s)[:len(*s) + len(e)]
    } else {
        *s = append(*s, e...)
    } // else
    copy((*s)[index + len(e):], (*s)[index:])
    copy((*s)[index:], e[:])
}

// The Swap method in sort.Interface.
func (s FloatSlice) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

// Remove removes the element at the specified position in this slice.
func (s *FloatSlice) Remove(index int) float64 {
    e := (*s)[index]
    *s = append((*s)[0:index], (*s)[index + 1:]...)
    return e
}

// RemoveRange removes all of the elements whose index is between from, inclusive, and to, exclusive, from this slice.
func (s *FloatSlice) RemoveRange(from, to int) {
    *s = append((*s)[0:from], (*s)[to:]...)
}

// Fill sets the elements between from, inclusive, and to, exclusive, to a speicified value.
func (s FloatSlice) Fill(from, to int, vl float64) {
    for i := from; i < to; i ++ {
        s[i] = vl
    } // for i
}

// Clear removes all of the elements from this slice.
func (s *FloatSlice) Clear() {
    *s = (*s)[:0]
}


// Equals returns true if a given slice has the same contents (with maximum error of epsilon) with the slice
func (s FloatSlice) Equals(t []float64, epsilon float64) bool {
    if len(s) != len(t) {
        return false
    } // if
    
    for i := range(s) {
        e := s[i] - t[i]
        if e > epsilon || e < -epsilon {
            return false
        } // if
    } // for i
    
    return true
}

