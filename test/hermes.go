package test

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mysteriumnetwork/payments/bindings"
)

func DeployHermesImplementation(opts *bind.TransactOpts, backend bind.ContractBackend, timeout time.Duration) (common.Address, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	opts.Context = ctx
	hermesImplementationAddress, _, _, err := bindings.DeployHermesImplementation(opts, backend)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to deploy hermes implementation: %w", err)
	}
	return hermesImplementationAddress, nil
}
