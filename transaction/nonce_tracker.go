package transaction

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

// DepotNonceTracker keeps track of nonces atomically.
type DepotNonceTracker struct {
	nonceTrackerBC nonceTrackerBC
	ds             DepotStorage

	nonces    map[Sender]uint64
	nonceLock sync.Mutex
}

type nonceTrackerBC interface {
	PendingNonceAt(chainID int64, account common.Address) (uint64, error)
	NonceAt(chainID int64, account common.Address, blockNum *big.Int) (uint64, error)
}

// NewDepotNonceTracker returns a new nonce tracker.
func NewDepotNonceTracker(nonceTrackerBC nonceTrackerBC, ds DepotStorage) *DepotNonceTracker {
	return &DepotNonceTracker{
		nonceTrackerBC: nonceTrackerBC,
		nonces:         make(map[Sender]uint64),
		ds:             ds,
	}
}

// GetNextNonce returns an atomically increasing nonce for the account.
func (nt *DepotNonceTracker) GetNextNonce(chainID int64, account common.Address) (uint64, error) {
	nt.nonceLock.Lock()
	defer nt.nonceLock.Unlock()

	key := NewSender(account, chainID)
	if v, ok := nt.nonces[key]; ok {
		v++
		nt.nonces[key] = v
		return v, nil
	}

	lastKnown, err := nt.ds.GetLastQueuedDelivery(chainID, account)
	if err != nil {
		return 0, err
	}

	persistentNonce := uint64(0)
	if lastKnown != nil {
		persistentNonce = lastKnown.Nonce + 1
	}

	bcNonce, err := nt.nonceTrackerBC.PendingNonceAt(chainID, account)
	if err != nil {
		return bcNonce, err
	}

	nonce := persistentNonce
	if bcNonce > persistentNonce {
		// If BC nonce is larger, that means we've missed some transaction and didnt account for them
		// in the persistent storage. Issue a nonce that is up to date.
		nonce = bcNonce
	}

	nt.nonces[key] = nonce
	return nonce, nil
}

func (nt *DepotNonceTracker) GetConfirmedNonce(chainID int64, account common.Address) (uint64, error) {
	bcNonce, err := nt.nonceTrackerBC.NonceAt(chainID, account, nil)
	if err != nil {
		return bcNonce, err
	}

	return bcNonce, nil
}

// ForceReloadNonce clears the nonce cache. This will force loading from BC next time.
func (nt *DepotNonceTracker) ForceReloadNonce(chainID int64, account common.Address) {
	nt.nonceLock.Lock()
	defer nt.nonceLock.Unlock()
	delete(nt.nonces, NewSender(account, chainID))
}
