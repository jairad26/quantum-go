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

func (q *Qubit) H() *Qubit {
	q.register.H(q.position)
	return q
}

func (q *Qubit) X() *Qubit {
	q.register.X(q.position)
	return q
}

func (q *Qubit) CNOT(target int) *Qubit {
	q.register.CNOT(q.position, target)
	return q
}
