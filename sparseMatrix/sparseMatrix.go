package sparseMatrix

import (
    "fmt"
    "strings"
    "goBiCGSTAB/denseMatrix"
)

type Matrix struct {
    RowN, ColN int
    Row, Col   []int
    Val        []float64
}

func New(n int) *Matrix {
    return &Matrix{Val: make([]float64, n), Row: make([]int, n),
        Col: make([]int, n)}
}

func (self *Matrix) FillFromDenseMatrix(matrix *denseMatrix.Matrix) {
    var i int
    for r := range matrix.Val {
        for c := range matrix.Val[r] {
            if matrix.Val[r][c] != 0.0 {
                self.Row[i] = r
                self.Col[i] = c
                self.Val[i] = matrix.Val[r][c]
                if self.Row[i] > self.RowN {
                    self.RowN = self.Row[i]
                }
                if self.Col[i] > self.ColN {
                    self.ColN = self.Col[i]
                }
                i += 1
            }
        }
    }
}

func (self *Matrix) String() string {
    printString := "[ "
    r := 0
    prevCol := 0
    for i := range self.Val {
        if self.Row[i] > r {
            r += 1
            prevCol = 0
            printString += fmt.Sprintf("\n  ")
        }
        if self.Col[i]-prevCol > 0 {
            printString += strings.Repeat("          ", self.Col[i]-prevCol)
            prevCol = self.Col[i]
        }
        printString += fmt.Sprintf("%+.2e ", self.Val[i])
        prevCol += 1
    }
    return printString + "]"
}
