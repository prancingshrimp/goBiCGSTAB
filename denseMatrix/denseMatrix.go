package denseMatrix

import (
    "fmt"
    "strconv"
    "strings"
)

type Matrix struct {
    Val [][]float64
    RowN, ColN, NoneZero int
}

func New() *Matrix {
    return &Matrix{}
}

func splitToFloats(s, sep string) []float64 {
    sepSave := 0
    n := strings.Count(s, sep) + 1
    c := sep[0]
    start := 0
    a := make([]float64, n)
    na := 0
    for i := 0; i+len(sep) <= len(s) && na+1 < n; i++ {
        if s[i] == c && (len(sep) == 1 || s[i:i+len(sep)] == sep) {
            a[na], _ = strconv.ParseFloat(s[start:i+sepSave], 64)
            na++
            start = i + len(sep)
            i += len(sep) - 1
        }
    }
    a[na], _ = strconv.ParseFloat(s[start:], 64)
    return a[0 : na+1]
}

func (self *Matrix) Set(matrix string) {
    matrix = strings.Replace(matrix, " ", "", -1)
    var startI, endI int
    for i, r := range matrix {
        if r == '[' {
            startI = i + 1
        }
        if r == ']' {
            endI = i
            self.Val = append(self.Val,
                splitToFloats(matrix[startI:endI], ","))
        }

    }
    for r := range self.Val {
        for c := range self.Val[r] {
            if self.Val[r][c] != 0 {
                self.NoneZero += 1
            }
        }
    }
    self.RowN = len(self.Val)
    self.ColN = len(self.Val[0])
}

func (self *Matrix) String() string {
    var printString string
    for _, item := range self.Val {
        printString += fmt.Sprintf("%+.2e \n", item)
    }
    printString = strings.Replace(printString,
        "+0.00e+00", "         ", -1)
    return strings.TrimRight(printString, "\n")
}
