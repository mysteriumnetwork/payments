/* Mysterium network payment library.
 *
 * Copyright (C) 2020 BlockDev AG
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package crypto

import (
	"crypto/ecdsa"
	"encoding/hex"
	"io/ioutil"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestGetHash(t *testing.T) {
	promise := getPromise("provider")
	expectedHash, _ := hex.DecodeString("e3d72d47da46d3b9ca2a9b84136f626a1ac82630c09edc20d206fffe69a2b4a3")
	assert.Equal(t, expectedHash, promise.GetHash())

	promise = getPromise("consumer")
	expectedHash, _ = hex.DecodeString("a932aa523e07da71e257cadd57a8ed49cf1ee9d0ce16c97636c052496b9da73c")
	assert.Equal(t, expectedHash, promise.GetHash())
}

func TestGetSignatureHex(t *testing.T) {
	promise := getPromise("provider")
	signature := promise.GetSignatureHexString()

	expectedSignature := "0x517e12816b37dd08750a805d45b375c0d672ef8141dc441bb092683f6599718772a8b8308c1baee662cd898d7d5dbfac22269a6dc0d469a0277042db347131a51c"
	assert.Equal(t, expectedSignature, signature)
}

func TestCreateConsumerSignature(t *testing.T) {
	dir, ks := tmpKeyStore(t, false)
	defer os.RemoveAll(dir)

	account, err := ks.ImportECDSA(getPrivKey("consumer"), "")
	assert.Nil(t, err)
	if err := ks.Unlock(account, ""); err != nil {
		t.Fatal(err)
	}

	promise := getPromise("consumer")

	signature, err := promise.CreateSignature(ks, account.Address)
	assert.Nil(t, err)

	ReformatSignatureVForBC(signature)
	assert.Equal(t, promise.Signature, signature)
}

func TestCreateProviderSignature(t *testing.T) {
	dir, ks := tmpKeyStore(t, false)
	defer os.RemoveAll(dir)

	account, err := ks.ImportECDSA(getPrivKey("provider"), "")
	assert.Nil(t, err)
	if err := ks.Unlock(account, ""); err != nil {
		t.Fatal(err)
	}

	promise := getPromise("provider")

	signature, err := promise.CreateSignature(ks, account.Address)
	assert.Nil(t, err)

	ReformatSignatureVForBC(signature)
	assert.Equal(t, promise.Signature, signature)
}

func TestConsumerPromiseValidation(t *testing.T) {
	promise := getPromise("consumer")
	expectedSigner := common.HexToAddress("0xf53acdd584ccb85ee4ec1590007ad3c16fdff057")
	assert.True(t, promise.IsPromiseValid(expectedSigner))
}

func TestProviderPromiseValidation(t *testing.T) {
	promise := getPromise("provider")
	expectedSigner := common.HexToAddress("0x354bd098b4ef8c9e70b7f21be2d455df559705d7")
	assert.True(t, promise.IsPromiseValid(expectedSigner))

	wrongSigner := common.HexToAddress("0xf53acdd584ccb85ee4ec1590007ad3c16fdff057")
	assert.False(t, promise.IsPromiseValid(wrongSigner))
}

func TestRecoverSigner(t *testing.T) {
	promise := getPromise("consumer")
	expectedSigner := common.HexToAddress("0xf53acdd584ccb85ee4ec1590007ad3c16fdff057")
	recoveredSigner, err := promise.RecoverSigner()
	assert.Nil(t, err)
	assert.Equal(t, expectedSigner, recoveredSigner)
}

// This test ensures that promise validation functions will not mutate original signature
func TestValidationAndRecoveryImmutability(t *testing.T) {
	promise := getPromise("consumer")

	originalSignature := make([]byte, 65)
	copy(originalSignature, promise.Signature)

	// Ensure that IsPromiseValid will not mutate original signature
	expectedSigner := common.HexToAddress("0xf53acdd584ccb85ee4ec1590007ad3c16fdff057")
	promise.IsPromiseValid(expectedSigner)
	assert.Equal(t, originalSignature, promise.Signature)

	// Ensure that RecoverSigner will not mutate original signature
	promise.RecoverSigner()
	assert.Equal(t, originalSignature, promise.Signature)
}

func TestCreatePromise(t *testing.T) {
	dir, ks := tmpKeyStore(t, false)
	defer os.RemoveAll(dir)

	account, err := ks.ImportECDSA(getPrivKey("consumer"), "")
	assert.Nil(t, err)
	if err := ks.Unlock(account, ""); err != nil {
		t.Fatal(err)
	}

	p := getParams("consumer")
	channelID := hex.EncodeToString(p.ChannelID)
	amount := big.NewInt(0).SetUint64(p.Amount)
	fee := big.NewInt(0).SetUint64(p.Fee)
	hashlock := hex.EncodeToString(p.Hashlock)

	promise, err := CreatePromise(channelID, 1, amount, fee, hashlock, ks, account.Address)
	assert.NoError(t, err)
	assert.Equal(t, p.PromiseSignature, promise.Signature)

	// ChannelID can also be provided with prefix
	promise, err = CreatePromise("0x"+channelID, 1, amount, fee, hashlock, ks, account.Address)
	assert.NoError(t, err)
	assert.Equal(t, p.PromiseSignature, promise.Signature)

	// Should fail when provided not correct channel ID
	_, err = CreatePromise("NotHex", 1, amount, fee, hashlock, ks, account.Address)
	assert.Error(t, err)
	assert.Equal(t, "channelID and hashlock have to be proper hex strings", err.Error())
}

func TestNewPromise(t *testing.T) {
	p := getParams("provider")
	channelID := hex.EncodeToString(p.ChannelID)
	amount := big.NewInt(0).SetUint64(p.Amount)
	fee := big.NewInt(0).SetUint64(p.Fee)
	preimage := hex.EncodeToString(p.R)
	signature := hex.EncodeToString(p.PromiseSignature)

	promise, err := NewPromise(1, channelID, amount, fee, preimage, signature)
	assert.NoError(t, err)
	assert.Equal(t, big.NewInt(0).SetUint64(p.Amount), promise.Amount)
	assert.Equal(t, p.ChannelID, promise.ChannelID)
	assert.Equal(t, p.Hashlock, promise.Hashlock)
	assert.Equal(t, p.PromiseSignature, promise.Signature)
}

func TestSign(t *testing.T) {
	dir, ks := tmpKeyStore(t, false)
	defer os.RemoveAll(dir)

	account, err := ks.ImportECDSA(getPrivKey("provider"), "")
	assert.Nil(t, err)
	if err := ks.Unlock(account, ""); err != nil {
		t.Fatal(err)
	}

	p := getParams("provider")
	amount := big.NewInt(0).SetUint64(p.Amount)
	fee := big.NewInt(0).SetUint64(p.Fee)

	promise, err := NewPromise(1,
		hex.EncodeToString(p.ChannelID),
		amount,
		fee,
		hex.EncodeToString(p.R),
		"",
	)
	assert.NoError(t, err)

	promise.Sign(ks, account.Address)
	assert.Equal(t, p.PromiseSignature, promise.Signature)
}

const (
	veryLightScryptN = 2
	veryLightScryptP = 1
)

func tmpKeyStore(t *testing.T, encrypted bool) (string, *keystore.KeyStore) {
	d, err := ioutil.TempDir("", "eth-keystore-test")
	if err != nil {
		t.Fatal(err)
	}
	newKs := keystore.NewPlaintextKeyStore
	if encrypted {
		newKs = func(kd string) *keystore.KeyStore {
			return keystore.NewKeyStore(kd, veryLightScryptN, veryLightScryptP)
		}
	}
	return d, newKs(d)
}

func getPrivKey(userType string) *ecdsa.PrivateKey {
	var privKey []byte

	if userType == "consumer" {
		privKey, _ = hex.DecodeString("1fd26004a85a0ae04e16b444893ba7408dc7215b3962dda782d1afc6d7441d1e")
	}

	if userType == "provider" {
		privKey, _ = hex.DecodeString("45bb96530f3d1972fdcd2005c1987a371d0b6d378b77561c6beeaca27498f46b")
	}

	pk, _ := crypto.ToECDSA(privKey)
	return pk
}

func getPromise(userType string) Promise {
	p := getParams(userType)
	amount := big.NewInt(0).SetUint64(p.Amount)
	fee := big.NewInt(0).SetUint64(p.Fee)
	promise := Promise{
		ChannelID: p.ChannelID,
		Amount:    amount,
		Fee:       fee,
		Hashlock:  p.Hashlock,
		Signature: p.PromiseSignature,
		ChainID:   1,
	}

	return promise
}

type Params struct {
	ChannelID                []byte
	Amount                   uint64
	Fee                      uint64
	R                        []byte
	Hashlock                 []byte
	PromiseSignature         []byte
	ExchangeMessageSignature string
	Provider                 common.Address
}

func getParams(userType string) Params {
	var channelID []byte
	var promiseSig []byte

	amount := uint64(1401)
	fee := uint64(0)
	provider := common.HexToAddress("0xf10021ba3b10d023e671668d20daeff821561d09")
	preimage, _ := hex.DecodeString("5b6b3f31a3acd0e317173d25c8b60503547b741a0e81d6068bb88486967839fa")
	hashlock, _ := hex.DecodeString("4e8444e4bd5721ba00ceb2c6c180c21b2ae43e590172f1b39e51f46312243633")
	messageSig := "3310f94760eeb288598d69aaf6214e123766d3cc988f5d3c6a77d5d5b70dc4b6617eeeb0f49df467527ed4a9cdf6f5c7ae3a1477a84f853b9ef8587607d0be431b"

	if userType == "consumer" {
		channelID, _ = hex.DecodeString("000000000000000000000000d2c94475763fa7e81076ab0bde4dc4b902191498")
		promiseSig, _ = hex.DecodeString("72051de6f5ca0eaef86ecca3b303fe03b7aca32cc5dc4cbf894f73737f43e24d7599220bfd69a87252f3826e07e28d6ff4704ec316885252bc4f534abcb413d41b")
	}

	if userType == "provider" {
		channelID, _ = hex.DecodeString("40684b7886aa813a04bce9efd1c6d45379178e9248bf1cc926a316e4e4924ae8")
		promiseSig, _ = hex.DecodeString("517e12816b37dd08750a805d45b375c0d672ef8141dc441bb092683f6599718772a8b8308c1baee662cd898d7d5dbfac22269a6dc0d469a0277042db347131a51c")
	}

	return Params{
		ChannelID:                channelID,
		Amount:                   amount,
		Fee:                      fee,
		R:                        preimage,
		Hashlock:                 hashlock,
		PromiseSignature:         promiseSig,
		ExchangeMessageSignature: messageSig,
		Provider:                 provider,
	}
}
