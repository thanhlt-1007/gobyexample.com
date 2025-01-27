package main

import (
    "fmt"
)

type State int

const (
    IdleState State = 1
    ConnectedState State = 2
    ErrorState State = 3
    ReryState State = 4
)

var StateName = map[State]string {
    IdleState: "idle",
    ConnectedState: "connected",
    ErrorState: "error",
    ReryState: "retry",
}

func (state State) String() string {
    return StateName[state]
}

func nextState(state State) State {
    switch state {
    case IdleState:
        return ConnectedState
    case ConnectedState:
        panic("ConnectedState is final state")
    case ErrorState:
        return ReryState
    default:
        return IdleState
    }
}

func main() {
    fmt.Println(IdleState)
    // fmt.Println(nextState(ConnectedState))
    fmt.Println(nextState(IdleState))
}
