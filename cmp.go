package villa;

// CmpFunc is a function comparing two elements. The function returns a positive value if a > b, a negative value if a < b, and 0 otherwise.
type CmpFunc func(a, b interface{}) int

// IntCmpFunc is a function comparing two int elements. The function returns a positive value if a > b, a negative value if a < b, and 0 otherwise.
type IntCmpFunc func(a, b int) int

// FloatCmpFunc is a function comparing two float elements. The function returns a positive value if a > b, a negative value if a < b, and 0 otherwise.
type FloatCmpFunc func(a, b float64) int

// ComplexCmpFunc is a function comparing two complex128 elements. The function returns a positive value if a > b, a negative value if a < b, and 0 otherwise.
type ComplexCmpFunc func(a, b complex128) int


// IntValueCompare compares the input int values a and b, returns -1 if a < b, 1 if a > b, and 0 otherwise.
// This is a natural IntCmpFunc.
func IntValueCompare(a, b int) int {
    if a < b {
        return -1
    } else if a > b {
        return 1
    } // else if
    
    return 0
}


// FloatValueCompare compares the input float64 values a and b, returns -1 if a < b, 1 if a > b, and 0 otherwise.
// This is a natural FloatCmpFunc.
func FloatValueCompare(a, b float64) int {
    if a < b {
        return -1
    } else if a > b {
        return 1
    } // else if
    
    return 0
}

