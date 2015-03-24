package villa

import (
	"fmt"
)

// Errorf is similar to fmt.Errorf with caller's position in the source
func Errorf(format string, a ...interface{}) error {
	return fmt.Errorf("%s: %s", LinePos(1), fmt.Sprintf(format, a...))
}

// Error is similar to fmt.Error with caller's position in the source
func Error(v ...interface{}) error {
	return fmt.Errorf("%s: %s", LinePos(1), fmt.Sprint(v...))
}

// ErrorfN is similar to villa.Errorf with a specified line-pos depth. when n == 1, it is
// equivalent to villa.Errorf.
func ErrorfN(n int, format string, a ...interface{}) error {
	return fmt.Errorf("%s: %s", LinePos(n), fmt.Sprintf(format, a...))
}

// ErrorN is similar to villa.Error with a specified line-pos depth. when n == 1, it is
// equivalent to villa.Error.
func ErrorN(n int, v ...interface{}) error {
	return fmt.Errorf("%s: %s", LinePos(n), fmt.Sprint(v...))
}
