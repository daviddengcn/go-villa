package villa

import(
    "fmt"
    "bytes"
)

/*
IntMatrix is 2D array of integers. Elements are store in a single int slice and slices of each row are created.

NOTE the matrix can be sized of 0x0, but never 0x10 or 10x0.
*/
type IntMatrix [][]int


// NewIntMatrix creates a new IntMatrix instance with specified number of rows and columns
func NewIntMatrix(nRow, nCol int) IntMatrix {
    s := make([]int, nCol*nRow)
    mat := make(IntMatrix, nRow)
    for i, p := 0, 0; i < nRow; i ++ {
        mat[i] = s[p:p+nCol]
        p += nCol
    } // for i
    
    return mat
}

// Clone clones an IntMatrix
func (m IntMatrix) Clone() IntMatrix {
    mat := NewIntMatrix(m.Rows(), m.Cols())
    
    n := m.Rows() * m.Cols()
    if n > 0 {
        copy(mat[0][:n], m[0][:n])
    } // if
    
    return mat
}

// Cols returns the number of columns
func (m IntMatrix) Cols() int {
    if len(m) == 0 {
        return 0
    } // if
    
    return len(m[0])
}

// Rows returns the number of rows
func (m IntMatrix) Rows() int {
    return len(m)
}

// PrettyString returns a pretty text form of the matrix.
// This function is mainly for debugging.
func (m IntMatrix) PrettyString() string {
    sa := make([][]string, 0, m.Rows())
    for _, row := range(m) {
        sr := make([]string, 0, len(row))
        for _, cell := range(row) {
            sr = append(sr, fmt.Sprint(cell))
        } // for cell
        sa = append(sa, sr)
    } // for row
    
    wds := make([]int, m.Cols())
    for i := 0; i < m.Rows(); i ++ {
        for j := 0; j < m.Cols(); j ++ {
            if len(sa[i][j]) > wds[j] {
                wds[j] = len(sa[i][j])
            } //  if
        } // for j
    } // for i
    
    var res bytes.Buffer
    for i, row := range(sa) {
        if i == 0 {
            res.WriteString("[")
        } else {
            res.WriteString(" ")
        } // else
        res.WriteString("[")
        for j, cell := range(row) {
            if j > 0 {
                res.WriteString(" ")
            } // if
            fmt.Fprintf(&res, "%*s", wds[j], cell)
        } // for j, cell
        res.WriteString("]")
        if i == len(sa) - 1 {
            fmt.Fprintf(&res, "](%dx%d)", m.Rows(), m.Cols())
        } // else
        res.WriteString("\n")
    } // for row
    
    return res.String()
}

// Fill sets all elements of the matrix to a specified value
func (m IntMatrix) Fill(vl int) {
    if len(m) == 0 {
        return;
    } // if
    
    n := m.Rows() * m.Cols()
    IntSlice(m[0][:n]).Fill(0, n, vl)
}