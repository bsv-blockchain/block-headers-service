// Copyright (c) 2024 The bitcoin-sv developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chainhash

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

// FuzzDecode tests the Decode function with random string inputs.
// Decode parses hexadecimal strings into 256-bit hashes with byte reversal.
// The function should never panic regardless of input.
func FuzzDecode(f *testing.F) {
	// Seed corpus with valid Bitcoin hashes
	f.Add("000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f") // Genesis
	f.Add("000000000003ba27aa200b1cecaad478d2b00432346c3f1f3986da1afd33e506") // Block 100000
	f.Add("00000000000000000007878ec04bb2b2e12317804810f4c26033585b3f81ffaa") // Recent block

	// Edge cases - valid hex
	f.Add("")                                                                 // Empty string
	f.Add("0")                                                                // Single zero
	f.Add("1")                                                                // Single digit
	f.Add("00")                                                               // Two zeros
	f.Add("ff")                                                               // Max byte
	f.Add("abcdef")                                                           // Short valid hex
	f.Add("ABCDEF")                                                           // Uppercase hex
	f.Add("AbCdEf")                                                           // Mixed case
	f.Add("0000000000000000000000000000000000000000000000000000000000000000") // All zeros (64 chars)
	f.Add("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff") // All ones (64 chars)
	f.Add("123456789abcdef")                                                  // Odd length (will be padded)

	// Edge cases - boundary lengths
	f.Add("00000000000000000000000000000000000000000000000000000000000000000") // 65 chars (too long)

	// Invalid inputs (should return error, not panic)
	f.Add("ghijkl")                                                             // Invalid hex chars
	f.Add("0x1234")                                                             // Prefix (invalid)
	f.Add("12 34")                                                              // Space (invalid)
	f.Add("12-34")                                                              // Dash (invalid)
	f.Add("\x00\x01\x02")                                                       // Binary data
	f.Add("zzzzzzzz")                                                           // All invalid
	f.Add("12345g")                                                             // Mixed valid/invalid
	f.Add("\n\r\t")                                                             // Whitespace
	f.Add("000000000000000000000000000000000000000000000000000000000000000000") // 66 chars

	f.Fuzz(func(t *testing.T, src string) {
		var dst Hash

		// Decode should never panic regardless of input
		err := Decode(&dst, src)
		if err != nil {
			// Error is acceptable - verify the hash wasn't modified unexpectedly
			// Just return since error handling is the expected path for invalid input
			return
		}

		// If no error, verify the result is valid

		// The decoded hash should have valid length
		require.Len(t, dst, HashSize,
			"Decoded hash should be %d bytes", HashSize)

		// For valid decodes, verify we can convert back to string
		hashStr := dst.String()
		require.NotEmpty(t, hashStr, "String() should not return empty for valid hash")
		require.Len(t, hashStr, MaxHashStringSize,
			"String() should return %d char hex string", MaxHashStringSize)

		// Verify the string is valid hex
		_, hexErr := hex.DecodeString(hashStr)
		require.NoError(t, hexErr, "String() output should be valid hex")

		// Verify round-trip: decode the string output and compare
		// Note: Due to byte reversal and padding, we can't always do exact equality
		// but we can verify the round-trip is consistent
		var roundTrip Hash
		roundTripErr := Decode(&roundTrip, hashStr)
		require.NoError(t, roundTripErr, "Round-trip decode should succeed")
		require.Equal(t, dst, roundTrip, "Round-trip should produce identical hash")
	})
}

// FuzzNewHashFromStr tests the NewHashFromStr function with random string inputs.
// This is a wrapper around Decode that allocates a new Hash.
func FuzzNewHashFromStr(f *testing.F) {
	// Reuse the same seed corpus as FuzzDecode
	f.Add("000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f")
	f.Add("")
	f.Add("1")
	f.Add("abcdef")
	f.Add("ghijkl") // Invalid
	f.Add("0000000000000000000000000000000000000000000000000000000000000000")
	f.Add("00000000000000000000000000000000000000000000000000000000000000000") // Too long

	f.Fuzz(func(t *testing.T, hashStr string) {
		// NewHashFromStr should never panic
		hash, err := NewHashFromStr(hashStr)
		if err != nil {
			// Error is acceptable - verify hash is nil
			require.Nil(t, hash, "Hash should be nil when error is returned")
			return
		}

		// If no error, hash must not be nil
		require.NotNil(t, hash, "Hash should not be nil when no error")

		// Verify it's the correct size
		require.Len(t, *hash, HashSize,
			"Hash should be %d bytes", HashSize)
	})
}
