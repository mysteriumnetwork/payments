package transaction

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

// NonceTracker keeps track of nonces atomically.
type NonceTracker struct {
	nonceTrackerBC nonceTrackerBC
	ds             DepotStorage

	nonces    map[Sender]uint64
	nonceLock sync.Mutex
}

type nonceTrackerBC interface {
	PendingNonceAt(chainID int64, account common.Address) (uint64, error)
	NonceAt(chainID int64, account common.Address, blockNum *big.Int) (uint64, error)
}

// NewNonceTracker returns a new nonce tracker.
func NewNonceTracker(nonceTrackerBC nonceTrackerBC, ds DepotStorage) *NonceTracker {
	return &NonceTracker{
		nonceTrackerBC: nonceTrackerBC,
		nonces:         make(map[Sender]uint64),
		ds:             ds,
	}
}

type nonceSetFn func(nonce uint64) error

// GetNextNonce returns an atomically increasing nonce for the account.
func (nt *NonceTracker) SetNextNonce(chainID int64, account common.Address, fn nonceSetFn) error {
	nt.nonceLock.Lock()
	defer nt.nonceLock.Unlock()

	key := NewSender(account, chainID)
	if v, ok := nt.nonces[key]; ok {
		v++
		return nt.setWithExec(key, v, fn)
	}

	lastKnown, err := nt.ds.GetLastQueuedDelivery(chainID, account)
	if err != nil {
		return err
	}

	persistentNonce := uint64(0)
	if lastKnown != nil {
		persistentNonce = lastKnown.Nonce + 1
	}

	bcNonce, err := nt.nonceTrackerBC.PendingNonceAt(chainID, account)
	if err != nil {
		return err
	}

	nonce := persistentNonce
	if bcNonce > persistentNonce {
		// If BC nonce is larger, that means we've missed some transaction and didnt account for them
		// in the persistent storage. Issue a nonce that is up to date.
		nonce = bcNonce
	}

	return nt.setWithExec(key, nonce, fn)
}

func (nt *NonceTracker) setWithExec(key Sender, nonce uint64, fn nonceSetFn) error {
	if err := fn(nonce); err != nil {
		return err
	}

	nt.nonces[key] = nonce
	return nil
}

func (nt *NonceTracker) GetConfirmedNonce(chainID int64, account common.Address) (uint64, error) {
	bcNonce, err := nt.nonceTrackerBC.NonceAt(chainID, account, nil)
	if err != nil {
		return bcNonce, err
	}

	return bcNonce, nil
}

// ForceReloadNonce clears the nonce cache. This will force loading from BC next time.
func (nt *NonceTracker) ForceReloadNonce(chainID int64, account common.Address) {
	nt.nonceLock.Lock()
	defer nt.nonceLock.Unlock()
	delete(nt.nonces, NewSender(account, chainID))
}
