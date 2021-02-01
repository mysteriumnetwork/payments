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

package bindings

import (
	"context"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

// EstimateOpts is the collection of authorization data required to perform validation of Ethereum transaction.
type EstimateOpts struct {
	From   common.Address // Ethereum account to send the transaction from
	Method string
	Params []interface{}

	Context context.Context // Network context to support cancellation and timeouts (nil = no timeout)
}

// ContractEstimator is an dryrun-only Go binding around an Ethereum contract.
type ContractEstimator struct {
	contractAddress common.Address          // Deployment address of the contract on the Ethereum blockchain
	contractAbi     abi.ABI                 // Reflect based ABI to access the correct Ethereum methods
	transactor      bind.ContractTransactor // Write interface to interact with the blockchain
}

// NewContractEstimator creates a new instance of ContractEstimator, bound to a specific deployed contract.
func NewContractEstimator(address common.Address, binding string, transactor bind.ContractTransactor) (*ContractEstimator, error) {
	parsed, err := abi.JSON(strings.NewReader(binding))
	if err != nil {
		return nil, err
	}

	return &ContractEstimator{
		contractAddress: address,
		contractAbi:     parsed,
		transactor:      transactor,
	}, nil
}

// Estimate simulates the (paid) contract method with params as input values and estimates gas needed.
func (drt *ContractEstimator) Estimate(opts *EstimateOpts) (uint64, error) {
	input, err := drt.contractAbi.Pack(opts.Method, opts.Params...)
	if err != nil {
		return 0, errors.Wrap(err, "could not pack input")
	}

	ctx := opts.Context
	if ctx == nil {
		ctx = context.TODO()
	}

	return drt.transactor.EstimateGas(ctx, ethereum.CallMsg{
		From: opts.From,
		To:   &drt.contractAddress,
		Data: input,
	})
}
