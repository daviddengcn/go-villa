package villa

import (
	"testing"
)

func TestIntMatrix(t *testing.T) {
	defer __(o_(t))

	mat := NewIntMatrix(5, 4)
	AssertEquals(t, "mat.Rows()", mat.Rows(), 5)
	AssertEquals(t, "mat.Cols()", mat.Cols(), 4)
	AssertStringEquals(t, "mat", mat, "[[0 0 0 0] [0 0 0 0] [0 0 0 0] [0 0 0 0] [0 0 0 0]]")

	mat.Fill(10)
	AssertStringEquals(t, "mat", mat, "[[10 10 10 10] [10 10 10 10] [10 10 10 10] [10 10 10 10] [10 10 10 10]]")

	mat[1][1] = 0
	mat[3][2] = 12345
	mat[2][0] = -998
	t.Logf("%s", mat.PrettyString())

	// Clone for null matrix, assure no panic
	mat = IntMatrix(nil)
	mat.Clone()
	mat.Fill(0)
	AssertEquals(t, "mat.Rows()", mat.Rows(), 0)
	AssertEquals(t, "mat.Cols()", mat.Cols(), 0)
}
