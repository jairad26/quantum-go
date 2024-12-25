package qstate

import (
	"errors"
	"math/cmplx"
	"math/rand/v2"
)

var ErrInvalidQState = errors.New("invalid quantum state, the sum of the squares of the amplitudes must be 1")

type QState struct {
	Alpha complex128
	Beta  complex128
}

func New(alpha, beta complex128) (*QState, error) {
	if !validate(alpha, beta) {
		return nil, ErrInvalidQState
	}
	return &QState{
		Alpha: alpha,
		Beta:  beta,
	}, nil
}

func validate(alpha, beta complex128) bool {
	return cmplx.Abs(alpha)*cmplx.Abs(alpha)+cmplx.Abs(beta)*cmplx.Abs(beta) == 1
}

// For a quantum state |ψ⟩ = α|0⟩ + β|1⟩, normalization formula is: √(|α|² + |β|²) = 1
// α_normalized = α/√(|α|² + |β|²)
// β_normalized = β/√(|α|² + |β|²)
func (q *QState) Normalize() {
	sum := cmplx.Abs(q.Alpha)*cmplx.Abs(q.Alpha) + cmplx.Abs(q.Beta)*cmplx.Abs(q.Beta)
	norm := 1 / cmplx.Sqrt(complex(sum, 0))
	q.Alpha *= norm
	q.Beta *= norm
}

func (q *QState) Prob0() float64 {
	return cmplx.Abs(q.Alpha) * cmplx.Abs(q.Alpha)
}

func (q *QState) Prob1() float64 {
	return cmplx.Abs(q.Beta) * cmplx.Abs(q.Beta)
}

func (q *QState) Measure() bool {
	if rand.Float64() < q.Prob0() {
		q.Alpha = 1
		q.Beta = 0
		return false
	}
	q.Alpha = 0
	q.Beta = 1
	return true
}
