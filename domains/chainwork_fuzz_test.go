package domains

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// FuzzCompactToBig tests the CompactToBig function with random uint32 inputs.
// CompactToBig converts Bitcoin's compact difficulty representation to a big.Int.
// The function should never panic regardless of input.
func FuzzCompactToBig(f *testing.F) {
	// Seed corpus with known Bitcoin difficulty values from real blocks
	f.Add(uint32(0x1d00ffff)) // Genesis block difficulty
	f.Add(uint32(0x1b04864c)) // Block 100,000
	f.Add(uint32(0x1a05db8b)) // Block 200,000
	f.Add(uint32(0x1900896c)) // Block 300,000
	f.Add(uint32(0x1806b99f)) // Block 400,000
	f.Add(uint32(0x1809b91a)) // Block 500,000

	// Edge cases
	f.Add(uint32(0))          // Zero
	f.Add(uint32(1))          // Minimum non-zero
	f.Add(uint32(0xffffffff)) // Maximum uint32
	f.Add(uint32(0x00800000)) // Sign bit edge case
	f.Add(uint32(0x007fffff)) // Max mantissa, zero exponent
	f.Add(uint32(0x01000001)) // Exponent 1
	f.Add(uint32(0x02000001)) // Exponent 2
	f.Add(uint32(0x03000001)) // Exponent 3 (boundary)
	f.Add(uint32(0x04000001)) // Exponent 4 (above boundary)

	f.Fuzz(func(t *testing.T, compact uint32) {
		// CompactToBig should never panic
		result := CompactToBig(compact)

		// Result must always be a valid big.Int (not nil)
		require.NotNil(t, result, "CompactToBig returned nil for input 0x%08x", compact)

		// Extract components for validation
		mantissa := compact & 0x007fffff
		isNegative := compact&0x00800000 != 0
		exponent := uint(compact >> 24)

		// If mantissa is zero, result should be zero regardless of exponent
		if mantissa == 0 {
			require.Equal(t, int64(0), result.Int64(),
				"Expected zero result for zero mantissa, input 0x%08x", compact)
		}

		// If negative flag is set and mantissa is non-zero, result should be negative
		if isNegative && mantissa != 0 {
			require.LessOrEqual(t, result.Sign(), 0,
				"Expected non-positive result for negative flag, input 0x%08x", compact)
		}

		// For non-negative values with non-zero mantissa, result should be positive
		if !isNegative && mantissa != 0 {
			require.GreaterOrEqual(t, result.Sign(), 0,
				"Expected non-negative result, input 0x%08x", compact)
		}

		// Verify the result is finite (not an uninitialized pointer)
		_ = result.BitLen()

		// For very small exponents (<=3), verify mantissa shift is correct
		if exponent <= 3 && mantissa != 0 && !isNegative {
			shiftedMantissa := mantissa >> (8 * (3 - exponent))
			require.Equal(t, int64(shiftedMantissa), result.Int64(),
				"Mantissa shift incorrect for exponent %d, input 0x%08x", exponent, compact)
		}
	})
}
