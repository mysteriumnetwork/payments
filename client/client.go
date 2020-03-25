package client

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/mysteriumnetwork/payments/bindings"
	"github.com/mysteriumnetwork/payments/crypto"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// DefaultBackoff is the default backoff for the client
var DefaultBackoff = time.Second * 3

// ProviderChannel represents the provider channel
type ProviderChannel struct {
	Beneficiary   common.Address
	Balance       *big.Int
	Settled       *big.Int
	Loan          *big.Int
	LastUsedNonce *big.Int
	Timelock      *big.Int
}

// Blockchain contains all the useful blockchain utilities for the payment off chain messaging
type Blockchain struct {
	client    *ethclient.Client
	bcTimeout time.Duration
}

// NewBlockchain returns a new instance of blockchain
func NewBlockchain(c *ethclient.Client, timeout time.Duration) *Blockchain {
	return &Blockchain{
		client:    c,
		bcTimeout: timeout,
	}
}

// GetAccountantFee fetches the accountant fee from blockchain
func (bc *Blockchain) GetAccountantFee(accountantAddress common.Address) (uint16, error) {
	caller, err := bindings.NewAccountantImplementationCaller(accountantAddress, bc.client)
	if err != nil {
		return 0, errors.Wrap(err, "could not create accountant implementation caller")
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	res, err := caller.LastFee(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return 0, errors.Wrap(err, "could not get accountant fee")
	}

	return res.Value, err
}

// CalculateAccountantFee calls blockchain for calculation of accountant fee
func (bc *Blockchain) CalculateAccountantFee(accountantAddress common.Address, value uint64) (*big.Int, error) {
	caller, err := bindings.NewAccountantImplementationCaller(accountantAddress, bc.client)
	if err != nil {
		return nil, errors.Wrap(err, "could not create accountant implementation caller")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return caller.CalculateAccountantFee(&bind.CallOpts{
		Context: ctx,
	}, new(big.Int).SetUint64(value))
}

// IsRegisteredAsProvider checks if the provider is registered with the accountant properly
func (bc *Blockchain) IsRegisteredAsProvider(accountantAddress, registryAddress, addressToCheck common.Address) (bool, error) {
	registered, err := bc.IsRegistered(registryAddress, addressToCheck)
	if err != nil {
		return false, errors.Wrap(err, "could not check registration status")
	}

	if !registered {
		return false, nil
	}

	res, err := bc.getProviderChannelLoan(accountantAddress, addressToCheck)
	if err != nil {
		return false, errors.Wrap(err, "could not get provider channel loan amount")
	}

	return res.Cmp(big.NewInt(0)) == 1, nil
}

// SubscribeToMystTokenTransfers subscribes to myst token transfers
func (bc *Blockchain) SubscribeToMystTokenTransfers(mystSCAddress common.Address) (chan *bindings.MystTokenTransfer, func(), error) {
	sink := make(chan *bindings.MystTokenTransfer)
	mtc, err := bindings.NewMystTokenFilterer(mystSCAddress, bc.client)
	if err != nil {
		return sink, nil, err
	}
	sub := event.Resubscribe(DefaultBackoff, func(ctx context.Context) (event.Subscription, error) {
		return mtc.WatchTransfer(&bind.WatchOpts{
			Context: ctx,
		}, sink, []common.Address{}, []common.Address{})
	})
	go func() {
		subErr := <-sub.Err()
		if subErr != nil {
			log.Error().Err(err).Msg("subscription error")
		}
		close(sink)
	}()

	return sink, sub.Unsubscribe, nil
}

// SubscribeToConsumerBalanceEvent subscribes to balance change events in blockchain
func (bc *Blockchain) SubscribeToConsumerBalanceEvent(channel, mystSCAddress common.Address, timeout time.Duration) (chan *bindings.MystTokenTransfer, func(), error) {
	sink := make(chan *bindings.MystTokenTransfer)
	mtc, err := bindings.NewMystTokenFilterer(mystSCAddress, bc.client)
	if err != nil {
		return sink, nil, err
	}
	sub := event.Resubscribe(DefaultBackoff, func(ctx context.Context) (event.Subscription, error) {
		return mtc.WatchTransfer(&bind.WatchOpts{
			Context: ctx,
		}, sink, []common.Address{}, []common.Address{channel})
	})

	go func() {
		select {
		case <-time.After(timeout):
			sub.Unsubscribe()
		}
	}()

	go func() {
		subErr := <-sub.Err()
		if subErr != nil {
			log.Error().Err(err).Msg("subscription error")
		}
		close(sink)
	}()
	return sink, sub.Unsubscribe, nil
}

// GetProviderChannel returns the provider channel
func (bc *Blockchain) GetProviderChannel(accountantAddress common.Address, addressToCheck common.Address) (ProviderChannel, error) {
	addressBytes, err := bc.getProviderChannelAddressBytes(accountantAddress, addressToCheck)
	if err != nil {
		return ProviderChannel{}, errors.Wrap(err, "could not calculate provider channel address")
	}
	caller, err := bindings.NewAccountantImplementationCaller(accountantAddress, bc.client)
	if err != nil {
		return ProviderChannel{}, errors.Wrap(err, "could not create accountant caller")
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	ch, err := caller.Channels(&bind.CallOpts{
		Context: ctx,
	}, addressBytes)
	return ch, errors.Wrap(err, "could not get provider channel from bc")
}

func (bc *Blockchain) getProviderChannelLoan(accountantAddress common.Address, addressToCheck common.Address) (*big.Int, error) {
	ch, err := bc.GetProviderChannel(accountantAddress, addressToCheck)
	return ch.Loan, errors.Wrap(err, "could not get provider channel from bc")
}

func (bc *Blockchain) getProviderChannelAddressBytes(accountantAddress, addressToCheck common.Address) ([32]byte, error) {
	addressBytes := [32]byte{}

	addr, err := crypto.GenerateProviderChannelID(addressToCheck.Hex(), accountantAddress.Hex())
	if err != nil {
		return addressBytes, errors.Wrap(err, "could not generate channel address")
	}

	copy(addressBytes[:], crypto.Pad(common.Hex2Bytes(strings.TrimPrefix(addr, "0x")), 32))

	return addressBytes, nil
}

// SubscribeToPromiseSettledEvent subscribes to promise settled events
func (bc *Blockchain) SubscribeToPromiseSettledEvent(providerID, accountantID common.Address) (sink chan *bindings.AccountantImplementationPromiseSettled, cancel func(), err error) {
	addr, err := bc.getProviderChannelAddressBytes(accountantID, providerID)
	if err != nil {
		return sink, cancel, errors.Wrap(err, "could not get provider channel address")
	}
	return bc.SubscribeToPromiseSettledEventByChannelID(accountantID, [][32]byte{addr})
}

// IsRegistered checks wether the given identity is registered or not
func (bc *Blockchain) IsRegistered(registryAddress, addressToCheck common.Address) (bool, error) {
	caller, err := bindings.NewRegistryCaller(registryAddress, bc.client)
	if err != nil {
		return false, errors.Wrap(err, "could not create registry caller")
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	res, err := caller.IsRegistered(&bind.CallOpts{
		Context: ctx,
	}, addressToCheck)
	return res, errors.Wrap(err, "could not check registration status")
}

// GetMystBalance returns myst balance
func (bc *Blockchain) GetMystBalance(mystAddress, identity common.Address) (*big.Int, error) {
	c, err := bindings.NewMystTokenCaller(mystAddress, bc.client)
	if err != nil {
		return nil, err
	}
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, bc.bcTimeout)
	defer cancel()
	return c.BalanceOf(&bind.CallOpts{
		Context: ctx,
	}, identity)
}

// GetRegistrationFee returns fee required by registry
func (bc *Blockchain) GetRegistrationFee(registryAddress common.Address) (*big.Int, error) {
	// TODO to reduce amount of blockchain calls, it could get registration fee from cache (updated once in a day)
	c, err := bindings.NewRegistryCaller(registryAddress, bc.client)
	if err != nil {
		return nil, errors.Wrap(err, "could not get registration fee")
	}
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, bc.bcTimeout)
	defer cancel()
	res, err := c.RegistrationFee(&bind.CallOpts{
		Context: ctx,
	})
	return res, errors.Wrap(err, "could not get registration fee")
}

// RegistrationRequest contains all the parameters for the registration request
type RegistrationRequest struct {
	WriteRequest
	AccountantID    common.Address
	Loan            *big.Int
	TransactorFee   *big.Int
	Beneficiary     common.Address
	Signature       []byte
	RegistryAddress common.Address
}

// WriteRequest contains the required params for a write request
type WriteRequest struct {
	Identity common.Address
	Signer   bind.SignerFn
	GasLimit uint64
	GasPrice *big.Int
}

// GetGasLimit returns the gas limit
func (wr WriteRequest) GetGasLimit() uint64 {
	return wr.GasLimit
}

// EstimateGas exposes the clients internal estimate gas
func (bc *Blockchain) EstimateGas(msg ethereum.CallMsg) (uint64, error) {
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, bc.bcTimeout)
	defer cancel()
	return bc.client.EstimateGas(ctx, msg)
}

// RegisterIdentity registers the given identity on blockchain
func (bc *Blockchain) RegisterIdentity(rr RegistrationRequest) (*types.Transaction, error) {
	transactor, err := bindings.NewRegistryTransactor(rr.RegistryAddress, bc.client)
	if err != nil {
		return nil, err
	}
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, bc.bcTimeout)
	defer cancel()

	nonce, err := bc.getNonce(rr.Identity)
	if err != nil {
		return nil, errors.Wrap(err, "could not get nonce")
	}
	tx, err := transactor.RegisterIdentity(&bind.TransactOpts{
		From:     rr.Identity,
		Signer:   rr.Signer,
		Context:  ctx,
		GasLimit: rr.GasLimit,
		GasPrice: rr.GasPrice,
		Nonce:    big.NewInt(0).SetUint64(nonce),
	},
		rr.AccountantID,
		rr.Loan,
		rr.TransactorFee,
		rr.Beneficiary,
		rr.Signature,
	)
	return tx, err
}

