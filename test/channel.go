package test

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mysteriumnetwork/payments/bindings"
)

func DeployChannelImplementation(opts *bind.TransactOpts, backend bind.ContractBackend, timeout time.Duration) (common.Address, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	opts.Context = ctx
	address, _, _, err := bindings.DeployChannelImplementation(opts, backend)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to deploy channel implementation: %w", err)
	}
	return address, nil
}
