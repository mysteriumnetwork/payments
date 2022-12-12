/* Mysterium network payment library.
 *
 * Copyright (C) 2021 BlockDev AG
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

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
)

type SignedIdentityRequest struct {
	Identity  Address `json:"identity,omitempty"`
	Signature string  `json:"signature,omitempty"`
}

func CreateSignedIdentityRequest(ks hashSigner, identityToSign, signer Address) (SignedIdentityRequest, error) {
	message := identityToSign.Bytes()
	hash := crypto.Keccak256(message)

	signature, err := ks.SignHash(
		accounts.Account{Address: signer.ToCommon()},
		hash,
	)
	if err != nil {
		return SignedIdentityRequest{}, err
	}

	return SignedIdentityRequest{
		Signature: hex.EncodeToString(signature),
		Identity:  identityToSign,
	}, nil
}

func (sir SignedIdentityRequest) RecoverSigner() (Address, error) {
	b, err := hex.DecodeString(sir.Signature)
	if err != nil {
		return Address{}, err
	}

	h := crypto.Keccak256(sir.Identity.Bytes())
	k, err := crypto.Ecrecover(h, b)
	if err != nil {
		return Address{}, err
	}
	pubKey, err := crypto.UnmarshalPubkey(k)
	if err != nil {
		return Address{}, err
	}

	return FromCommonAddress(crypto.PubkeyToAddress(*pubKey)), nil
}
