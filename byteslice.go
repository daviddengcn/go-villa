package villa

import (
	"errors"
	"io"
	"unicode/utf8"
)

// ByteSlice is a wrapper type for []byte.
// Its pointer form, *ByteSlice, satisfies io.Reader, io.Writer, io.ByteReader,
// io.Closer, io.ReaderFrom, io.WriterTo and io.RuneReader interfaces.
type ByteSlice []byte

// Read implements io.Reader interface.
// After some bytes are read, the slice shrinks.
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

// Write implements  io.Writer interface
// Bytes are appended to the tail of the slice.
func (s *ByteSlice) Write(p []byte) (n int, err error) {
	*s = append(*s, p...)
	return len(p), nil
}

// ReadByte implements io.ByteReader interface
func (s *ByteSlice) ReadByte() (c byte, err error) {
	if len(*s) < 1 {
		return 0, io.EOF
	}

	c = (*s)[0]
	*s = (*s)[1:]
	return c, nil
}

// Close implements io.Closer interface.
// It does nothing.
func (s ByteSlice) Close() error {
	return nil
}

// ReadFrom implements io.ReaderFrom interface.
func (s *ByteSlice) ReadFrom(r io.Reader) (n int64, err error) {
	const buf_SIZE = 32 * 1024
	buf := make([]byte, buf_SIZE)
	for {
		nRead, err := r.Read(buf)
		if nRead == 0 {
			if err == io.EOF {
				return n, nil
			}
			break
		}
		n += int64(nRead)
		*s = append(*s, buf[:nRead]...)
		if err == io.EOF {
			return n, nil
		}

		if err != nil {
			break
		}
	}

	return n, err
}

// WriteTo implements io.WriterTo interface.
func (s ByteSlice) WriteTo(w io.Writer) (n int64, err error) {
	nWrite, err := w.Write(s)
	return int64(nWrite), err
}

// ReadRune implements io.RuneReader interface.
func (s *ByteSlice) ReadRune() (r rune, size int, err error) {
	if !utf8.FullRune(*s) {
		return utf8.RuneError, 0, io.ErrUnexpectedEOF
	}
	r, size = utf8.DecodeRune(*s)
	if r != utf8.RuneError {
		*s = (*s)[size:]
	}

	return r, size, err
}

var ErrInvalidRune = errors.New("villa.ByteSlice: invalid rune")

// WriteRune writes a single Unicode code point, returning the number of bytes
// written and any error.
func (s *ByteSlice) WriteRune(r rune) (size int, err error) {
	if r < utf8.RuneSelf {
		*s = append(*s, byte(r))
		return 1, nil
	}

	l := utf8.RuneLen(r)
	if l < 0 {
		return 0, ErrInvalidRune
	}

	*s = append(*s, make([]byte, l)...)
	utf8.EncodeRune((*s)[len(*s)-l:], r)
	return l, nil
}

// WriteString appends the contents of str to the slice, growing the slice as
// needed. The return value n is the length of str; err is always nil.
func (s *ByteSlice) WriteString(str string) (size int, err error) {
	*s = append(*s, str...)
	return len(str), nil
}
