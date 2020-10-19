package bindings

import (
	"context"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
)

// DryRunOpts is the collection of authorization data required to perform validation of Ethereum transaction.
type DryRunOpts struct {
	From    common.Address  // Ethereum account to send the transaction from
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
func (drt *ContractEstimator) Estimate(opts *DryRunOpts, method string, params ...interface{}) (uint64, error) {
	input, err := drt.contractAbi.Pack(method, params...)
	if err != nil {
		return 0, errors.Wrap(err, "could not pack input")
	}

	ctx := opts.Context
	if ctx == nil {
		ctx = context.TODO()
	}

	gas, err := drt.transactor.EstimateGas(ctx, ethereum.CallMsg{
		From: opts.From,
		To:   &drt.contractAddress,
		Data: input,
	})

	return gas, errors.Wrap(err, "could not estimate gas")
}

// DryRun simulates the (paid) contract method with params as input values.
func (drt *ContractEstimator) DryRun(opts *DryRunOpts, method string, params ...interface{}) error {
	_, err := drt.Estimate(opts, method, params)
	if err == nil {
		return nil
	}

	// Extract error thrown in contract
	if rpcCauseErr, hasCause := errors.Cause(err).(rpc.Error); hasCause {
		if rpcCauseErr.ErrorCode() == -32000 && strings.Contains(rpcCauseErr.Error(), "VM Exception while processing transaction: revert") {
			err = &ErrorTransactionReverted{
				Err:    rpcCauseErr,
				Reason: strings.TrimPrefix(rpcCauseErr.Error(), "VM Exception while processing transaction: revert "),
			}
		}
	}

	return err
}

// ErrorTransactionReverted when contract execution is interrupted with an error./
type ErrorTransactionReverted struct {
	Err    rpc.Error
	Reason string
}

func (e ErrorTransactionReverted) Error() string {
	return e.Err.Error()
}
