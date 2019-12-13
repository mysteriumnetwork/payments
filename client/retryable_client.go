package client

import (
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/mysteriumnetwork/payments/bindings"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type blockchain interface {
	GetAccountantFee(accountantAddress common.Address) (uint16, error)
	IsRegisteredAsProvider(accountantAddress, registryAddress, addressToCheck common.Address) (bool, error)
	GetProviderChannel(accountantAddress common.Address, addressToCheck common.Address) (ProviderChannel, error)
	IsRegistered(registryAddress, addressToCheck common.Address) (bool, error)
	SubscribeToPromiseSettledEvent(providerID, accountantID common.Address) (sink chan *bindings.AccountantImplementationPromiseSettled, cancel func(), err error)
	GetMystBalance(mystSCAddress, address common.Address) (*big.Int, error)
	SubscribeToConsumerBalanceEvent(channel, mystSCAddress common.Address) (chan *bindings.MystTokenTransfer, func(), error)
	GetRegistrationFee(registryAddress common.Address) (*big.Int, error)
	RegisterIdentity(rr RegistrationRequest) (*types.Transaction, error)
	TransferMyst(req TransferRequest) (tx *types.Transaction, err error)
	IsAccountantRegistered(registryAddress, acccountantID common.Address) (bool, error)
	GetAccountantOperator(accountantID common.Address) (common.Address, error)
	SettleAndRebalance(req SettleAndRebalanceRequest) (*types.Transaction, error)
	GetChannel(acc common.Address, chID []byte) (ProviderChannel, error)
}

// BlockchainWithRetries takes in the plain blockchain implementation and exposes methods that will retry the underlying bc methods before giving up.
// This is required as the ethereum client will occasionally spit a TLS error if running for prolonged periods of time.
type BlockchainWithRetries struct {
	delay      time.Duration
	maxRetries int
	bc         blockchain
	stop       chan struct{}
	once       sync.Once
}

// ErrStopped represents an error when a call is interrupted
var ErrStopped = errors.New("call stopped")

// NewBlockchainWithRetries returns a new instance of blockchain with retries
func NewBlockchainWithRetries(bc blockchain, delay time.Duration, maxRetries int) *BlockchainWithRetries {
	return &BlockchainWithRetries{
		bc:         bc,
		delay:      delay,
		maxRetries: maxRetries,
	}
}

func (bwr *BlockchainWithRetries) callWithRetry(f func() error) error {
	for i := 0; i < bwr.maxRetries; i++ {
		err := f()
		if err == nil {
			return nil
		}
		if i+1 == bwr.maxRetries {
			return err
		}

		log.Warn().Err(err).Msgf("retry %v of %v", i+1, bwr.maxRetries)
		select {
		case <-bwr.stop:
			return ErrStopped
		case <-time.After(bwr.delay):
		}
	}
	return nil
}

// GetAccountantFee fetches the accountant fee from blockchain
func (bwr *BlockchainWithRetries) GetAccountantFee(accountantAddress common.Address) (uint16, error) {
	var res uint16
	err := bwr.callWithRetry(func() error {
		r, err := bwr.bc.GetAccountantFee(accountantAddress)
		if err != nil {
			return errors.Wrap(err, "could not get accountant fee")
		}
		res = r
		return nil
	})
	return res, err
}

// IsRegisteredAsProvider checks if the provider is registered with the accountant properly
func (bwr *BlockchainWithRetries) IsRegisteredAsProvider(accountantAddress, registryAddress, addressToCheck common.Address) (bool, error) {
	var res bool
	err := bwr.callWithRetry(func() error {
		r, err := bwr.bc.IsRegisteredAsProvider(accountantAddress, registryAddress, addressToCheck)
		if err != nil {
			return errors.Wrap(err, "could not check if registered as provider")
		}
		res = r
		return nil
	})
	return res, err
}

// GetProviderChannel returns the provider channel
func (bwr *BlockchainWithRetries) GetProviderChannel(accountantAddress, addressToCheck common.Address) (ProviderChannel, error) {
	var res ProviderChannel
	err := bwr.callWithRetry(func() error {
		r, err := bwr.bc.GetProviderChannel(accountantAddress, addressToCheck)
		if err != nil {
			return errors.Wrap(err, "could not get provider channel")
		}
		res = r
		return nil
	})

	return res, err
}

