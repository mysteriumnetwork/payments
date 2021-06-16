/* Mysterium network payment library.
 *
 * Copyright (C) 2021 BlockDev AG
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package fees

import (
	"errors"
	"fmt"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

func TestGasPriceIncrementor(t *testing.T) {
	chid := int64(9223372036854775790)
	cfg := GasIncrementorConfig{
		PullInterval:      time.Millisecond,
		MaxQueuePerSigner: 100,
	}
	t.Run("gas price is increased and tx is successfull", func(t *testing.T) {
		originalGasPrice := big.NewInt(1)
		org := types.NewTransaction(1, common.HexToAddress("0x1"), big.NewInt(1), 1, originalGasPrice, []byte{})
		opts := defaultOpts()
		st := &mockStorage{}
		c := newClient(big.NewInt(2))

		sender := common.HexToAddress("")
		sg := signer{}
		inc := NewGasPriceIncremenetor(cfg, st, c, map[common.Address]SignatureFunc{
			sender: sg.SignatureFunc,
		})
		go inc.Run()
		inc.InsertInitial(org, opts, sender)
		assert.Eventually(t, func() bool {
			txs, _ := st.GetIncrementorTransactionsToCheck([]string{sender.Hex()})
			if len(txs) == 0 {
				return false
			}
			tx := txs[0]
			return TxStateSucceed == tx.State
		}, time.Second, time.Millisecond*10)

		inc.Stop()
		assert.True(t, st.inserted)
		assert.True(t, st.pulled)
		assert.True(t, c.sent, "should be sent")
		assert.True(t, c.checked, "should be checked")
		assert.True(t, sg.signed, "should be signed")
		altered, err := st.tx.getLatestTx()
		assert.NoError(t, err)

		assert.Equal(t,
			[]TransactionState{TxStateCreated, TxStatePriceIncreased, TxStateSucceed},

			st.stateHistory)

		assert.Equal(t, chid, st.tx.ChainID)
		assert.Equal(t, big.NewInt(2), altered.GasPrice(), "gas price should increase")
	})
	t.Run("tx already mined", func(t *testing.T) {
		originalGasPrice := big.NewInt(1)
		org := types.NewTransaction(1, common.HexToAddress("0x1"), big.NewInt(1), 1, originalGasPrice, []byte{})
		opts := defaultOpts()
		st := &mockStorage{}
		c := newClient(big.NewInt(0))

		sender := common.HexToAddress("")
		sg := signer{}
		inc := NewGasPriceIncremenetor(cfg, st, c, map[common.Address]SignatureFunc{
			sender: sg.SignatureFunc,
		})
		go inc.Run()
		inc.InsertInitial(org, opts, sender)
		assert.Eventually(t, func() bool {
			txs, _ := st.GetIncrementorTransactionsToCheck([]string{sender.Hex()})
			if len(txs) == 0 {
				return false
			}
			tx := txs[0]

			return TxStateSucceed == tx.State
		}, time.Second, time.Millisecond*10)

		inc.Stop()
		assert.True(t, st.inserted)
		assert.True(t, st.pulled)
		assert.False(t, c.sent, "already mind, should not check")
		assert.False(t, sg.signed, "tx should not be signed")
		assert.True(t, c.checked, "should be checked")
		assert.Equal(t,
			[]TransactionState{TxStateCreated, TxStateSucceed},

			st.stateHistory)
		altered, err := st.tx.getLatestTx()
		assert.NoError(t, err)
		assert.Equal(t, chid, st.tx.ChainID)
		assert.Equal(t, originalGasPrice, altered.GasPrice(), "original gas price should stay the same")
	})
	t.Run("tx fails if gas price is too big", func(t *testing.T) {
		originalGasPrice := big.NewInt(2)
		org := types.NewTransaction(1, common.HexToAddress("0x1"), big.NewInt(1), 1, originalGasPrice, []byte{})
		opts := defaultOpts()
		opts.MaxPrice = new(big.Int).SetInt64(5)

		st := &mockStorage{}
		c := newClient(new(big.Int).Add(opts.MaxPrice, big.NewInt(5)))

		sender := common.HexToAddress("")
		sg := signer{}
		inc := NewGasPriceIncremenetor(cfg, st, c, map[common.Address]SignatureFunc{
			sender: sg.SignatureFunc,
		})
		go inc.Run()
		inc.InsertInitial(org, opts, sender)
		assert.Eventually(t, func() bool {
			txs, _ := st.GetIncrementorTransactionsToCheck([]string{sender.Hex()})
			if len(txs) == 0 {
				return false
			}
			tx := txs[0]

			return TxStateFailed == tx.State
		}, time.Second, time.Millisecond*10)

		inc.Stop()
		assert.True(t, st.inserted)
		assert.True(t, st.pulled)
		assert.True(t, c.sent, "should be sent once")
		assert.True(t, sg.signed, "tx should be signed")
		assert.Equal(t,
			[]TransactionState{TxStateCreated, TxStatePriceIncreased, TxStateFailed},

			st.stateHistory)

		altered, err := st.tx.getLatestTx()
		assert.Equal(t, chid, st.tx.ChainID)
		assert.NoError(t, err)
		assert.Equal(t, new(big.Int).SetInt64(4), altered.GasPrice(), "gas price should be increased once")
	})
	t.Run("invalid opts, not inserted", func(t *testing.T) {
		org := types.NewTransaction(1, common.HexToAddress("0x1"), big.NewInt(1), 1, big.NewInt(1), []byte{})

		st := &mockStorage{}
		c := newClient(nil)

		sender := common.HexToAddress("")
		sg := signer{}
		inc := NewGasPriceIncremenetor(cfg, st, c, map[common.Address]SignatureFunc{
			sender: sg.SignatureFunc,
		})
		assert.Error(t, inc.InsertInitial(org, TransactionOpts{}, sender))
	})
}

func Test_syncer(t *testing.T) {
	s := newSyncer()

	tx := Transaction{UniqueID: "0x0"}
	s.txMarkBeingWatched(tx)
	assert.True(t, s.txBeingWatched(tx), "transaction should be watched")
	s.txRemoveWatched(tx)
	assert.False(t, s.txBeingWatched(tx), "transaction should no longer be watched")
}

func defaultOpts() TransactionOpts {
	return TransactionOpts{
		PriceMultiplier:  2.0,
		MaxPrice:         new(big.Int).SetInt64(100),
		Timeout:          time.Minute,
		IncreaseInterval: time.Millisecond * 60,
		CheckInterval:    time.Millisecond * 40,
	}
}

type mockStorage struct {
	tx           Transaction
	stateHistory []TransactionState
	inserted     bool
	pulled       bool

	m sync.Mutex
}

func (s *mockStorage) UpsertIncrementorTransaction(tx Transaction) error {
	s.m.Lock()
	defer s.m.Unlock()

	if len(s.stateHistory) == 0 {
		s.stateHistory = []TransactionState{}
	}

	s.stateHistory = append(s.stateHistory, tx.State)
	s.tx = tx
	s.inserted = true
	return nil
}

func (s *mockStorage) GetIncrementorTransactionsToCheck(signers []string) (tx []Transaction, err error) {
	s.m.Lock()
	defer s.m.Unlock()
	s.pulled = true
	give := []Transaction{s.tx}
	return give, nil
}

func (s *mockStorage) GetIncrementorSenderQueue(sender string) (length int, err error) {
	return 0, nil
}

type mockClient struct {
	gasTreshold *big.Int
	currentGas  *big.Int
	sent        bool
	checked     bool
}

func newClient(gasTh *big.Int) *mockClient {
	return &mockClient{
		gasTreshold: gasTh,
		currentGas:  big.NewInt(0),
	}
}

func (c *mockClient) TransactionReceipt(chainID int64, hash common.Hash) (*types.Receipt, error) {
	c.checked = true
	if c.currentGas.Cmp(c.gasTreshold) >= 0 {
		return &types.Receipt{
			Status: types.ReceiptStatusSuccessful,
		}, nil
	}

	return &types.Receipt{
		Status: types.ReceiptStatusFailed,
	}, nil
}

func (c *mockClient) TransactionByHash(chainID int64, hash common.Hash) (*types.Transaction, bool, error) {
	return nil, false, nil
}

func (c *mockClient) SendTransaction(chainID int64, tx *types.Transaction) error {
	c.currentGas = tx.GasPrice()
	c.sent = true
	return nil
}

type signer struct {
	signed bool
}

func (s *signer) SignatureFunc(tx *types.Transaction, chainID int64) (*types.Transaction, error) {
	s.signed = true
	return tx, nil
}

func TestGasPriceIncremenetor_isBlockchainErrorUnhandleable(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "detects ethereum not found error",
			args: args{
				err: fmt.Errorf("tortilla %w", ethereum.NotFound),
			},
			want: true,
		},
		{
			name: "detects core nonce too low error",
			args: args{
				err: fmt.Errorf("multiwrap %w", fmt.Errorf("tortilla %w", core.ErrNonceTooLow)),
			},
			want: true,
		},
		{
			name: "detects nonce too low by string",
			args: args{
				err: fmt.Errorf("tortilla %w", errors.New("nonce too low")),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &GasPriceIncremenetor{}
			if got := i.isBlockchainErrorUnhandleable(tt.args.err); got != tt.want {
				t.Errorf("GasPriceIncremenetor.isBlockchainErrorUnhandleable() = %v, want %v", got, tt.want)
			}
		})
	}
}
