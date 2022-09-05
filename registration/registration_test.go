package registration

import (
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegistrationRequest(t *testing.T) {
	req := &Request{
		ChainID:         1,
		HermesID:        "0x0000000000000000000000000000000000000001",
		Stake:           big.NewInt(2),
		Fee:             big.NewInt(3),
		Beneficiary:     "0x0000000000000000000000000000000000000002",
		RegistryAddress: "0x0000000000000000000000000000000000000003",
	}

	t.Run("getter", func(t *testing.T) {
		assert.Equal(t, big.NewInt(2), req.GetStakeAmount())
		assert.Equal(t, big.NewInt(3), req.GetFee())
	})

	t.Run("sign", func(t *testing.T) {
		pk, err := crypto.GenerateKey()
		require.NoError(t, err)
		publicKeyECDSA, ok := pk.Public().(*ecdsa.PublicKey)
		require.True(t, ok)
		address := crypto.PubkeyToAddress(*publicKeyECDSA)
		signature, err := crypto.Sign(crypto.Keccak256(req.GetMessage()), pk)
		assert.NoError(t, err)
		req.Signature = hex.EncodeToString(signature)

		signer, err := req.RecoverIdentity()
		assert.NoError(t, err)
		assert.Equal(t, address, signer)
	})
}

func TestOpenChannelRequest(t *testing.T) {
	req := &OpenConsumerChannelRequest{
		ChainID:         1,
		HermesID:        "0x0000000000000000000000000000000000000001",
		RegistryAddress: "0x0000000000000000000000000000000000000003",
		TransactorFee:   big.NewInt(3),
	}

	t.Run("sign", func(t *testing.T) {
		pk, err := crypto.GenerateKey()
		require.NoError(t, err)
		publicKeyECDSA, ok := pk.Public().(*ecdsa.PublicKey)
		require.True(t, ok)
		address := crypto.PubkeyToAddress(*publicKeyECDSA)
		signature, err := crypto.Sign(crypto.Keccak256(req.GetMessage()), pk)
		assert.NoError(t, err)
		req.Signature = hex.EncodeToString(signature)

		signer, err := req.RecoverIdentity()
		assert.NoError(t, err)
		assert.Equal(t, address, signer)
	})
}
