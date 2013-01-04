package villa

import "fmt"
import "reflect"

/*
Slice is a wrapper to a slice of interface{}.

You can either create an IntSlice instance directly, or converting the type when necessary.

Usage 1:
    var s Slice
    s.Add(10, "20", 30)
    s.InsertSlice(len(s), []A{A{50, 60}, A{70, 80}})
    s.Insert(1, 40, 50)
    s.Swap(1, len(s) - 1)
    s.RemoveRange(1, 3)
    s.Fill(0, len(s), 55)
    s.Clear()
    
Usage 2:
    type A struct {
        B, C int
    }
    
    var s []interface{}
    s = append(s, 10, "20", 30)
    (*Slice)(&s).InsertSlice(len(s), []A{A{50, 60}, A{70, 80}})
    (*Slice)(&s).Insert(1, 40, 50)
    Slice(s).Swap(1, len(s) - 1)
    (*Slice)(&s).RemoveRange(1, 3)
    Slice(s).Fill(0, len(s), 55)
    s = s[:0]
*/
type Slice []interface{}

// Add appends the specified element to the end of this slice.
func (s *Slice) Add(e... interface{}) {
    *s = append(*s, e...)
}

// Insert inserts specified elements at the specified position in this slice.
// NOTE: the insertion algorithm is much better than the slice-trick in go community wiki
func (s *Slice) Insert(index int, e... interface{}) {
    if cap(*s) >= len(*s) + len(e) {
        *s = (*s)[:len(*s) + len(e)]
    } else {
        *s = append(*s, e...)
    } // else
    copy((*s)[index + len(e):], (*s)[index:])
    copy((*s)[index:], e)
}

// InsertSlice inserts the elements of a slice at the specified position in this slice.
// This method is useful when some elements in a slice *NOT* of type []interface{} are to be inserted.
func (s *Slice) InsertSlice(index int, src interface{}) {
    v := reflect.ValueOf(src)
    if v.Kind() != reflect.Slice {
        panic(fmt.Sprintf("%v is not a slice!", src))
    } // if
    
    n := v.Len()
    if cap(*s) >= len(*s) + n {
        *s = (*s)[:len(*s) + n]
        copy((*s)[index+n:], (*s)[index:])
    } else {
        ss := make([]interface{}, len(*s) + n)
        copy(ss[:index], *s)
        copy(ss[index + n:], (*s)[index:])
        *s = ss
    } // else
    
    for i := 0; i < n; i ++ {
        (*s)[i + index] = v.Index(i).Interface()
    } // for i
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
