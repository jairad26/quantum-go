/*
Quantum computing is based on qubits, quantum states, quantum gates, and quantum circuits.

## General Things to Know

### Notation
- |0⟩: [1,0] Represents the state first basis state of a single qubit. 
- |1⟩: [0,1] Represents the state second basis state of a single qubit.

For multiple qubits
- |00⟩: [1,0,0,0] Represents the state first basis state of two qubits.
- |01⟩: [0,1,0,0] Represents the state second basis state of two qubits.
- |10⟩: [0,0,1,0] Represents the state third basis state of two qubits.
- |11⟩: [0,0,0,1] Represents the state fourth basis state of two qubits.

Think of it as left to right binary representation of the state of the qubits.

### Bitwise operations in Go
- & : Bitwise AND e.g. 1010 & 1100 = 1000
- | : Bitwise OR e.g. 1010 | 1100 = 1110
- ^ : Bitwise XOR e.g. 1010 ^ 1100 = 0110
- << : Left shift e.g. 1010 << 1 = 10100
- >> : Right shift e.g. 1010 >> 1 = 101
- &^ : Bit clear (AND NOT) e.g. 1010 &^ 1100 = 0010

## Qubits
A qubit is the basic unit of quantum information. It is a two-level quantum system that can exist in a superposition of states. 
The state of a qubit can be represented as a linear combination of the basis states |0⟩ and |1⟩, denoted as α|0⟩ + β|1⟩, 
where α and β are complex numbers, used to represent the quantum state of the qubit.

## Quantum States
A quantum state is a vector in a complex vector space that represents the state of a quantum system.
The state of a quantum system can be represented as a superposition of basis states, where the coefficients of the 
basis states are complex numbers, representing the amplitude.

## Quantum Gates
Quantum gates are the building blocks of quantum circuits. They are unitary operators that act on qubits to perform operations.
Quantum gates can be used to manipulate the quantum state of qubits, entangle qubits, and perform quantum computations.

### Common Quantum Gates

#### Pauli X Gate
The Pauli X gate is a quantum gate that acts as a quantum NOT gate, flipping the state of a qubit from |0⟩ to |1⟩ and vice versa.

For multiple qubits, the Pauli X gate acts on a single qubit, leaving the other qubits unchanged.

for example:
- |00⟩ -> X on qubit 0 -> |01⟩
- |01⟩ -> X on qubit 0 -> |00⟩
- |10⟩ -> X on qubit 0 -> |11⟩
- |11⟩ -> X on qubit 0 -> |10⟩


#### Hadamard Gate
The Hadamard gate is a quantum gate that creates superposition by putting a qubit into an equal superposition of |0⟩ and |1⟩.

For multiple qubits, the Hadamard gate acts on a single qubit, leaving the other qubits unchanged.

for example:
- |0⟩ -> H -> (|0⟩ + |1⟩) / √2
- |1⟩ -> H -> (|0⟩ - |1⟩) / √2
- |00⟩ -> H on first qubit -> (|00⟩ + |10⟩) / √2
- |01⟩ -> H on first qubit -> (|01⟩ + |11⟩) / √2

## Quantum Circuits
A quantum circuit is a sequence of quantum gates applied to qubits to perform quantum computations.
*/