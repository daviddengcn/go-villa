package villa

import(
    "testing"
    "fmt"
)

func TestIntMatrix(t *testing.T) {
    defer __(o_())

    mat := NewIntMatrix(5, 4)
    AssertEquals(t, "mat.Rows()", mat.Rows(), 5)
    AssertEquals(t, "mat.Cols()", mat.Cols(), 4)
    AssertStringEquals(t, "mat", mat, "[[0 0 0 0] [0 0 0 0] [0 0 0 0] [0 0 0 0] [0 0 0 0]]")
    
    mat.Fill(10)
    fmt.Println(mat.PrettyString())
    AssertStringEquals(t, "mat", mat, "[[10 10 10 10] [10 10 10 10] [10 10 10 10] [10 10 10 10] [10 10 10 10]]")
    
    // Clone for null matrix, assure no panic
    mat = IntMatrix(nil)
    mat.Clone()
    mat.Fill(0)
    AssertEquals(t, "mat.Rows()", mat.Rows(), 0)
    AssertEquals(t, "mat.Cols()", mat.Cols(), 0)
}
