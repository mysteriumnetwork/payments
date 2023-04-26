package test

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mysteriumnetwork/payments/bindings"
	pc "github.com/mysteriumnetwork/payments/crypto"
	"github.com/mysteriumnetwork/payments/registration"
)

func DeployRegistry(opts *bind.TransactOpts, backend bind.ContractBackend, timeout time.Duration) (common.Address, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	opts.Context = ctx
	address, _, _, err := bindings.DeployRegistry(opts, backend)
	if err != nil {
		return address, fmt.Errorf("failed to deploy registry: %w", err)
	}
	return address, nil
}

func InitializeRegistry(opts *bind.TransactOpts, registryAddress common.Address, transactor bind.ContractTransactor, timeout time.Duration, tokenAddress, dexAddress common.Address, minimalHermesStake *big.Int, channelImplementation, hermesImplementation common.Address, parentRegistry *common.Address) error {
	registryTransactor, err := bindings.NewRegistryTransactor(registryAddress, transactor)
	if err != nil {
		return fmt.Errorf("failed to initialize registry transactor: %w", err)
	}
	pr := common.Address{}
	if parentRegistry != nil {
		pr = *parentRegistry
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	opts.Context = ctx
	_, err = registryTransactor.Initialize(opts, tokenAddress, dexAddress, minimalHermesStake, channelImplementation, hermesImplementation, pr)
	if err != nil {
		return fmt.Errorf("failed to initialize registry: %w", err)
	}
	return nil
}

func DeployAndInitializeRegistry(opts *bind.TransactOpts, backend bind.ContractBackend, tokenAddress, dexAddress common.Address, minimalHermesStake *big.Int, channelImplementation, hermesImplementation common.Address, parentRegistry *common.Address, timeout time.Duration) (common.Address, error) {
	registryAddress, err := DeployRegistry(opts, backend, timeout)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to deploy registry: %w", err)
	}
	err = InitializeRegistry(opts, registryAddress, backend, timeout, tokenAddress, dexAddress, minimalHermesStake, channelImplementation, hermesImplementation, parentRegistry)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to initialize registry: %w", err)
	}
	return registryAddress, nil
}

func RegisterHermes(opts *bind.TransactOpts, registryAddress common.Address, backend bind.ContractBackend, timeout time.Duration, operator common.Address, hermesStake *big.Int, hermesFee uint16, minChannelStake *big.Int, maxChannelStake *big.Int, url string) error {
	registryCaller, err := bindings.NewRegistry(registryAddress, backend)
	if err != nil {
		return fmt.Errorf("failed to initialize registry transactor: %w", err)
	}
	tokenAddress, err := registryCaller.Token(nil)
	if err != nil {
		return fmt.Errorf("failed to get token address: %w", err)
	}
	tokenCaller, err := bindings.NewMystToken(tokenAddress, backend)
	if err != nil {
		return fmt.Errorf("failed to make token caller: %w", err)
	}
	allowance, err := tokenCaller.Allowance(nil, opts.From, registryAddress)
	if err != nil {
		return fmt.Errorf("failed to get allowance: %w", err)
	}
	if allowance.Cmp(hermesStake) < 0 {
		ctxAp, cancelAp := context.WithTimeout(context.Background(), timeout)
		defer cancelAp()
		opts.Context = ctxAp
		_, err = tokenCaller.Approve(opts, registryAddress, hermesStake)
		if err != nil {
			return fmt.Errorf("failed to approve token allowance: %w", err)
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	opts.Context = ctx
	_, err = registryCaller.RegisterHermes(opts, operator, hermesStake, hermesFee, minChannelStake, maxChannelStake, []byte(url))
	if err != nil {
		return fmt.Errorf("failed to register hermes: %w", err)
	}
	return nil
}

func GenerateRegisterIdentitySignature(privateKey *ecdsa.PrivateKey, registryAddress, hermesAddress common.Address, stakeAmount *big.Int, transactorFee *big.Int, beneficiary common.Address, chainID int64) ([]byte, error) {
	req := registration.Request{
		RegistryAddress: strings.ToLower(registryAddress.Hex()),
		HermesID:        strings.ToLower(hermesAddress.Hex()),
		Stake:           stakeAmount,
		Fee:             transactorFee,
		Beneficiary:     strings.ToLower(beneficiary.Hex()),
		ChainID:         chainID,
	}
	signature, err := crypto.Sign(crypto.Keccak256(req.GetMessage()), privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign register identity request: %w", err)
	}
	err = pc.ReformatSignatureVForBC(signature)
	if err != nil {
		return nil, fmt.Errorf("failed to reformat signature: %w", err)
	}
	return signature, nil
}

func RegisterIdentityWithSignature(opts *bind.TransactOpts, registryAddress common.Address, transactor bind.ContractTransactor, timeout time.Duration, hermesAddress common.Address, stakeAmount *big.Int, transactorFee *big.Int, beneficiary common.Address, signature []byte) error {
	registryTransactor, err := bindings.NewRegistryTransactor(registryAddress, transactor)
	if err != nil {
		return fmt.Errorf("failed to initialize registry transactor: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	opts.Context = ctx
	_, err = registryTransactor.RegisterIdentity(opts, hermesAddress, stakeAmount, transactorFee, beneficiary, signature)
	if err != nil {
		return fmt.Errorf("failed to register identity: %w", err)
	}
	return nil
}

func RegisterIdentityWithPrivateKey(opts *bind.TransactOpts, registryAddress common.Address, transactor bind.ContractTransactor, timeout time.Duration, hermesAddress common.Address, stakeAmount *big.Int, transactorFee *big.Int, beneficiary common.Address, privateKey *ecdsa.PrivateKey, chainID int64) error {
	signature, err := GenerateRegisterIdentitySignature(privateKey, registryAddress, hermesAddress, stakeAmount, transactorFee, beneficiary, chainID)
	if err != nil {
		return fmt.Errorf("failed to generate register identity signature: %w", err)
	}
	registryTransactor, err := bindings.NewRegistryTransactor(registryAddress, transactor)
	if err != nil {
		return fmt.Errorf("failed to initialize registry transactor: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	opts.Context = ctx
	_, err = registryTransactor.RegisterIdentity(opts, hermesAddress, stakeAmount, transactorFee, beneficiary, signature)
	if err != nil {
		return fmt.Errorf("failed to register identity: %w", err)
	}
	return nil
}
