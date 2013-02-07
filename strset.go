package villa

import "fmt"

// StrSet is a set of strings
type StrSet map[string] bool

// NewStrSet create a string set with specified elements
func NewStrSet(els ...string) (s StrSet) {
	return s.Put(els...)
}

// Put add elements to the set. The set can be nil
func (s *StrSet) Put(els ...string) StrSet {
	if *s == nil {
		*s = make(map[string]bool)
	}
	for _, el := range els {
		(*s)[el] = true
	}
	
	return *s
}

// Delete removes elements from the set
func (s StrSet) Delete(els ...string) StrSet {
	for _, el := range els {
		delete(s, el)
	}
	
	return s
}

// In returns true if the specified element is in the set, false otherwise
func (s StrSet) In(el string) bool {
	_, in := s[el]
	return in
}

// Elements returns all elements in the set as a string slice
func (s StrSet) Elements() (els []string) {
	els = make([]string, 0, len(s))
	for el := range s {
		els = append(els, el)
	}
	
	return els
}

// String returns the string presentation of the set
func (s StrSet) String() string {
	return fmt.Sprint(s.Elements())
}