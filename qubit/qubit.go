package qubit

import (
	"quantum-go/gates"
	"quantum-go/qstate"
)

type Qubit struct {
	state *qstate.QState
}

func NewQubit() *Qubit {
	state, _ := qstate.New(1, 0) // Initialize to |0⟩
	return &Qubit{state: state}
}

func (q *Qubit) ApplyX() {
	gates.X(q.state)
}

func (q *Qubit) ApplyH() {
	gates.H(q.state)
}

func (q *Qubit) Measure() bool {
	return q.state.Measure()
}

func (q *Qubit) Reset() {
	q.state.Alpha = 1
	q.state.Beta = 0
}

func (q *Qubit) State() *qstate.QState {
	return q.state
}
