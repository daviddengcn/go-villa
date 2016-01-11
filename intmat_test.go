package villa

import (
	"testing"

	"github.com/golangplus/testing/assert"
)

func TestIntMatrix(t *testing.T) {
	defer __(o_(t))

	mat := NewIntMatrix(5, 4)
	assert.Equal(t, "mat.Rows()", mat.Rows(), 5)
	assert.Equal(t, "mat.Cols()", mat.Cols(), 4)
	assert.StringEqual(t, "mat", mat, "[[0 0 0 0] [0 0 0 0] [0 0 0 0] [0 0 0 0] [0 0 0 0]]")

	mat.Fill(10)
	assert.StringEqual(t, "mat", mat, "[[10 10 10 10] [10 10 10 10] [10 10 10 10] [10 10 10 10] [10 10 10 10]]")

	mat[1][1] = 0
	mat[3][2] = 12345
	mat[2][0] = -998
	t.Logf("%s", mat.PrettyString())
}

func TestIntMatrix_Clone(t *testing.T) {
	// Clone for null matrix, assure no panic
	mat := IntMatrix(nil)
	mat.Clone()
	mat.Fill(0)
	assert.Equal(t, "mat.Rows()", mat.Rows(), 0)
	assert.Equal(t, "mat.Cols()", mat.Cols(), 0)

	// Clone for non-null matrix
	mat = NewIntMatrix(2, 2)
	mat.Fill(1)
	matB := mat.Clone()
	// Change contents of mat, matB should not be changed.
	mat.Fill(2)
	assert.StringEqual(t, "mat", mat, "[[2 2] [2 2]]")
	assert.StringEqual(t, "matB", matB, "[[1 1] [1 1]]")
}
