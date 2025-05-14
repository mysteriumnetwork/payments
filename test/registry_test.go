package test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mysteriumnetwork/payments/v3/bindings"
	"github.com/stretchr/testify/assert"
)

func TestDeployRegistry(t *testing.T) {
	address, privateKey, err := GetKeyPair(privateKey0)
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	chainID, err := client.ChainID(ctx)
	assert.NoError(t, err)

	registryAddress, err := DeployRegistry(GetTransactOpts(address, privateKey, chainID), client, 10*time.Second)
	assert.NoError(t, err)
	assert.NotEqual(t, (common.Address{}).Hex(), registryAddress.Hex())

	registryCaller, err := bindings.NewRegistry(registryAddress, client)
	assert.NoError(t, err)

	owner, err := registryCaller.Owner(nil)
	assert.NoError(t, err)
	assert.Equal(t, (common.Address{}).Hex(), owner.Hex())
}

func TestInitializeRegistry(t *testing.T) {
	address, privateKey, err := GetKeyPair(privateKey0)
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	chainID, err := client.ChainID(ctx)
	assert.NoError(t, err)

	topts := GetTransactOpts(address, privateKey, chainID)
	registryAddress, err := DeployRegistry(topts, client, 10*time.Second)
	assert.NoError(t, err)
	assert.NotEqual(t, (common.Address{}).Hex(), registryAddress.Hex())

	registryCaller, err := bindings.NewRegistry(registryAddress, client)
	assert.NoError(t, err)

	owner, err := registryCaller.Owner(nil)
	assert.NoError(t, err)
	assert.Equal(t, (common.Address{}).Hex(), owner.Hex())

	tokenAddresses, err := DeployTokenV2WithDependencies(GetTransactOpts(address, privateKey, chainID), client, 10*time.Second)
	assert.NoError(t, err)
	channelImplementationAddress, err := DeployChannelImplementation(GetTransactOpts(address, privateKey, chainID), client, 10*time.Second)
	assert.NoError(t, err)
	hermesImplementationAddress, err := DeployHermesImplementation(GetTransactOpts(address, privateKey, chainID), client, 10*time.Second)
	assert.NoError(t, err)

	mockAddress1 := common.HexToAddress("0x1")
	minimalHermesStake := big.NewInt(100)

	err = InitializeRegistry(topts, registryAddress, client, 10*time.Second, tokenAddresses.TokenV2Address, mockAddress1, minimalHermesStake, channelImplementationAddress, hermesImplementationAddress, nil)
	assert.NoError(t, err)

	owner, err = registryCaller.Owner(nil)
	assert.NoError(t, err)
	assert.Equal(t, address.Hex(), owner.Hex())

	tokenAddress, err := registryCaller.Token(nil)
	assert.NoError(t, err)
	assert.Equal(t, tokenAddresses.TokenV2Address.Hex(), tokenAddress.Hex())

	dexAddress, err := registryCaller.Dex(nil)
	assert.NoError(t, err)
	assert.Equal(t, mockAddress1.Hex(), dexAddress.Hex())

	mhs, err := registryCaller.MinimalHermesStake(nil)
	assert.NoError(t, err)
	assert.Equal(t, minimalHermesStake.String(), mhs.String())

	cia, err := registryCaller.GetChannelImplementation0(nil)
	assert.NoError(t, err)
	assert.Equal(t, channelImplementationAddress.Hex(), cia.Hex())

	hia, err := registryCaller.GetHermesImplementation0(nil)
	assert.NoError(t, err)
	assert.Equal(t, hermesImplementationAddress.Hex(), hia.Hex())
}

