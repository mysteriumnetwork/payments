package crypto

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExit(t *testing.T) {
	var er *ExitRequest
	ch := common.HexToAddress("0x1")
	benef := common.HexToAddress("0x2")
	validUntil := big.NewInt(20)

	t.Run("new", func(t *testing.T) {
		er = NewExitRequest(ch, benef, validUntil)
		assert.Equal(t, ch, er.ChannelID)
		assert.Equal(t, benef, er.Beneficiary)
		assert.Equal(t, validUntil, er.ValidUntil)
	})

	pk, err := crypto.GenerateKey()
	require.NoError(t, err)
	publicKeyECDSA, ok := pk.Public().(*ecdsa.PublicKey)
	require.True(t, ok)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	t.Run("sign", func(t *testing.T) {
		err = er.Sign(&pkHashSigner{pk: pk, address: address}, address)
		assert.NoError(t, err)
		assert.NotEqual(t, make([]byte, 65), er.Signature)
	})

	t.Run("recover signer", func(t *testing.T) {
		signer, err := er.RecoverSigner()
		assert.NoError(t, err)
		assert.Equal(t, address, signer)
	})
}