// TransferRequest contains all the parameters for a transfer request
type TransferRequest struct {
	MystAddress common.Address
	Recipient   common.Address
	Amount      *big.Int
	WriteRequest
}

// TransferMyst transfers myst
func (bc *Blockchain) TransferMyst(req TransferRequest) (tx *types.Transaction, err error) {
	transactor, err := bindings.NewMystTokenTransactor(req.MystAddress, bc.client)
	if err != nil {
		return tx, err
	}

	nonce, err := bc.getNonce(req.Identity)
	if err != nil {
		return nil, errors.Wrap(err, "could not get nonce")
	}

	return transactor.Transfer(&bind.TransactOpts{
		From:     req.Identity,
		Signer:   req.Signer,
		GasPrice: req.GasPrice,
		GasLimit: req.GasLimit,
		Nonce:    big.NewInt(0).SetUint64(nonce),
	}, req.Recipient, req.Amount)
}

// IsAccountantRegistered checks if given accountant is registered and returns true or false.
func (bc *Blockchain) IsAccountantRegistered(registryAddress, acccountantID common.Address) (bool, error) {
	caller, err := bindings.NewRegistryCaller(registryAddress, bc.client)
	if err != nil {
		return false, err
	}
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, bc.bcTimeout)
	defer cancel()
	return caller.IsAccountant(&bind.CallOpts{
		Pending: false,
		Context: ctx,
	}, acccountantID)
}

