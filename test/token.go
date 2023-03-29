package test

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mysteriumnetwork/payments/bindings"
)

func DeployOldToken(opts *bind.TransactOpts, backend bind.ContractBackend, timeout time.Duration) (common.Address, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	opts.Context = ctx
	address, _, _, err := bindings.DeployOldMystToken(opts, backend)
	if err != nil {
		return address, fmt.Errorf("failed to deploy old token: %w", err)
	}
	return address, nil
}

func DeployTokenV2(opts *bind.TransactOpts, backend bind.ContractBackend, oldTokenAddress common.Address, timeout time.Duration) (common.Address, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	opts.Context = ctx
	address, _, _, err := bindings.DeployMystToken(opts, backend, oldTokenAddress)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to deploy token v2: %w", err)
	}
	return address, nil
}

type TokenContractAddresses struct {
	TokenAddress   common.Address
	TokenV2Address common.Address
}

func DeployTokenV2WithDependencies(opts *bind.TransactOpts, backend bind.ContractBackend, timeout time.Duration) (TokenContractAddresses, error) {
	oldTokenAddress, err := DeployOldToken(opts, backend, timeout)
	if err != nil {
		return TokenContractAddresses{}, fmt.Errorf("failed to deploy old token: %w", err)
	}
	tokenAddress, err := DeployTokenV2(opts, backend, oldTokenAddress, timeout)
	if err != nil {
		return TokenContractAddresses{}, fmt.Errorf("failed to deploy token v2: %w", err)
	}
	return TokenContractAddresses{
		TokenAddress:   oldTokenAddress,
		TokenV2Address: tokenAddress,
	}, nil
}

func MintTokens(opts *bind.TransactOpts, tokenAddress common.Address, transactor bind.ContractTransactor, timeout time.Duration, recipient common.Address, amount *big.Int) error {
	tokenTransactor, err := bindings.NewMystTokenTransactor(tokenAddress, transactor)
	if err != nil {
		return fmt.Errorf("failed to create token transactor: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	opts.Context = ctx
	_, err = tokenTransactor.Mint(opts, recipient, amount)
	if err != nil {
		return fmt.Errorf("failed to mint tokens: %w", err)
	}
	return nil
}
