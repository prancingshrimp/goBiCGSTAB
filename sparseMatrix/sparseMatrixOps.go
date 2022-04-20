package sparseMatrix

import (
    vp "goBiCGSTAB/vector"
)

// x = v1 - m * v2
func VmMV(v1 *vp.Vector, m *Matrix, v2, x *vp.Vector) {
    var r, row int
    var sum float64
    for c := range m.Val {
        // new row: reset counter and sum
        if m.Row[c] != row {
            x.Val[row] = v1.Val[row] - sum
            row += 1
            sum = 0
            r = m.Col[c]
        }
        sum += m.Val[c] * v2.Val[r]
        r += 1
    }
    // last row needs special treatment
    x.Val[row] = v1.Val[row] - sum
}

// x = m * v
func MV(m *Matrix, v, x *vp.Vector) {
    var r, row int
    var sum float64
    for c := range m.Val {
        // new row: reset counter and sum
        if m.Row[c] != row {
            x.Val[row] = sum
            row += 1
            sum = 0
            r = m.Col[c]
        }
        sum += m.Val[c] * v.Val[r]
        r += 1
    }
    // last row needs special treatment
    x.Val[row] = sum
}
