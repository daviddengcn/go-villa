package villa

// N is a very long slice of elements with size of zero-bytes.
// It is useful for generating a serial of numbers in a for-range clause. e.g.
//    for i := range villa.N[:200] { ... }
// The above i in the above for clause will range from 0 to 200(exclusive).
var N [int(^uint(0) >> 1)]struct{}

// An variable of zero-size bytes
type Empty struct{}
