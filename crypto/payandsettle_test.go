package crypto

import (
	"crypto/ecdsa"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPayAndSettle(t *testing.T) {
	var pas *PayAndSettleBeneficiaryPayload
	id := common.HexToAddress("0x1")
	benef := common.HexToAddress("0x2")
	hermes := common.HexToAddress("0x3")
	amount := big.NewInt(20)
	chain := int64(1)

	t.Run("new", func(t *testing.T) {
		chId, err := generateProviderChannelID(id.Hex(), hermes.Hex(), ProviderChannelTypePayAndSettle)
		assert.NoError(t, err)
		r, err := GenerateR()
		assert.NoError(t, err)
		var r32 [32]byte
		copy(r32[:], r)
		pas = NewPayAndSettleBeneficiaryPayload(benef, chain, chId, amount, r32)
		assert.Equal(t, benef, pas.Beneficiary)
		assert.Equal(t, chain, pas.ChainID)
		assert.Equal(t, strings.TrimPrefix(chId, "0x"), pas.ProviderChannelIDForWithdrawal)
		assert.Equal(t, amount, pas.Amount)
		assert.Equal(t, r32, pas.R)
	})

	pk, err := crypto.GenerateKey()
	require.NoError(t, err)
	publicKeyECDSA, ok := pk.Public().(*ecdsa.PublicKey)
	require.True(t, ok)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	t.Run("sign", func(t *testing.T) {
		err = pas.Sign(&pkHashSigner{pk: pk, address: address}, address)
		assert.NoError(t, err)
		assert.NotEqual(t, make([]byte, 65), pas.Signature)
	})

	t.Run("recover signer", func(t *testing.T) {
		signer, err := pas.RecoverSigner()
		assert.NoError(t, err)
		assert.Equal(t, address, signer)
	})
}
