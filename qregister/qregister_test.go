package qregister_test

import (
	"math"
	"quantum-go/qregister"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	qr := qregister.New(2)

	require.NotNil(t, qr)
	require.Equal(t, 2, qr.GetNumQubits())
	require.Equal(t, 4, len(qr.GetState()))
	require.Equal(t, complex(1, 0), qr.GetState()[0])
}

func TestH(t *testing.T) {
	qr := qregister.New(1)
	qr.H(0)

	expected := 1.0 / math.Sqrt(2)

	require.Equal(t, complex(expected, 0), qr.GetState()[0])
	require.Equal(t, complex(expected, 0), qr.GetState()[1])
}

func TestX(t *testing.T) {
	qr := qregister.New(1)
	qr.X(0)

	require.Equal(t, complex(0, 0), qr.GetState()[0])
	require.Equal(t, complex(1, 0), qr.GetState()[1])
}

func TestCNOT(t *testing.T) {
	tests := []struct {
		name          string
		setupFunc     func(*qregister.QRegister)
		expectedState []complex128
		control       int
		target        int
	}{
		{
			name:          "CNOT on |00⟩",
			setupFunc:     func(qr *qregister.QRegister) {},
			expectedState: []complex128{1, 0, 0, 0},
			control:       0,
			target:        1,
		},
		{
			name: "CNOT on |10⟩",
			setupFunc: func(qr *qregister.QRegister) {
				qr.X(0)
			},
			expectedState: []complex128{0, 0, 0, 1},
			control:       0,
			target:        1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qr := qregister.New(2)
			tt.setupFunc(qr)
			qr.CNOT(tt.control, tt.target)

			require.Equal(t, tt.expectedState, qr.GetState())
		})
	}
}

func TestMeasure(t *testing.T) {
	qr := qregister.New(2)
	qr.H(0) // Create superposition of first qubit

	result := qr.Measure()

	require.Len(t, result, 2)

	// After measurement, state should be collapsed
	numNonZero := 0
	for _, amp := range qr.GetState() {
		if amp != 0 {
			numNonZero++
		}
	}
	require.Equal(t, 1, numNonZero)
}
