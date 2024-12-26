package qregister

import (
	"math"
	"math/cmplx"
	"math/rand/v2"
)

// QRegister represents a quantum register, a collection of qubits.
type QRegister struct {
	numQubits int
	state     []complex128
}

// New creates a new quantum register with the specified number of qubits.
func New(numQubits int) *QRegister {
	size := 1 << numQubits
	state := make([]complex128, size)
	state[0] = 1
	return &QRegister{
		numQubits: numQubits,
		state:     state,
	}
}

// GetState returns a copy of the current quantum state.
func (qr *QRegister) GetState() []complex128 {
	stateCopy := make([]complex128, len(qr.state))
	copy(stateCopy, qr.state)
	return stateCopy
}

// GetNumQubits returns  the number of qubits in the quantum register.
func (qr *QRegister) GetNumQubits() int {
	return qr.numQubits
}

// ApplyH applies the Hadamard gate to the qubit at the specified index.
/*
Explanation:
The Hadamard gate is a single-qubit gate that maps the basis states |0⟩ and |1⟩ to superposition states.
It is defined by the following matrix:
H = 1/√2 * [[1, 1], [1, -1]]
The Hadamard gate can be represented as a unitary matrix acting on a single qubit.

Example:
Consider a quantum register with 2 qubits. The Hadamard gate can be applied to the first qubit as follows:
Let's say index = 1 (second qubit)
mask = 1 << 1 = binary 10

For i = 0 (binary 00):
bit = (00 & 10) >> 1 = 0
pair = 00 ^ 10 = 10 (decimal 2)

For i = 1 (binary 01):
bit = (01 & 10) >> 1 = 1
pair = 01 ^ 10 = 11 (decimal 3)

For i = 2 (binary 10):
bit = (10 & 10) >> 1 = 1
pair = 10 ^ 10 = 00 (decimal 0)

For i = 3 (binary 11):
bit = (11 & 10) >> 1 = 1
pair = 11 ^ 10 = 01 (decimal 1)
*/
func (qr *QRegister) H(index int) *QRegister {
	if index >= qr.numQubits {
		return qr
	}

	factor := complex(1.0/math.Sqrt(2), 0)

	// shift 1 left by index bits
	mask := 1 << index
	newState := make([]complex128, len(qr.state))

	for i := range qr.state {
		bit := (i & mask) >> index
		pair := i ^ mask // Find the corresponding pair state

		if bit == 0 {
			// |0⟩ -> (|0⟩ + |1⟩)/√2
			newState[i] += factor * qr.state[i]
			newState[i] += factor * qr.state[pair]
		} else {
			// |1⟩ -> (|0⟩ - |1⟩)/√2
			newState[i] += factor * qr.state[pair]
			newState[i] -= factor * qr.state[i]
		}
	}

	qr.state = newState
	return qr
}

// ApplyX applies the Pauli-X gate to the qubit at the specified index.
/*
Explanation:
The Pauli-X gate is a single-qubit gate that flips the state of a qubit.
It is defined by the following matrix:
X = [[0, 1], [1, 0]]
The Pauli-X gate can be represented as a unitary matrix acting on a single qubit.

Example:
Consider a quantum register with 2 qubits. The Pauli-X gate can be applied to the first qubit as follows:
Let's say index = 1 (second qubit)
mask = 1 << 1 = binary 10

For i = 0 (binary 00):
flipped = 00 ^ 10 = 10 (decimal 2)

For i = 1 (binary 01):
flipped = 01 ^ 10 = 11 (decimal 3)

For i = 2 (binary 10):
flipped = 10 ^ 10 = 00 (decimal 0)

For i = 3 (binary 11):
flipped = 11 ^ 10 = 01 (decimal 1)
*/
func (qr *QRegister) X(index int) *QRegister {
	if index >= qr.numQubits {
		return qr
	}

	// shift 1 left by index bits
	mask := 1 << index
	newState := make([]complex128, len(qr.state))

	for i := range qr.state {
		// Flip the bit at the specified index
		flipped := i ^ mask
		newState[flipped] = qr.state[i]
	}

	qr.state = newState
	return qr
}

// ApplyCNOT applies the CNOT gate to the qubits at the specified control and target indices.

/*
Explanation:
The CNOT gate is a two-qubit gate that flips the target qubit if the control qubit is in the |1⟩ state.
It is defined by the following matrix:
CNOT = [[1, 0, 0, 0], [0, 1, 0, 0], [0, 0, 0, 1], [0, 0, 1, 0]]
The CNOT gate can be represented as a unitary matrix acting on two qubits.

Example:
Consider a quantum register with 2 qubits. The CNOT gate can be applied to the first qubit as the control qubit
and the second qubit as the target qubit as follows:
Let's say controlIndex = 0 (first qubit) and targetIndex = 1 (second qubit)
controlMask = 1 << 0 = binary 1
targetMask = 1 << 1 = binary 10

For i = 0 (binary 00):
controlBit = (00 & 01) >> 0 = 0
newState[0] = state[0]

For i = 1 (binary 01):
controlBit = (01 & 01) >> 0 = 1
flipped = 01 ^ 10 = 11 (decimal 3)

For i = 2 (binary 10):
controlBit = (10 & 01) >> 0 = 0
newState[2] = state[2]

For i = 3 (binary 11):
controlBit = (11 & 01) >> 0 = 1
flipped = 11 ^ 10 = 01 (decimal 1)

The CNOT gate flips the target qubit if the control qubit is in the |1⟩ state.
*/
func (qr *QRegister) CNOT(controlIndex, targetIndex int) *QRegister {
	if controlIndex >= qr.numQubits || targetIndex >= qr.numQubits {
		return qr
	}

	controlMask := 1 << controlIndex
	targetMask := 1 << targetIndex
	newState := make([]complex128, len(qr.state))

	for i := range qr.state {
		controlBit := (i & controlMask) >> controlIndex

		if controlBit == 1 {
			// If control qubit is |1⟩, flip the target qubit
			flipped := i ^ targetMask
			newState[flipped] = qr.state[i]
		} else {
			// If control qubit is |0⟩, leave the state unchanged
			newState[i] = qr.state[i]
		}
	}

	qr.state = newState
	return qr
}

// Measure collapses the quantum register into a classical state by measuring each qubit.
func (qr *QRegister) Measure() []int {
	// Calculate probabilities for each basis state
	probs := make([]float64, len(qr.state))
	for i, amp := range qr.state {
		probs[i] = real(amp * cmplx.Conj(amp))
	}

	// Generate random number between 0 and 1
	r := rand.Float64()

	// Find which state we collapsed to
	var cumulative float64
	var outcome int
	for i, prob := range probs {
		cumulative += prob
		if r <= cumulative {
			outcome = i
			break
		}
	}

	// Convert outcome to binary representation
	result := make([]int, qr.numQubits)
	for i := 0; i < qr.numQubits; i++ {
		result[i] = (outcome >> i) & 1
	}

	// Collapse state vector
	newState := make([]complex128, len(qr.state))
	newState[outcome] = 1
	qr.state = newState

	return result
}

func (qr *QRegister) Probabilities() []float64 {
	probs := make([]float64, len(qr.state))

	for i, amp := range qr.state {
		// Probability = |amplitude|^2
		probs[i] = real(amp * cmplx.Conj(amp))
	}

	return probs
}
