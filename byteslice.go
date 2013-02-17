package villa

import (
	"io"
)

// ByteSlice is a wrapper type for []byte.
// It satisfies io.Reader and io.Writer interfaces.
type ByteSlice []byte

// Read is the read function of io.Reader.
// After some bytes are read, the slice shrinks
func (s *ByteSlice) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}

	if len(p) == 0 {
		return 0, io.EOF
	}
	n = copy(p, *s)

	if n == len(*s) {
		*s = nil
	} else {
		*s = (*s)[n:]
	}

	return n, nil
}

// Write is the write function of io.Writer.
// Bytes are appended to the tail of the slice.
func (s *ByteSlice) Write(p ...byte) (n int, err error) {
	*s = append(*s, p...)
	return len(p), nil
}
