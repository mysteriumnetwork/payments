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
	"encoding/hex"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

type Metadata struct {
	ChannelID   string
	Beneficiary string
	Nonce       *big.Int
	Signature   string
	Message     string
}

func getMetadata() Metadata {
	return Metadata{
		ChannelID:   "0xcdce1530994e9d86c9446c42209057b4b6cee45ff44c6e746265f50b2051d019",
		Beneficiary: "0xB6352B90f85CE9992706a80760F0A025FCEE7b0A",
		Nonce:       big.NewInt(1),
		Signature:   "333f2648ebdbc7d8e4f8d857f4c24a8a8b43951ab9c84f176bdbc8427a640f736f53201a8491b2c22bae91ddee8efbcbfff5bb47957c1d2af83d0245f2de63b41b",
		Message:     "cdce1530994e9d86c9446c42209057b4b6cee45ff44c6e746265f50b2051d019b6352b90f85ce9992706a80760f0a025fcee7b0a0000000000000000000000000000000000000000000000000000000000000001",
	}
}

func TestBeneficiarySignature(t *testing.T) {
	dir, ks := tmpKeyStore(t, false)
	defer os.RemoveAll(dir)

	account, err := ks.ImportECDSA(getPrivKey("provider"), "")
	assert.Nil(t, err)
	if err := ks.Unlock(account, ""); err != nil {
		t.Fatal(err)
	}

	data := getMetadata()

	request, err := CreateBeneficiaryRequest(data.ChannelID, data.Beneficiary, data.Nonce, ks, account.Address)
	assert.Nil(t, err)

	assert.Equal(t, request.Signature, data.Signature)
}

func TestGetMessage(t *testing.T) {
	data := getMetadata()
	request, err := NewBeneficiaryRequest(data.ChannelID, data.Beneficiary, data.Nonce, data.Signature)
	assert.Nil(t, err)

	message := request.GetMessage()

	expectedMessage, _ := hex.DecodeString(data.Message)
	assert.Equal(t, expectedMessage, message)
}

func TestRecoverBeneficiarySigner(t *testing.T) {
	data := getMetadata()
	request, err := NewBeneficiaryRequest(data.ChannelID, data.Beneficiary, data.Nonce, data.Signature)
	assert.Nil(t, err)

	expectedSigner := common.HexToAddress("0x354bd098b4ef8c9e70b7f21be2d455df559705d7")
	recoveredSigner, err := request.RecoverSigner()
	assert.Nil(t, err)
	assert.Equal(t, expectedSigner, recoveredSigner)
}
