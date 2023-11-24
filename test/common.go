package test

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mysteriumnetwork/payments/bindings"
)

func GetKeyPair(privateKeyString string) (common.Address, *ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("error parsing private key: %w", err)
	}
	address, err := GetAddress(privateKey)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("error getting address: %w", err)
	}
	return address, privateKey, nil
}

func GenerateKeyPair() (common.Address, *ecdsa.PrivateKey, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to generate private key: %w", err)
	}
	address, err := GetAddress(privateKey)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("error getting address: %w", err)
	}
	return address, privateKey, nil
}

func GetAddress(privateKey *ecdsa.PrivateKey) (common.Address, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, fmt.Errorf("error casting public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address, nil
}

func GetTransactOpts(sender common.Address, privateKey *ecdsa.PrivateKey, chainID *big.Int) *bind.TransactOpts {
	return &bind.TransactOpts{
		From: sender,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
		},
	}
}

type BaseContractAddresses struct {
	TokenContractAddresses
	RegistryAddress              common.Address
	ChannelImplementationAddress common.Address
	HermesImplementationAddress  common.Address
}

type RegistryOpts struct {
	DexAddress         common.Address
	MinimalHermesStake *big.Int
	ParentRegistry     *common.Address
}

func DeployBaseSmartContracts(opts *bind.TransactOpts, backend bind.ContractBackend, timeout time.Duration, registryOpts RegistryOpts) (BaseContractAddresses, error) {
	registryAddress, err := DeployRegistry(opts, backend, timeout)
	if err != nil {
		return BaseContractAddresses{}, fmt.Errorf("failed to deploy registry: %w", err)
	}
	tokenAddresses, err := DeployTokenV2WithDependencies(opts, backend, timeout)
	if err != nil {
		return BaseContractAddresses{}, fmt.Errorf("failed to deploy token: %w", err)
	}
	channelImplementationAddress, err := DeployChannelImplementation(opts, backend, timeout)
	if err != nil {
		return BaseContractAddresses{}, fmt.Errorf("failed to deploy registry: %w", err)
	}
	hermesImplementationAddress, err := DeployHermesImplementation(opts, backend, timeout)
	if err != nil {
		return BaseContractAddresses{}, fmt.Errorf("failed to deploy registry: %w", err)
	}
	err = InitializeRegistry(opts, registryAddress, backend, 10*time.Second, tokenAddresses.TokenV2Address, registryOpts.DexAddress, registryOpts.MinimalHermesStake, channelImplementationAddress, hermesImplementationAddress, registryOpts.ParentRegistry)
	if err != nil {
		return BaseContractAddresses{}, fmt.Errorf("failed to initialize registry: %w", err)
	}
	return BaseContractAddresses{
		TokenContractAddresses:       tokenAddresses,
		RegistryAddress:              registryAddress,
		ChannelImplementationAddress: channelImplementationAddress,
		HermesImplementationAddress:  hermesImplementationAddress,
	}, nil
}

type HermesWithDependenciesContractAddresses struct {
	BaseContractAddresses
	HermesAddress common.Address
}

type RegisterHermesOpts struct {
	Operator        common.Address
	HermesStake     *big.Int
	HermesFee       uint16
	MinChannelStake *big.Int
	MaxChannelStake *big.Int
	Url             string
}

