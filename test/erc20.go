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

func DeployErc20(opts *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, initialSupply *big.Int, timeout time.Duration) (common.Address, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	opts.Context = ctx
	address, _, _, err := bindings.DeployErc20(opts, backend, name, symbol, initialSupply)
	if err != nil {
		return address, fmt.Errorf("failed to deploy erc20 token: %w", err)
	}
	return address, nil
}

func TransferErc20(opts *bind.TransactOpts, backend bind.ContractBackend, tokenAddress, to common.Address, amount *big.Int, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	opts.Context = ctx
	erc20Transactor, err := bindings.NewErc20Transactor(tokenAddress, backend)
	if err != nil {
		return fmt.Errorf("failed to create erc20 transactor: %w", err)
	}
	_, err = erc20Transactor.Transfer(opts, to, amount)
	if err != nil {
		return fmt.Errorf("failed to transfer erc20 token: %w", err)
	}
	return nil
}
