package villa

import (
	"testing"
)

func TestByteSlice(t *testing.T) {
	var bs ByteSlice
	AssertEquals(t, "len(bs)", len(bs), 0)
	AssertStringEquals(t, "bs", bs, "[]")

	bs.Write(1, 2, 3)
	AssertEquals(t, "len(bs)", len(bs), 3)
	AssertStringEquals(t, "bs", bs, "[1 2 3]")

	bs.Read(make([]byte, 2))
	AssertEquals(t, "len(bs)", len(bs), 1)
	AssertStringEquals(t, "bs", bs, "[3]")

	bs.Read(make([]byte, 1))
	AssertEquals(t, "len(bs)", len(bs), 0)
	AssertStringEquals(t, "bs", bs, "[]")

	bs.Write(4, 5)
	AssertEquals(t, "len(bs)", len(bs), 2)
	AssertStringEquals(t, "bs", bs, "[4 5]")
}