// SubscribeToPromiseSettledEvent subscribes to promise settled events
func (bwr *BlockchainWithRetries) SubscribeToPromiseSettledEvent(providerID, accountantID common.Address) (chan *bindings.AccountantImplementationPromiseSettled, func(), error) {
	var sink chan *bindings.AccountantImplementationPromiseSettled
	var cancel func()
	err := bwr.callWithRetry(func() error {
		s, c, err := bwr.bc.SubscribeToPromiseSettledEvent(providerID, accountantID)
		if err != nil {
			return errors.Wrap(err, "could not subscribe to settlement events")
		}
		sink = s
		cancel = c
		return nil
	})
	return sink, cancel, err
}

// SubscribeToConsumerBalanceEvent subscribes to the consumer balance change events
func (bwr *BlockchainWithRetries) SubscribeToConsumerBalanceEvent(channel, mystSCAddress common.Address) (chan *bindings.MystTokenTransfer, func(), error) {
	var sink chan *bindings.MystTokenTransfer
	var cancel func()
	err := bwr.callWithRetry(func() error {
		s, c, err := bwr.bc.SubscribeToConsumerBalanceEvent(channel, mystSCAddress)
		if err != nil {
			return errors.Wrap(err, "could not subscribe to settlement events")
		}
		sink = s
		cancel = c
		return nil
	})
	return sink, cancel, err
}

// IsRegistered checks wether the given identity is registered or not
func (bwr *BlockchainWithRetries) IsRegistered(registryAddress, addressToCheck common.Address) (bool, error) {
	var res bool
	err := bwr.callWithRetry(func() error {
		r, err := bwr.bc.IsRegistered(registryAddress, addressToCheck)
		if err != nil {
			return errors.Wrap(err, "check registration status")
		}
		res = r
		return nil
	})
	return res, err
}

// GetMystBalance returns the balance in myst
func (bwr *BlockchainWithRetries) GetMystBalance(mystSCAddress, channel common.Address) (*big.Int, error) {
	var res *big.Int
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.GetMystBalance(mystSCAddress, channel)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not get myst balance")
		}
		res = result
		return nil
	})
	return res, err
}

// GetRegistrationFee returns the registration fee
func (bwr *BlockchainWithRetries) GetRegistrationFee(registryAddress common.Address) (*big.Int, error) {
	var res *big.Int
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.GetRegistrationFee(registryAddress)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not get registration fee")
		}
		res = result
		return nil
	})
	return res, err
}

// RegisterIdentity registers the given identity on blockchain
func (bwr *BlockchainWithRetries) RegisterIdentity(rr RegistrationRequest) (*types.Transaction, error) {
	var res *types.Transaction
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.RegisterIdentity(rr)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not register identity")
		}
		res = result
		return nil
	})
	return res, err
}

// TransferMyst transfers myst to the provided address
func (bwr *BlockchainWithRetries) TransferMyst(req TransferRequest) (tx *types.Transaction, err error) {
	err = bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.TransferMyst(req)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not register identity")
		}
		tx = result
		return nil
	})
	return tx, err
}

// IsAccountantRegistered checks if given accountant is registered and returns true or false.
func (bwr *BlockchainWithRetries) IsAccountantRegistered(registryAddress, acccountantID common.Address) (bool, error) {
	var res bool
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.IsAccountantRegistered(registryAddress, acccountantID)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not register identity")
		}
		res = result
		return nil
	})
	return res, err
}

// GetAccountantOperator returns operator address of given accountant
func (bwr *BlockchainWithRetries) GetAccountantOperator(accountantID common.Address) (common.Address, error) {
	var res common.Address
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.GetAccountantOperator(accountantID)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not register identity")
		}
		res = result
		return nil
	})
	return res, err
}

// SettleAndRebalance is settling given accountant issued promise
func (bwr *BlockchainWithRetries) SettleAndRebalance(req SettleAndRebalanceRequest) (*types.Transaction, error) {
	var res *types.Transaction
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.SettleAndRebalance(req)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not register identity")
		}
		res = result
		return nil
	})
	return res, err
}

// GetChannel returns the given channel information
func (bwr *BlockchainWithRetries) GetChannel(acc common.Address, chID []byte) (ProviderChannel, error) {
	var res ProviderChannel
	err := bwr.callWithRetry(func() error {
		result, bcErr := bwr.bc.GetChannel(acc, chID)
		if bcErr != nil {
			return errors.Wrap(bcErr, "could not register identity")
		}
		res = result
		return nil
	})
	return res, err
}

// Stop stops the blockhain with retries aborting any waits for retries
func (bwr *BlockchainWithRetries) Stop() {
	bwr.once.Do(func() {
		close(bwr.stop)
	})
}
