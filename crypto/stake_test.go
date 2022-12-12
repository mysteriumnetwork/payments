package crypto

import (
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStake(t *testing.T) {
	id := HexToAddress("0x1")
	hermes := HexToAddress("0x3")
	amount := big.NewInt(20)
	tf := big.NewInt(10)
	nonce := big.NewInt(3)
	chain := int64(1)

	chId, err := generateProviderChannelID(id.Hex(), hermes.Hex(), ProviderChannelTypePayAndSettle)
	assert.NoError(t, err)
	chIdBytes, err := hex.DecodeString(strings.TrimPrefix(chId, "0x"))
	assert.NoError(t, err)
	var chIdBytes32 [32]byte
	copy(chIdBytes32[:], chIdBytes)
	dps := DecreaseProviderStakeRequest{
		ChannelID:     chIdBytes32,
		HermesID:      hermes,
		Amount:        amount,
		TransactorFee: tf,
		Nonce:         nonce,
		ChainID:       chain,
	}

	pk, err := crypto.GenerateKey()
	require.NoError(t, err)
	publicKeyECDSA, ok := pk.Public().(*ecdsa.PublicKey)
	require.True(t, ok)
	address := FromCommonAddress(crypto.PubkeyToAddress(*publicKeyECDSA))

	t.Run("sign", func(t *testing.T) {
		err = dps.Sign(&pkHashSigner{pk: pk, address: address}, address)
		assert.NoError(t, err)
		assert.NotEqual(t, make([]byte, 65), dps.Signature)
	})

	t.Run("is valid", func(t *testing.T) {
		valid := dps.IsValid(HexToAddress("0x1234"))
		assert.False(t, valid)

		valid = dps.IsValid(address)
		assert.True(t, valid)
	})
}
