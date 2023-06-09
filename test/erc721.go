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

func DeployErc721(opts *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, timeout time.Duration) (common.Address, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	opts.Context = ctx
	address, _, _, err := bindings.DeployErc721(opts, backend, name, symbol)
	if err != nil {
		return address, fmt.Errorf("failed to deploy erc721 token: %w", err)
	}
	return address, nil
}

func TransferErc721(opts *bind.TransactOpts, backend bind.ContractBackend, tokenAddress, to common.Address, tokenId *big.Int, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	opts.Context = ctx
	erc721Transactor, err := bindings.NewErc721(tokenAddress, backend)
	if err != nil {
		return fmt.Errorf("failed to create erc721 transactor: %w", err)
	}
	_, err = erc721Transactor.TransferFrom(opts, opts.From, to, tokenId)
	if err != nil {
		return fmt.Errorf("failed to transfer erc721 token: %w", err)
	}
	return nil
}

func MintErc721(opts *bind.TransactOpts, backend bind.ContractBackend, tokenAddress, to common.Address, tokenId *big.Int, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	opts.Context = ctx
	erc721Transactor, err := bindings.NewErc721(tokenAddress, backend)
	if err != nil {
		return fmt.Errorf("failed to create erc721 transactor: %w", err)
	}
	_, err = erc721Transactor.Mint(opts, to, tokenId)
	if err != nil {
		return fmt.Errorf("failed to mint erc721 token: %w", err)
	}
	return nil
}
