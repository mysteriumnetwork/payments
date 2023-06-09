package test

import (
	"context"
	"fmt"
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
