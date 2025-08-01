package main

import "fmt"

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]

}

func main() {

	ns := transition(StateIdle)
	fmt.Println(ns)
}

func transition(s ServerState) ServerState {

	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:

		return StateIdle
	case StateError:
		return StateRetrying
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