func TestRegisterHermes(t *testing.T) {
	address, privateKey, err := GetKeyPair(privateKey0)
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	chainID, err := client.ChainID(ctx)
	assert.NoError(t, err)

	topts := GetTransactOpts(address, privateKey, chainID)
	baseSmartContracts, err := DeployBaseSmartContracts(topts, client, 10*time.Second, RegistryOpts{
		DexAddress:         common.HexToAddress("0x1"),
		MinimalHermesStake: big.NewInt(100),
		ParentRegistry:     nil,
	})
	assert.NoError(t, err)
	hermesStake := big.NewInt(100)

	err = MintTokens(topts, baseSmartContracts.TokenV2Address, client, 10*time.Second, address, hermesStake)
	assert.NoError(t, err)

	operator := common.HexToAddress("0x2")
	hermesFee := uint16(200)
	minChannelStake := big.NewInt(456)
	maxChannelStake := big.NewInt(789)
	url := "https://hermes.mysterium.network"
	err = RegisterHermes(topts, baseSmartContracts.RegistryAddress, client, 10*time.Second, operator, hermesStake, hermesFee, minChannelStake, maxChannelStake, url)
	assert.NoError(t, err)

	registryCaller, err := bindings.NewRegistry(baseSmartContracts.RegistryAddress, client)
	assert.NoError(t, err)

	lastImplVersion, err := registryCaller.GetLastImplVer(nil)
	assert.NoError(t, err)
	hermesAddress, err := registryCaller.GetHermesAddress(nil, operator, lastImplVersion)
	assert.NoError(t, err)

	hermesCaller, err := bindings.NewHermesImplementation(hermesAddress, client)
	assert.NoError(t, err)

	hermesOperator, err := hermesCaller.GetOperator(nil)
	assert.NoError(t, err)
	assert.Equal(t, operator.Hex(), hermesOperator.Hex())

	hermesStakeValue, err := hermesCaller.GetHermesStake(nil)
	assert.NoError(t, err)
	assert.Equal(t, hermesStake.String(), hermesStakeValue.String())

	hermesFeeValue, err := hermesCaller.GetActiveFee(nil)
	assert.NoError(t, err)
	assert.Equal(t, int64(hermesFee), hermesFeeValue.Int64())

	hermesMinChannelStake, hermesMaxChannelStake, err := hermesCaller.GetStakeThresholds(nil)
	assert.NoError(t, err)
	assert.Equal(t, minChannelStake.String(), hermesMinChannelStake.String())
	assert.Equal(t, maxChannelStake.String(), hermesMaxChannelStake.String())

	hermesURL, err := registryCaller.GetHermesURL(nil, hermesAddress)
	assert.NoError(t, err)
	assert.Equal(t, url, string(hermesURL))
}

func TestRegisterIdentity(t *testing.T) {
	address, privateKey, err := GetKeyPair(privateKey0)
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	chainID, err := client.ChainID(ctx)
	assert.NoError(t, err)

	topts := GetTransactOpts(address, privateKey, chainID)
	hermesSmartContracts, err := DeployHermesWithDependencies(topts, client, 10*time.Second, RegistryOpts{
		DexAddress:         common.HexToAddress("0x1"),
		MinimalHermesStake: big.NewInt(100),
		ParentRegistry:     nil,
	}, RegisterHermesOpts{
		Operator:        topts.From,
		HermesStake:     big.NewInt(100),
		HermesFee:       200,
		MinChannelStake: big.NewInt(0),
		MaxChannelStake: big.NewInt(1000),
		Url:             "https://hermes.mysterium.network",
	})
	assert.NoError(t, err)

	identityAddress, identityPrivateKey, err := GetKeyPair(privateKey0)
	assert.NoError(t, err)
	beneficiary := common.HexToAddress("0x5")

	channelStake := big.NewInt(0)
	err = RegisterIdentityWithPrivateKey(topts, hermesSmartContracts.RegistryAddress, client, 10*time.Second, hermesSmartContracts.HermesAddress, channelStake, big.NewInt(0), beneficiary, identityPrivateKey, chainID.Int64())
	assert.NoError(t, err)

	registryCaller, err := bindings.NewRegistry(hermesSmartContracts.RegistryAddress, client)
	assert.NoError(t, err)

	channelAddress, err := registryCaller.GetChannelAddress(nil, identityAddress, hermesSmartContracts.HermesAddress)
	assert.NoError(t, err)

	channelCaller, err := bindings.NewChannelImplementation(channelAddress, client)
	assert.NoError(t, err)

	channelHermes, err := channelCaller.Hermes(nil)
	assert.NoError(t, err)
	assert.Equal(t, hermesSmartContracts.HermesAddress.Hex(), channelHermes.ContractAddress.Hex())
	assert.Equal(t, identityAddress.Hex(), channelHermes.Operator.Hex())
	assert.Equal(t, "0", channelHermes.Settled.String())

	channelOperator, err := channelCaller.Operator(nil)
	assert.NoError(t, err)
	assert.Equal(t, identityAddress.Hex(), channelOperator.Hex())

	channelOwner, err := channelCaller.Owner(nil)
	assert.NoError(t, err)
	assert.Equal(t, identityAddress.Hex(), channelOwner.Hex())

	init, err := channelCaller.IsInitialized(nil)
	assert.NoError(t, err)
	assert.True(t, init)

	beneficiaryValue, err := registryCaller.GetBeneficiary(nil, identityAddress)
	assert.NoError(t, err)
	assert.Equal(t, beneficiary.Hex(), beneficiaryValue.Hex())
}
