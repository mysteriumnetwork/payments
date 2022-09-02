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
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
)

// Invoice represent a payment request
type Invoice struct {
	AgreementID    *big.Int
	AgreementTotal *big.Int
	TransactorFee  *big.Int
	Hashlock       string
	Provider       string
	ChainID        int64
}

// CreateInvoice creates new invoice
func CreateInvoice(agreementID, agreementTotal, transactorFee *big.Int, r []byte, chainID int64) (Invoice, error) {
	if r == nil {
		var err error
		r, err = GenerateR()
		if err != nil {
			return Invoice{}, err
		}
	}

	return Invoice{
		AgreementID:    new(big.Int).Set(agreementID),
		AgreementTotal: new(big.Int).Set(agreementTotal),
		TransactorFee:  new(big.Int).Set(transactorFee),
		Hashlock:       hex.EncodeToString(crypto.Keccak256(r)),
		ChainID:        chainID,
	}, nil
}

// GenerateR creates a random R
func GenerateR() ([]byte, error) {
	r := make([]byte, 32)
	_, err := rand.Read(r)
	if err != nil {
		return nil, fmt.Errorf("failed to generate r: %w", err)
	}
	return r, nil
}
