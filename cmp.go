package villa;

// CmpFunc is the function compares two elements.
type CmpFunc func(interface{}, interface{}) int

// IntCmpFunc is the function compares two int elements.
type IntCmpFunc func(int, int) int

// FloatCmpFunc is the function compares two float elements.
type FloatCmpFunc func(float64, float64) int

// ComplexCmpFunc is the function compares two complex128 elements.
type ComplexCmpFunc func(complex128, complex128) int


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

