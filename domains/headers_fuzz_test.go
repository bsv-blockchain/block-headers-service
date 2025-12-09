package domains

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

// FuzzFastLog2Floor tests the FastLog2Floor function with random uint32 inputs.
// FastLog2Floor calculates the floor of base-2 logarithm using bitwise operations.
// The function should never panic and always return a value in [0, 31].
func FuzzFastLog2Floor(f *testing.F) {
	// Seed corpus with edge cases and powers of 2
	f.Add(uint32(0))          // Zero edge case
	f.Add(uint32(1))          // 2^0
	f.Add(uint32(2))          // 2^1
	f.Add(uint32(3))          // Between 2^1 and 2^2
	f.Add(uint32(4))          // 2^2
	f.Add(uint32(7))          // 2^3 - 1
	f.Add(uint32(8))          // 2^3
	f.Add(uint32(15))         // 2^4 - 1
	f.Add(uint32(16))         // 2^4
	f.Add(uint32(255))        // 2^8 - 1
	f.Add(uint32(256))        // 2^8
	f.Add(uint32(65535))      // 2^16 - 1
	f.Add(uint32(65536))      // 2^16
	f.Add(uint32(0x7fffffff)) // Max int32
	f.Add(uint32(0x80000000)) // 2^31
	f.Add(uint32(0xffffffff)) // Max uint32 (2^32 - 1)

	// All powers of 2 up to 2^31
	for i := uint32(0); i < 32; i++ {
		f.Add(uint32(1) << i)
	}

	f.Fuzz(func(t *testing.T, n uint32) {
		// FastLog2Floor should never panic
		result := FastLog2Floor(n)

		// Result must always be in valid range [0, 31]
		require.LessOrEqual(t, result, uint8(31),
			"Result %d exceeds maximum 31 for input %d", result, n)

		if n == 0 {
			// For n=0, the function returns 0 (edge case behavior)
			require.Equal(t, uint8(0), result,
				"Expected 0 for input 0, got %d", result)
			return
		}

		// For n > 0: verify that 2^result <= n < 2^(result+1)
		// This is the mathematical definition of floor(log2(n))

		// 2^result should be <= n
		lowerBound := uint32(1) << result
		require.LessOrEqual(t, lowerBound, n,
			"2^%d = %d should be <= %d", result, lowerBound, n)

		// 2^(result+1) should be > n (unless result is 31)
		if result < 31 {
			upperBound := uint32(1) << (result + 1)
			require.Less(t, n, upperBound,
				"Input %d should be < 2^%d = %d", n, result+1, upperBound)
		}

		// Cross-validate with math.Log2 for non-zero values
		expected := uint8(math.Floor(math.Log2(float64(n))))
		require.Equal(t, expected, result,
			"FastLog2Floor(%d) = %d, but math.Floor(math.Log2(%d)) = %d",
			n, result, n, expected)
	})
}
