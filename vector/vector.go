package vector

import (
    "fmt"
    "goBiCGSTAB/denseMatrix"
)

type Vector struct {
    Val []float64
    N       int
}

func New(n int) *Vector {
    return &Vector{Val: make([]float64, n), N: n}
}

// func Neww(n int) (...*Vector) {
//     return (...&Vector{Val: make([]float64, n), N: n})
// }


func NewFrom(v *Vector) *Vector {
    vecVal := make([]float64, v.N)
    for i := range v.Val {
        vecVal[i] = v.Val[i]
    }
    return &Vector{Val: vecVal, N: v.N}
}

func (self *Vector) Clear() {
    for i := range self.Val {
        self.Val[i] = 0
    }
}

func (self *Vector) FillFromDenseMatrix(matrix *denseMatrix.Matrix) {
    if matrix.RowN == 1 {
        for i := range matrix.Val[0] {
            self.Val[i] = matrix.Val[0][i]
        }
    }
    if matrix.ColN == 1 {
        for i := range matrix.Val {
            self.Val[i] = matrix.Val[i][0]
        }
    }
}

func (self *Vector) FillFromVector(v *Vector) {
    for i := range v.Val {
        self.Val[i] = v.Val[i]
    }
}

func (self *Vector) String() string {
    printString := "["
    for i := range self.Val {
        printString += fmt.Sprintf("%+.2e \n ", self.Val[i])
    }
    return printString[:len(printString)-3] + "]"
}
