package quantumgo_test

import (
	"quantum-go/qregister"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBellState(t *testing.T) {
	// Test initial superposition
	t.Run("Initial State", func(t *testing.T) {
		qr := qregister.New(2)
		qr.H(0)
		qr.CNOT(0, 1)

		probs := qr.Probabilities()
		tolerance := 1e-10

		require.InDelta(t, 0.5, probs[0], tolerance)
		require.InDelta(t, 0, probs[1], tolerance)
		require.InDelta(t, 0, probs[2], tolerance)
		require.InDelta(t, 0.5, probs[3], tolerance)
	})

	// Test entanglement behavior
	t.Run("Entanglement", func(t *testing.T) {
		trials := 1000
		matchingResults := 0

		for i := 0; i < trials; i++ {
			qr := qregister.New(2)
			qr.H(0)
			qr.CNOT(0, 1)

			result := qr.Measure()
			// Verify both qubits always match (both 0 or both 1)
			require.Equal(t, result[0], result[1],
				"Qubits should be perfectly correlated")
			if result[0] == 1 {
				matchingResults++
			}
		}

		// Check distribution is roughly 50/50
		ratio := float64(matchingResults) / float64(trials)
		require.InDelta(t, 0.5, ratio, 0.1,
			"Should get roughly equal 00 and 11 measurements")
	})
}