// GetAccountantOperator returns operator address of given accountant
func (bc *Blockchain) GetAccountantOperator(accountantID common.Address) (common.Address, error) {
	caller, err := bindings.NewAccountantImplementationCaller(accountantID, bc.client)
	if err != nil {
		return common.Address{}, err
	}

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, bc.bcTimeout)
	defer cancel()
	return caller.GetOperator(&bind.CallOpts{
		Pending: false,
		Context: ctx,
	})
}

// SettleAndRebalanceRequest represents all the parameters required for settle and rebalance
type SettleAndRebalanceRequest struct {
	WriteRequest
	AccountantID common.Address
	Promise      crypto.Promise
}

// SettleAndRebalance is settling given accountant issued promise
func (bc *Blockchain) SettleAndRebalance(req SettleAndRebalanceRequest) (*types.Transaction, error) {
	transactor, err := bindings.NewAccountantImplementationTransactor(req.AccountantID, bc.client)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	nonce, err := bc.getNonce(req.Identity)
	if err != nil {
		return nil, errors.Wrap(err, "could not get nonce")
	}

	return transactor.SettleAndRebalance(&bind.TransactOpts{
		From:     req.Identity,
		Signer:   req.Signer,
		Context:  ctx,
		GasLimit: req.GasLimit,
		GasPrice: req.GasPrice,
		Nonce:    big.NewInt(0).SetUint64(nonce),
	},
		toBytes32(req.Promise.ChannelID),
		big.NewInt(0).SetUint64(req.Promise.Amount),
		big.NewInt(0).SetUint64(req.Promise.Fee),
		toBytes32(req.Promise.R),
		req.Promise.Signature,
	)
}

