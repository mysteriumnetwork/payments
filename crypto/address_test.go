package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddress_Compare(t *testing.T) {
	a := HexToAddress("0xe7804c37c13166fF0b37F5aE0BB07A3aEbb6e245")
	b := HexToAddress("0xe7804c37c13166ff0b37F5ae0bb07a3aebb6e245")
	assert.Equal(t, a, b)
}
