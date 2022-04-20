package output

import (
    "fmt"
)

type BitFlag int

const (
    Solver BitFlag = 1
    Test   BitFlag = 2
    Misc   BitFlag = 3
)

type MessageChan struct {
    Content []string
}

func New(msgType BitFlag) chan MessageChan {
    msg := make(chan MessageChan)
    switch {
    case msgType == Solver:
        go solverOutput(msg)
        return msg

    case msgType == Misc:
        go miscOutput(msg)
        return msg

    default:
        panic("Unknown flag type")
    }
}

func solverOutput(msg chan MessageChan) {
    var out MessageChan
    for {
        out = <-msg
        fmt.Printf("%s: Solving %s, InitRes = %s, FinalRes = %s, NoIter = %s\n",
            out.Content[0], out.Content[1], out.Content[2], out.Content[3], out.Content[4])
    }
}

func miscOutput(msg chan MessageChan) {
    fmt.Println("Misc")
    for {
        fmt.Println(<-msg)
    }
}
