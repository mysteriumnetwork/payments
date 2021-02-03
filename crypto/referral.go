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
	"bytes"
	"encoding/hex"
	"errors"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type ReferralTokenRequest struct {
	Identity  common.Address
	Signature string
}

// CreateSignature signs promise using keystore
func CreateReferralTokenRequest(ks hashSigner, signer common.Address) (ReferralTokenRequest, error) {
	message := signer.Bytes()
	hash := crypto.Keccak256(message)
	signature, err := ks.SignHash(
		accounts.Account{Address: signer},
		hash,
	)
	if err != nil {
		return ReferralTokenRequest{}, err
	}

	return ReferralTokenRequest{
		Signature: hex.EncodeToString(signature),
		Identity:  signer,
	}, nil
}

func (rtr ReferralTokenRequest) IsValid() error {
	b, err := hex.DecodeString(rtr.Signature)
	if err != nil {
		return err
	}

	k, err := crypto.Ecrecover(crypto.Keccak256(rtr.Identity.Bytes()), b)
	if err != nil {
		return err
	}
	pubKey, err := crypto.UnmarshalPubkey(k)
	if err != nil {
		return err
	}

	recoveredAddress := crypto.PubkeyToAddress(*pubKey)

	if !bytes.Equal(rtr.Identity.Bytes(), recoveredAddress.Bytes()) {
		return errors.New("identities do not match")
	}

	return nil
}
