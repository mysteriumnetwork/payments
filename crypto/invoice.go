/*
 * Copyright (C) 2019 The "MysteriumNetwork/payments" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package crypto

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
)

// Invoice represent a payment request
type Invoice struct {
	AgreementID    uint64
	AgreementTotal uint64
	Fee            uint64
	Hashlock       string
	Provider       string
}

// CreateInvoice creates new invoice
func CreateInvoice(agreementID, agreementTotal, fee uint64, r []byte) Invoice {
	if r == nil {
		r = make([]byte, 64)
		rand.Read(r)
	}

	return Invoice{
		AgreementID:    agreementID,
		AgreementTotal: agreementTotal,
		Fee:            fee,
		Hashlock:       hex.EncodeToString(crypto.Keccak256(r)),
	}
}
