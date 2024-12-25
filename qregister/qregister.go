package qregister

import (
	"quantum-go/gates"
	"quantum-go/qubit"
)

// QRegister represents a quantum register, a collection of qubits.
type QRegister struct {
	qubits []*qubit.Qubit
}

// New creates a new quantum register with the specified number of qubits.
func New(numQubits int) *QRegister {
	qubits := make([]*qubit.Qubit, numQubits)
	for i := range qubits {
		qubits[i] = qubit.NewQubit()
	}
	return &QRegister{qubits: qubits}
}

// ApplyH applies the Hadamard gate to the qubit at the specified index.
func (qr *QRegister) ApplyH(index int) {
	qr.qubits[index].ApplyH()
}

// ApplyX applies the Pauli-X gate to the qubit at the specified index.
func (qr *QRegister) ApplyX(index int) {
	qr.qubits[index].ApplyX()
}

// ApplyCNOT applies the CNOT gate to the qubits at the specified control and target indices.
func (qr *QRegister) ApplyCNOT(controlIndex, targetIndex int) {
	gates.CNOT(qr.qubits[controlIndex].State(), qr.qubits[targetIndex].State())
}

// MeasureQubit measures the qubit at the specified index.
func (qr *QRegister) MeasureQubit(index int) bool {
	return qr.qubits[index].Measure()
}

// Measure measures all qubits in the register.
func (qr *QRegister) Measure() []bool {
	results := make([]bool, len(qr.qubits))
	for i := range qr.qubits {
		results[i] = qr.qubits[i].Measure()
	}
	return results
}
