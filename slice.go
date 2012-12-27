package villa

import "fmt"
import "reflect"

// Slice is wrapper to a slice of interface{}.
type Slice []interface{}

// Add appends the specified element to the end of this slice.
func (s *Slice) Add(e... interface{}) {
    *s = append(*s, e...)
}

// AddSlice appends the elements of a slice to the slice.
//
// The type of the elements of the slice could be any type. For instance,
//    type I interface {
//        Func()
//    }
//    type SI []I
//    si := make(SI, 2)
//    s.AddSlice(si)
func (s *Slice) AddSlice(src interface{}) {
    v := reflect.ValueOf(src)
    if v.Kind() != reflect.Slice {
        panic(fmt.Sprintf("%v is not a slice!", src))
    } // if
    
    n := v.Len()
    for i := 0; i < n; i ++ {
        *s = append(*s, v.Index(i).Interface())
    } // for i
}

// Insert inserts the specified element at the specified position in this slice.
// NOTE: the insertion algorithm is much better than the slice-trick in go community wiki
func (s *Slice) Insert(index int, e... interface{}) {
    if cap(*s) >= len(*s) + len(e) {
        *s = (*s)[:len(*s) + len(e)]
    } else {
        *s = append(*s, e...)
    } // else
    copy((*s)[index + len(e):], (*s)[index:])
    copy((*s)[index:], e[:])
}

// The Swap method in sort.Interface.
func (s Slice) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

// Remove removes the element at the specified position in this slice.
// The hole in the original slice is filled with a nil value.
func (s *Slice) Remove(index int) interface{} {
    e := (*s)[index]
    copy((*s)[index:], (*s)[index + 1:])
    (*s)[len(*s) - 1] = nil
    *s = (*s)[:len(*s) - 1]
    return e
}

// RemoveRange removes all of the elements whose index is between from, inclusive, and to, exclusive, from this slice.
// The holes in the original slice are filled with nil values.
func (s *Slice) RemoveRange(from, to int) {
    if to <= from {
        return
    } // if
    
    copy((*s)[from:], (*s)[to:])
    n := len(*s);  l := n - to + from
    for i := l; i < n; i ++ {
        (*s)[i] = nil
    } // for i
    *s = (*s)[:l]
}

// Fill sets the elements between from, inclusive, and to, exclusive, to a speicified value.
func (s Slice) Fill(from, to int, vl interface{}) {
    for i := from; i < to; i ++ {
        s[i] = vl
    } // for i
}

// Clear sets the slice to nil.
func (s *Slice) Clear() {
    *s = (*s)[:0]
}

// String returns the internal data's string format as a result
func (s *Slice) String() string {
    return fmt.Sprintf("%v", *s)
}
