package villa

import (
	"fmt"
)

// N is a very long slice of elements with size of zero-bytes.
// It is useful for generating a serial of numbers in a for-range clause. e.g.
//    for i := range villa.N[:200] { ... }
// The above i in the above for clause will range from 0 to 200(exclusive).
var N [int(^uint32(0) >> 1)]struct{}

// An variable of zero-size bytes
type Empty struct{}

// A channel for signaling stop
type Stop chan Empty

// Signal to stop
func (stop Stop) Stop() {
	stop <- Empty{}
}

// Create a Stop with buffer size 1
func NewStop() Stop {
	return make(Stop, 1)
}

/*
NestedError is an error with current message and nested error.

Use NestErrorf to generate a NestedError. It returns a nil for a nil nested
error.

Use NestedError.Deepest() to fetch the cause error.
*/
type NestedError interface {
	error
	// Message returns the messsage of this error
	Message() string
	// Nested returns the nested error
	Nested() error
	/*
		Deepest returns the deepest non-NestedError error, which is the
		original cause error.
	*/
	Deepest() error
}

type nestedError struct {
	message string
	nested error
}

// Error implements error interface
func (err *nestedError) Error() string {
	if err.nested == nil {
		return err.message
	}
	return err.message + ": " + err.nested.Error()
}

func (err *nestedError) Message() string {
	return err.message
}

func (err *nestedError) Nested() error {
	return err.nested
}

func (err *nestedError) Deepest() error {
	if err.nested == nil {
		return nil
	}
	
	ne, ok := err.nested.(NestedError)

	if !ok {
		return err.nested
	}

	return ne.Deepest()
}

/*
	DeepestNested returns the deepest nested error. If err is not *NestedError,
	it is directly returned.
*/
func DeepestNested(err error) error {
	ne, ok := err.(NestedError)
	if ok {
		err = ne.Deepest()
	}

	return err
}

/*
  NestErrorf returns nil if err == nil, otherwise it returns a *NestedError
  error with a formatted message
*/
func NestErrorf(err error, fmtstr string, args ...interface{}) NestedError {
	if err == nil {
		return nil
	}
	return &nestedError{
		message: fmt.Sprintf(fmtstr, args...),
		nested:  err,
	}
}
