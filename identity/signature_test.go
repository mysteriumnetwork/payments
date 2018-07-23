package identity

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

var testSignature = common.FromHex(
	"1F1F1F1F1F1F1F1F1F1F1F1F1F1F1F1F1F1F1F1F1F1F1F1F1F1F1F1F1F1F1F1F" + //R part - 32 bytes
		"ABABABABABABABABABABABABABABABABABABABABABABABABABABABABABABABAB" + //S part - 32 bytes
		"01", //V part - 32 bytes
)

func TestSignatureIsDecomposedCorrectly(t *testing.T) {
	decomposed, err := DecomposeSignature(testSignature)
	assert.NoError(t, err)
	assertArrayIsFilledWith(t, decomposed.R, 0x1f)
	assertArrayIsFilledWith(t, decomposed.S, 0xab)
	assert.Equal(t, uint8(28), decomposed.V)
}

func assertArrayIsFilledWith(t *testing.T, arr [32]byte, val uint8) {
	for i, v := range arr {
		assert.Equal(t, val, v, "Expected 0x%X on index: %d. Got: 0x%X", val, i, v)
	}
}
