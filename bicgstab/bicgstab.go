//testdoc
package bicgstab

import (
    // "fmt"
    "goBiCGSTAB/output"
    sp "goBiCGSTAB/sparseMatrix"
    "strconv"
    vp "goBiCGSTAB/vector"
)

type Options struct {
    Name    string
    Tol     float64
    MaxIter int
    MsgChan chan output.MessageChan
}

func setOptions(opt *Options) {
    if opt.Name == "" {
        opt.Name = "Unknown"
    }
    if opt.Tol == 0 {
        opt.Tol = 10.0e-6
    }
    if opt.MaxIter == 0 {
        opt.MaxIter = 1000
    }
    if opt.MsgChan == nil {
        opt.MsgChan = make(chan output.MessageChan)
        go func(trash chan output.MessageChan) {
            for {
                <-trash
            }
        }(opt.MsgChan)
    }
}

func Run(A *sp.Matrix, x0, b *vp.Vector, opt *Options) chan bool {

    switchChan := make(chan bool)

    go func(switchChan chan bool) {
        setOptions(opt)

        var normr, tolb float64
        var rr0, r1r0, vr0 float64
        var ts, tt float64
        var alpha, beta, omega float64
        // var normb float64

        r0 := vp.New(b.N)
        r := vp.New(b.N)
        p := vp.New(b.N)
        v := vp.New(b.N)
        s := vp.New(b.N)
        t := vp.New(b.N)

        solverMsg := make([]string, 5)
        solverMsg[0] = "BiCGStab"
        solverMsg[1] = opt.Name

        for {
            sp.VmMV(b, A, x0, r0)
            r.FillFromVector(r0)
            p.FillFromVector(r0)
            rr0 = vp.Dot(r, r0)

            // normb = b.Norm2()
            tolb = opt.Tol// * normb

            solverMsg[2] = strconv.FormatFloat(r.Norm2(), 'G', 8, 64)

            for i := 0; ; i++ {
                normr = r.Norm2()
                if normr <= tolb || i == opt.MaxIter {
                    solverMsg[3] = strconv.FormatFloat(normr, 'G', 8, 64)
                    solverMsg[4] = strconv.Itoa(i)
                    opt.MsgChan <- output.MessageChan{solverMsg}
                    break
                }

                sp.MV(A, p, v)
                vr0 = vp.Dot(v, r0)
                // if vr0 == 0 {
                //     fmt.Println("Bicgstab break-down")
                // }
                // if rr0 == 0 {
                //     fmt.Println("Bicgstab Loesung stagniert")
                // }
                alpha = rr0 / vr0
                vp.VmSV(r, alpha, v, s)
                sp.MV(A, s, t)

                ts = vp.Dot(s, t)
                tt = vp.Dot(t, t)
                // if tt == 0 || ts == 0 {
                //     fmt.Println("Bicgstab break-down")
                // }

                omega = ts / tt

                // x0 = x0 + alpha * p + omega * s
                vp.VpSVpSV(x0, alpha, p, omega, s)
                // r = s - omega * t
                vp.VmSV(s, omega, t, r)

                r1r0 = vp.Dot(r, r0)

                beta = (alpha * r1r0) / (omega * rr0)
                // p = r + beta * (p - omega * v)
                vp.VpS_VmSV(r, beta, omega, v, p)
                rr0 = r1r0
            }
            switchChan <- true
            <-switchChan
        }
    }(switchChan)
    return switchChan
}
