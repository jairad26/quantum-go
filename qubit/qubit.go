package qubit

import (
	"quantum-go/qregister"
)

type Qubit struct {
	register *qregister.QRegister
	position int
}

func New(register *qregister.QRegister, position int) *Qubit {
	return &Qubit{
		register: register,
		position: position,
	}
}

func (q *Qubit) H() {
	q.register.H(q.position)
}

func (q *Qubit) X() {
	q.register.X(q.position)
}

func (q *Qubit) CNOT(target int) {
	q.register.CNOT(q.position, target)
}