func toBytes32(arr []byte) (res [32]byte) {
	copy(res[:], arr)
	return res
}

// GetProviderChannelByID returns the given provider channel information
func (bc *Blockchain) GetProviderChannelByID(acc common.Address, chID []byte) (ProviderChannel, error) {
	caller, err := bindings.NewAccountantImplementationCaller(acc, bc.client)
	if err != nil {
		return ProviderChannel{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	return caller.Channels(&bind.CallOpts{
		Pending: false,
		Context: ctx,
	}, toBytes32(chID))
}

// ConsumersAccountant represents the consumers accountant
type ConsumersAccountant struct {
	Operator        common.Address
	ContractAddress common.Address
	Settled         *big.Int
}

// GetConsumerChannelsAccountant returns the consumer channels accountant
func (bc *Blockchain) GetConsumerChannelsAccountant(channelAddress common.Address) (ConsumersAccountant, error) {
	c, err := bindings.NewChannelImplementationCaller(channelAddress, bc.client)
	if err != nil {
		return ConsumersAccountant{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	return c.Accountant(&bind.CallOpts{Context: ctx})
}

// GetConsumerChannelOperator returns the consumer channel operator/identity
func (bc *Blockchain) GetConsumerChannelOperator(channelAddress common.Address) (common.Address, error) {
	c, err := bindings.NewChannelImplementationCaller(channelAddress, bc.client)
	if err != nil {
		return common.Address{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	return c.Operator(&bind.CallOpts{Context: ctx})
}

// SubscribeToIdentityRegistrationEvents subscribes to identity registration events
func (bc *Blockchain) SubscribeToIdentityRegistrationEvents(registryAddress common.Address, accountantIDs []common.Address) (sink chan *bindings.RegistryRegisteredIdentity, cancel func(), err error) {
	filterer, err := bindings.NewRegistryFilterer(registryAddress, bc.client)
	if err != nil {
		return sink, cancel, errors.Wrap(err, "could not create registry filterer")
	}
	sink = make(chan *bindings.RegistryRegisteredIdentity)
	sub := event.Resubscribe(DefaultBackoff, func(ctx context.Context) (event.Subscription, error) {
		return filterer.WatchRegisteredIdentity(&bind.WatchOpts{
			Context: ctx,
		}, sink, nil, accountantIDs)
	})
	go func() {
		subErr := <-sub.Err()
		if subErr != nil {
			log.Error().Err(err).Msg("subscription error")
		}
		close(sink)
	}()
	return sink, sub.Unsubscribe, nil
}

// SubscribeToConsumerChannelBalanceUpdate subscribes to consumer channel balance update events
func (bc *Blockchain) SubscribeToConsumerChannelBalanceUpdate(mystSCAddress common.Address, channelAddresses []common.Address) (sink chan *bindings.MystTokenTransfer, cancel func(), err error) {
	filterer, err := bindings.NewMystTokenFilterer(mystSCAddress, bc.client)
	if err != nil {
		return sink, cancel, errors.Wrap(err, "could not create myst token filterer")
	}

	sink = make(chan *bindings.MystTokenTransfer)
	sub := event.Resubscribe(DefaultBackoff, func(ctx context.Context) (event.Subscription, error) {
		return filterer.WatchTransfer(&bind.WatchOpts{
			Context: ctx,
		}, sink, nil, channelAddresses)
	})
	go func() {
		subErr := <-sub.Err()
		if subErr != nil {
			log.Error().Err(err).Msg("subscription error")
		}
		close(sink)
	}()
	return sink, sub.Unsubscribe, nil
}

// SubscribeToProviderChannelBalanceUpdate subscribes to provider channel balance update events
func (bc *Blockchain) SubscribeToProviderChannelBalanceUpdate(accountantAddress common.Address, channelAddresses [][32]byte) (sink chan *bindings.AccountantImplementationChannelBalanceUpdated, cancel func(), err error) {
	filterer, err := bindings.NewAccountantImplementationFilterer(accountantAddress, bc.client)
	if err != nil {
		return sink, cancel, errors.Wrap(err, "could not create accountant implementation filterer")
	}

	sink = make(chan *bindings.AccountantImplementationChannelBalanceUpdated)
	sub := event.Resubscribe(DefaultBackoff, func(ctx context.Context) (event.Subscription, error) {
		return filterer.WatchChannelBalanceUpdated(&bind.WatchOpts{
			Context: ctx,
		}, sink, channelAddresses)
	})
	go func() {
		subErr := <-sub.Err()
		if subErr != nil {
			log.Error().Err(err).Msg("subscription error")
		}
		close(sink)
	}()
	return sink, sub.Unsubscribe, nil
}

// SettleRequest represents all the parameters required for settle
type SettleRequest struct {
	WriteRequest
	ChannelID common.Address
	Promise   crypto.Promise
}

// SettlePromise is settling the given consumer issued promise
func (bc *Blockchain) SettlePromise(req SettleRequest) (*types.Transaction, error) {
	transactor, err := bindings.NewChannelImplementationTransactor(req.ChannelID, bc.client)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()

	amount := big.NewInt(0).SetUint64(req.Promise.Amount)
	fee := big.NewInt(0).SetUint64(req.Promise.Fee)
	lock := [32]byte{}
	copy(lock[:], req.Promise.R)

	nonce, err := bc.getNonce(req.Identity)
	if err != nil {
		return nil, errors.Wrap(err, "could not get nonce")
	}

	return transactor.SettlePromise(&bind.TransactOpts{
		From:     req.Identity,
		Signer:   req.Signer,
		Context:  ctx,
		GasLimit: req.GasLimit,
		GasPrice: req.GasPrice,
		Nonce:    big.NewInt(0).SetUint64(nonce),
	},
		amount, fee, lock, req.Promise.Signature,
	)
}

func (bc *Blockchain) getNonce(identity common.Address) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()
	return bc.client.PendingNonceAt(ctx, identity)
}

// SubscribeToChannelOpenedEvents subscribes to provider channel opened events
func (bc *Blockchain) SubscribeToChannelOpenedEvents(accountantAddress common.Address) (sink chan *bindings.AccountantImplementationChannelOpened, cancel func(), err error) {
	filterer, err := bindings.NewAccountantImplementationFilterer(accountantAddress, bc.client)
	if err != nil {
		return sink, cancel, errors.Wrap(err, "could not create accountant implementation filterer")
	}

	sink = make(chan *bindings.AccountantImplementationChannelOpened)
	sub := event.Resubscribe(DefaultBackoff, func(ctx context.Context) (event.Subscription, error) {
		return filterer.WatchChannelOpened(&bind.WatchOpts{
			Context: ctx,
		}, sink)
	})
	go func() {
		subErr := <-sub.Err()
		if subErr != nil {
			log.Error().Err(err).Msg("subscription error")
		}
		close(sink)
	}()
	return sink, sub.Unsubscribe, nil
}

// SubscribeToPromiseSettledEventByChannelID subscribes to promise settled events
func (bc *Blockchain) SubscribeToPromiseSettledEventByChannelID(accountantID common.Address, providerAddresses [][32]byte) (sink chan *bindings.AccountantImplementationPromiseSettled, cancel func(), err error) {
	caller, err := bindings.NewAccountantImplementationFilterer(accountantID, bc.client)
	if err != nil {
		return sink, cancel, errors.Wrap(err, "could not create accountant caller")
	}
	sink = make(chan *bindings.AccountantImplementationPromiseSettled)

	sub := event.Resubscribe(DefaultBackoff, func(ctx context.Context) (event.Subscription, error) {
		return caller.WatchPromiseSettled(&bind.WatchOpts{
			Context: ctx,
		}, sink, providerAddresses)
	})
	go func() {
		subErr := <-sub.Err()
		if subErr != nil {
			log.Error().Err(err).Msg("subscription error")
		}
		cancel()
		close(sink)
	}()

	return sink, sub.Unsubscribe, nil
}

// NetworkID returns the network id
func (bc *Blockchain) NetworkID() (*big.Int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), bc.bcTimeout)
	defer cancel()
	return bc.client.NetworkID(ctx)
}

// ConsumerChannel represents the consumer channel
type ConsumerChannel struct {
	Settled *big.Int
	Balance *big.Int
}

// GetConsumerChannel returns the consumer channel
func (bc *Blockchain) GetConsumerChannel(addr common.Address, mystSCAddress common.Address) (ConsumerChannel, error) {
	ad := common.BytesToAddress(addr.Bytes())
	party, err := bc.GetConsumerChannelsAccountant(ad)
	if err != nil {
		return ConsumerChannel{}, err
	}

	balance, err := bc.GetMystBalance(mystSCAddress, ad)
	return ConsumerChannel{
		Settled: party.Settled,
		Balance: balance,
	}, err
}
