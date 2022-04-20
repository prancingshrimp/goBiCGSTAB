package vector

import (
    "math"
)

func (self *Vector) Norm2() float64 {
    var returnVal float64
    for i := range self.Val {
        returnVal += self.Val[i] * self.Val[i]
    }
    return math.Sqrt(returnVal)
}

// x = v1 * v2
func  Dot(v1, v2 *Vector) float64 {
    var returnVal float64
    for i := range v1.Val {
        returnVal += v1.Val[i] * v2.Val[i]
    }
    return returnVal
}

// x = v1 - s * v2
func  VmSV(v1 *Vector, scalar float64, v2, x *Vector) {
    for i := range v1.Val {
        x.Val[i] = v1.Val[i] - v2.Val[i] * scalar
    }
}

// x = x + s1 * v1 + s2 * v2
func  VpSVpSV(x *Vector, s1 float64, v1 *Vector, s2 float64, v2 *Vector) {
    for i := range x.Val {
        x.Val[i] += s1 * v1.Val[i] + s2 * v2.Val[i]
    }
}

// x = v1 + s1 * (x - s2 * v2)
func  VpS_VmSV(v1 *Vector, s1, s2 float64, v2,x *Vector) {
    for i := range x.Val {
        x.Val[i] = v1.Val[i] + s1 * (x.Val[i] - s2 * v2.Val[i])
    }
}
