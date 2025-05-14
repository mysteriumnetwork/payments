package test

import (
	"context"
	"crypto/ecdsa"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mysteriumnetwork/payments/v3/bindings"
	pc "github.com/mysteriumnetwork/payments/v3/crypto"
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

func UpdateHermesURL(opts *bind.TransactOpts, backend bind.ContractBackend, registryAddress, hermesAddress common.Address, ownerPk *ecdsa.PrivateKey, url string, timeout time.Duration) error {
	registry, err := bindings.NewRegistry(registryAddress, backend)
	if err != nil {
		return fmt.Errorf("failed to create registry transactor: %w", err)
	}

	lastNonce, err := registry.LastNonce(nil)
	if err != nil {
		return fmt.Errorf("failed to get last nonce: %w", err)
	}
	signature, err := getHermesUrlUpdateSignature(registryAddress, hermesAddress, url, lastNonce.Uint64(), ownerPk)
	if err != nil {
		return fmt.Errorf("failed to get hermes url update signature: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	opts.Context = ctx
	_, err = registry.UpdateHermesURL(opts, hermesAddress, []byte(url), signature)
	if err != nil {
		return fmt.Errorf("failed to update hermes url: %w", err)
	}
	return nil
}

func getHermesUrlUpdateSignature(registry, hermesId common.Address, url string, lastNonce uint64, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	message := []byte{}
	message = append(message, registry.Bytes()...)
	message = append(message, hermesId.Bytes()...)
	message = append(message, []byte(url)...)
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, lastNonce)
	message = append(message, pc.Pad(b, 32)...)

	hash := crypto.Keccak256(message)
	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		return nil, err
	}
	if err := pc.ReformatSignatureVForBC(signature); err != nil {
		return nil, fmt.Errorf("failed to reformat signature: %w", err)
	}
	return signature, nil
}
