package villa

import "fmt"
import "reflect"

/*
StringSlice is a wrapper to a slice of interface{}.

You can either create an StringSlice instance directly, or converting the type when necessary.

Usage 1:
    type A struct {
        B string
        C int
    }

    var s StringSlice
    s.Add(10, "B", 30)
    s.InsertSlice(len(s), []A{A{"E", 60}, A{"G", 80}})
    s.Insert(1, "D", "E")
    s.Swap(1, len(s) - 1)
    s.RemoveRange(1, 3)
    s.Fill(0, len(s), "EE")
    s.Clear()

Usage 2:
    type A struct {
        B string
        C int
    }

    var s []string
    (*StringSlice)(&s).Add(10, "B", 30)
    (*StringSlice)(&s).InsertSlice(len(s), []A{A{"E", 60}, A{"G", 80}})
    (*StringSlice)(&s).Insert(1, "D", "E")
    StringSlice(s).Swap(1, len(s) - 1)
    (*StringSlice)(&s).RemoveRange(1, 3)
    StringSlice(s).Fill(0, len(s), "EE")
    (*StringSlice)(&s).Clear()
*/
type StringSlice []string

// Add appends string presentation of the specified elements to the end of this slice.
func (s *StringSlice) Add(e ...interface{}) *StringSlice {
	s.InsertSlice(len(*s), e)
	return s
}

// Insert inserts string presentation of the specified elements at the specified position.
// NOTE: the insertion algorithm is much better than the slice-trick in go community wiki
func (s *StringSlice) Insert(index int, e ...interface{}) {
	s.InsertSlice(index, e)
}

// InsertSlice inserts string presentation of elements of a slice at the specified position.
func (s *StringSlice) InsertSlice(index int, src interface{}) {
	v := reflect.ValueOf(src)
	if v.Kind() != reflect.Slice {
		panic(fmt.Sprintf("%v is not a slice!", src))
	} // if

	n := v.Len()
	if cap(*s) >= len(*s)+n {
		*s = (*s)[:len(*s)+n]
		copy((*s)[index+n:], (*s)[index:])
	} else {
		ss := make([]string, len(*s)+n)
		copy(ss[:index], *s)
		copy(ss[index+n:], (*s)[index:])
		*s = ss
	} // else

	for i := 0; i < n; i++ {
		vi := v.Index(i)
		if vi.Kind() == reflect.String {
			(*s)[i+index] = vi.String()
		} else {
			(*s)[i+index] = fmt.Sprint(vi.Interface())
		} // else
	} // for i
}

// The Swap method in sort.Interface.
func (s StringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Remove removes the element at the specified position in this slice.
// The hole in the original slice is filled with a zero value.
func (s *StringSlice) Remove(index int) string {
	e := (*s)[index]
	copy((*s)[index:], (*s)[index+1:])
	(*s)[len(*s)-1] = ""
	*s = (*s)[:len(*s)-1]
	return e
}

// RemoveRange removes all of the elements whose index is between from, inclusive, and to, exclusive, from this slice.
// The holes in the original slice are filled with nil values.
func (s *StringSlice) RemoveRange(from, to int) {
	if to <= from {
		return
	} // if

	copy((*s)[from:], (*s)[to:])
	n := len(*s)
	l := n - to + from
	for i := l; i < n; i++ {
		(*s)[i] = ""
	} // for i
	*s = (*s)[:l]
}

// Pop removes and returns the last element in the slice
func (s *StringSlice) Pop() string {
	return s.Remove(len(*s) - 1)
}

// Fill sets the elements between from, inclusive, and to, exclusive, to a speicified value.
func (s StringSlice) Fill(from, to int, vl string) {
	for i := from; i < to; i++ {
		s[i] = vl
	} // for i
}

// Clear sets the slice to an zero-length slice.
func (s *StringSlice) Clear() {
	s.Fill(0, len(*s), "")
	*s = (*s)[:0]
}

// Equals returns true if a given slice has the same contents with the slice
func (s StringSlice) Equals(t []string) bool {
	if len(s) != len(t) {
		return false
	} // if

	for i := range s {
		if s[i] != t[i] {
			return false
		} // if
	} // for i

	return true
}