func DeployHermesWithDependencies(opts *bind.TransactOpts, backend bind.ContractBackend, timeout time.Duration, registryOpts RegistryOpts, registerHermesOpts RegisterHermesOpts) (HermesWithDependenciesContractAddresses, error) {
	baseContractAddresses, err := DeployBaseSmartContracts(opts, backend, timeout, registryOpts)
	if err != nil {
		return HermesWithDependenciesContractAddresses{}, fmt.Errorf("failed to deploy base smart contracts: %w", err)
	}

	err = MintTokens(opts, baseContractAddresses.TokenV2Address, backend, timeout, opts.From, registerHermesOpts.HermesStake)
	if err != nil {
		return HermesWithDependenciesContractAddresses{}, fmt.Errorf("failed to mint tokens: %w", err)
	}
	err = RegisterHermes(opts, baseContractAddresses.RegistryAddress, backend, timeout,
		registerHermesOpts.Operator,
		registerHermesOpts.HermesStake,
		registerHermesOpts.HermesFee,
		registerHermesOpts.MinChannelStake,
		registerHermesOpts.MaxChannelStake,
		registerHermesOpts.Url)
	if err != nil {
		return HermesWithDependenciesContractAddresses{}, fmt.Errorf("failed to register hermes: %w", err)
	}

	registryCaller, err := bindings.NewRegistry(baseContractAddresses.RegistryAddress, backend)
	if err != nil {
		return HermesWithDependenciesContractAddresses{}, fmt.Errorf("failed to create registry caller: %w", err)
	}
	lastImplVersion, err := registryCaller.GetLastImplVer(nil)
	if err != nil {
		return HermesWithDependenciesContractAddresses{}, fmt.Errorf("failed to get last impl version: %w", err)
	}
	hermesAddress, err := registryCaller.GetHermesAddress(nil, registerHermesOpts.Operator, lastImplVersion)
	if err != nil {
		return HermesWithDependenciesContractAddresses{}, fmt.Errorf("failed to get hermes address: %w", err)
	}

	return HermesWithDependenciesContractAddresses{
		BaseContractAddresses: baseContractAddresses,
		HermesAddress:         hermesAddress,
	}, nil
}

func DeployHermes(opts *bind.TransactOpts, backend bind.ContractBackend, timeout time.Duration, baseContractAddresses BaseContractAddresses, registerHermesOpts RegisterHermesOpts) (HermesWithDependenciesContractAddresses, error) {
	err := MintTokens(opts, baseContractAddresses.TokenV2Address, backend, timeout, opts.From, registerHermesOpts.HermesStake)
	if err != nil {
		return HermesWithDependenciesContractAddresses{}, fmt.Errorf("failed to mint tokens: %w", err)
	}
	err = RegisterHermes(opts, baseContractAddresses.RegistryAddress, backend, timeout,
		registerHermesOpts.Operator,
		registerHermesOpts.HermesStake,
		registerHermesOpts.HermesFee,
		registerHermesOpts.MinChannelStake,
		registerHermesOpts.MaxChannelStake,
		registerHermesOpts.Url)
	if err != nil {
		return HermesWithDependenciesContractAddresses{}, fmt.Errorf("failed to register hermes: %w", err)
	}

	registryCaller, err := bindings.NewRegistry(baseContractAddresses.RegistryAddress, backend)
	if err != nil {
		return HermesWithDependenciesContractAddresses{}, fmt.Errorf("failed to create registry caller: %w", err)
	}
	lastImplVersion, err := registryCaller.GetLastImplVer(nil)
	if err != nil {
		return HermesWithDependenciesContractAddresses{}, fmt.Errorf("failed to get last impl version: %w", err)
	}
	hermesAddress, err := registryCaller.GetHermesAddress(nil, registerHermesOpts.Operator, lastImplVersion)
	if err != nil {
		return HermesWithDependenciesContractAddresses{}, fmt.Errorf("failed to get hermes address: %w", err)
	}

	return HermesWithDependenciesContractAddresses{
		BaseContractAddresses: baseContractAddresses,
		HermesAddress:         hermesAddress,
	}, nil
}

func TransferEth(opts *bind.TransactOpts, backend bind.ContractBackend, chainID *big.Int, to common.Address, value *big.Int, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	nonce, err := backend.PendingNonceAt(ctx, opts.From)
	if err != nil {
		return fmt.Errorf("failed to get nonce: %w", err)
	}
	gasFeeCap, err := backend.SuggestGasPrice(ctx)
	if err != nil {
		return fmt.Errorf("failed to get gas price: %w", err)
	}
	tx := types.NewTx(&types.DynamicFeeTx{
		Nonce:     nonce,
		ChainID:   chainID,
		To:        &to,
		Value:     value,
		Gas:       21000,
		GasFeeCap: gasFeeCap,
	})
	signedTx, err := opts.Signer(opts.From, tx)
	if err != nil {
		return fmt.Errorf("failed to sign transaction: %w", err)
	}
	err = backend.SendTransaction(ctx, signedTx)
	if err != nil {
		return fmt.Errorf("failed to send transaction: %w", err)
	}
	return nil
}
