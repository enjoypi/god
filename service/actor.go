package service

import (
	sc "github.com/enjoypi/gostatechart"
)

type Actor struct {
	*sc.StateMachine
}

func (a *Actor) Run() {

}
