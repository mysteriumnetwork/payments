package crypto

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBeneficiary(t *testing.T) {
	t.Run("beneficiary set request", func(t *testing.T) {
		chain := int64(1)
		id := HexToAddress("0x1")
		reg := HexToAddress("0x2")
		ben := HexToAddress("0x3")
		nonce := big.NewInt(1)
		signature := common.Hex2Bytes(signatureHex)
		err := ReformatSignatureVForBC(signature)
		assert.NoError(t, err)

		t.Run("create", func(t *testing.T) {
			req, err := CreateBeneficiaryRequest(chain, id.Hex(), reg.Hex(), ben.Hex(), nonce, &mockHashSigner{}, id)
			assert.NoError(t, err)
			assert.NotNil(t, req)
			assert.Equal(t, chain, req.ChainID)
			assert.Equal(t, nonce, req.Nonce)
			assert.Equal(t, strings.TrimPrefix(id.Hex(), "0x"), req.Identity)
			assert.Equal(t, strings.TrimPrefix(reg.Hex(), "0x"), req.Registry)
			assert.Equal(t, strings.TrimPrefix(ben.Hex(), "0x"), req.Beneficiary)
			assert.Equal(t, hex.EncodeToString(signature), req.Signature)
		})

		t.Run("new", func(t *testing.T) {
			req, err := NewBeneficiaryRequest(chain, id.Hex(), reg.Hex(), ben.Hex(), nonce, hex.EncodeToString(signature))
			assert.NoError(t, err)
			assert.NotNil(t, req)
			assert.Equal(t, chain, req.ChainID)
			assert.Equal(t, nonce, req.Nonce)
			assert.Equal(t, strings.TrimPrefix(id.Hex(), "0x"), req.Identity)
			assert.Equal(t, strings.TrimPrefix(reg.Hex(), "0x"), req.Registry)
			assert.Equal(t, strings.TrimPrefix(ben.Hex(), "0x"), req.Beneficiary)
			assert.Equal(t, hex.EncodeToString(signature), req.Signature)
		})

		t.Run("bytes raw", func(t *testing.T) {
			req, err := NewBeneficiaryRequest(chain, id.Hex(), reg.Hex(), ben.Hex(), nonce, hex.EncodeToString(signature))
			require.NoError(t, err)
			signatureBytes := req.GetSignatureBytesRaw()
			assert.Equal(t, signature, signatureBytes)
		})

		t.Run("recover signer", func(t *testing.T) {
			pk, err := crypto.GenerateKey()
			require.NoError(t, err)
			publicKeyECDSA, ok := pk.Public().(*ecdsa.PublicKey)
			require.True(t, ok)
			address := FromCommonAddress(crypto.PubkeyToAddress(*publicKeyECDSA))
			req, err := CreateBeneficiaryRequest(chain, address.Hex(), reg.Hex(), ben.Hex(), nonce, &pkHashSigner{pk: pk, address: address}, address)
			require.NoError(t, err)

			signer, err := req.RecoverSigner()
			assert.NoError(t, err)
			assert.Equal(t, address, signer)
		})
	})
}

const signatureHex = "abcd1231231231231231231231231231231231231231231231231231231231231abcd12312312312312312312312312312312312312312312312312312312312312"

type mockHashSigner struct{}

func (mhs *mockHashSigner) SignHash(a accounts.Account, hash []byte) ([]byte, error) {
	return common.Hex2Bytes(signatureHex), nil
}

type pkHashSigner struct {
	pk      *ecdsa.PrivateKey
	address Address
}

func (phs *pkHashSigner) SignHash(a accounts.Account, hash []byte) ([]byte, error) {
	if FromCommonAddress(a.Address) != phs.address {
		return nil, fmt.Errorf("wrong address")
	}
	return crypto.Sign(hash, phs.pk)
}
