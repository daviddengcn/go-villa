package villa

import "fmt"

/*
IntMatrix is 2D array of integers. Elements are store in a single int slice and slices of each row are created.
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

// PrettyString returns a pretty string form of the matrix
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
    for j := 0; j < m.Cols(); j ++ {
        for i := 0; i < m.Rows(); i ++ {
            if len(sa[i][j]) > wds[j] {
                wds[j] = len(sa[i][j])
            } //  if
        } // for i
    } // for i
    
    res := ""
    for i, row := range(sa) {
        if i == 0 {
            res += "["
        } else {
            res += " "
        } // else
        res += "["
        for j, cell := range(row) {
            if j > 0 {
                res += " "
            } // if
            res += fmt.Sprintf(fmt.Sprintf("%%%ds", wds[j]), cell)
        } // for j, cell
        res += "]"
        if i == len(sa) - 1 {
            res += fmt.Sprintf("](%dx%d)", m.Rows(), m.Cols())
        } // else
        res += "\n"
    } // for row
    
    return res
}


func (m IntMatrix) Fill(vl int) {
    n := m.Rows() * m.Cols()
    IntSlice(m[0][:n]).Fill(0, n, vl)
}