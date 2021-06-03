package client

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/mysteriumnetwork/payments/bindings"
)

type BC interface {
	CalculateHermesFee(hermesAddress common.Address, value *big.Int) (*big.Int, error)
	IsHermesActive(hermesID common.Address) (bool, error)
	IsHermesRegistered(registryAddress, acccountantID common.Address) (bool, error)

	GetHermesFee(hermesAddress common.Address) (uint16, error)
	GetHermessAvailableBalance(hermesAddress common.Address) (*big.Int, error)
	GetHermesURL(registryID, hermesID common.Address) (string, error)
	GetHermes(registryID, hermesID common.Address) (Hermes, error)
	GetHermesOperator(hermesID common.Address) (common.Address, error)
	GetProviderChannel(hermesAddress common.Address, addressToCheck common.Address, pending bool) (ProviderChannel, error)
	GetMystBalance(mystSCAddress, address common.Address) (*big.Int, error)
	GetConsumerChannelsHermes(channelAddress common.Address) (ConsumersHermes, error)
	GetConsumerChannelOperator(channelAddress common.Address) (common.Address, error)
	GetProviderChannelByID(acc common.Address, chID []byte) (ProviderChannel, error)
	GetConsumerChannel(addr common.Address, mystSCAddress common.Address) (ConsumerChannel, error)
	GetEthBalance(address common.Address) (*big.Int, error)
	GetStakeThresholds(hermesID common.Address) (min, max *big.Int, err error)
	GetBeneficiary(registryAddress, identity common.Address) (common.Address, error)
	GetLastRegistryNonce(registry common.Address) (*big.Int, error)
	GetChannelImplementationByVersion(registryID common.Address, version *big.Int) (common.Address, error)

	IsRegisteredAsProvider(hermesAddress, registryAddress, addressToCheck common.Address) (bool, error)
	IsRegistered(registryAddress, addressToCheck common.Address) (bool, error)
	RegisterIdentity(rr RegistrationRequest) (*types.Transaction, error)
	IncreaseProviderStake(req ProviderStakeIncreaseRequest) (*types.Transaction, error)

	SettleAndRebalance(req SettleAndRebalanceRequest) (*types.Transaction, error)
	SettleWithBeneficiary(req SettleWithBeneficiaryRequest) (*types.Transaction, error)
	SettlePromise(req SettleRequest) (*types.Transaction, error)
	DecreaseProviderStake(req DecreaseProviderStakeRequest) (*types.Transaction, error)
	SettleIntoStake(req SettleIntoStakeRequest) (*types.Transaction, error)

	SubscribeToPromiseSettledEvent(providerID, hermesID common.Address) (sink chan *bindings.HermesImplementationPromiseSettled, cancel func(), err error)
	SubscribeToConsumerBalanceEvent(channel, mystSCAddress common.Address, timeout time.Duration) (chan *bindings.MystTokenTransfer, func(), error)
	SubscribeToIdentityRegistrationEvents(registryAddress common.Address) (sink chan *bindings.RegistryRegisteredIdentity, cancel func(), err error)
	SubscribeToConsumerChannelBalanceUpdate(mystSCAddress common.Address, channelAddresses []common.Address) (sink chan *bindings.MystTokenTransfer, cancel func(), err error)
	SubscribeToPromiseSettledEventByChannelID(hermesID common.Address, providerAddresses [][32]byte) (sink chan *bindings.HermesImplementationPromiseSettled, cancel func(), err error)
	SubscribeToMystTokenTransfers(mystSCAddress common.Address) (chan *bindings.MystTokenTransfer, func(), error)
	FilterLogs(q ethereum.FilterQuery) ([]types.Log, error)

	NetworkID() (*big.Int, error)
	SuggestGasPrice() (*big.Int, error)
	HeaderByNumber(number *big.Int) (*types.Header, error)

	TransferMyst(req TransferRequest) (tx *types.Transaction, err error)
	TransferEth(etr EthTransferRequest) (*types.Transaction, error)

	SendTransaction(tx *types.Transaction) error
	TransactionReceipt(hash common.Hash) (*types.Receipt, error)
	TransactionByHash(hash common.Hash) (*types.Transaction, bool, error)

	RewarderTotalPayoutsFor(rewarderAddress common.Address, payoutsFor common.Address) (*big.Int, error)
	RewarderAirDrop(req RewarderAirDrop) (*types.Transaction, error)
	RewarderUpdateRoot(req RewarderUpdateRoot) (*types.Transaction, error)
	RewarderTotalClaimed(rewarderAddress common.Address) (*big.Int, error)
	CustodyTransferTokens(req CustodyTokensTransfer) (*types.Transaction, error)

	PayAndSettle(psr PayAndSettleRequest) (*types.Transaction, error)
	IsChannelOpened(registryID, identity, hermesID common.Address) (bool, error)
}
