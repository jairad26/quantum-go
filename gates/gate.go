package gates

import (
	"math"
	"quantum-go/qstate"
)

func X(q *qstate.QState) {
	q.Alpha, q.Beta = q.Beta, q.Alpha
}

func H(q *qstate.QState) {
	newAlpha := (q.Alpha + q.Beta) / complex(math.Sqrt(2), 0)
	newBeta := (q.Alpha - q.Beta) / complex(math.Sqrt(2), 0)
	q.Alpha, q.Beta = newAlpha, newBeta
}

func CNOT(control, target *qstate.QState) {
	if control.Prob1() == 1 {
		X(target)
	}
}
