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

package transfer

import (
	"errors"
	"fmt"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

func TestGasPriceIncrementor(t *testing.T) {
	chid := int64(9223372036854775790)
	cfg := GasIncrementorConfig{
		WorkerCount:       2,
		PullEntryCount:    100,
		PullInterval:      time.Millisecond * 3,
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
		err := inc.InsertInitial(org, opts, sender)
		assert.NoError(t, err)

		for i := 0; i < 10; i++ {
			<-time.After(cfg.PullInterval)
			txs, _ := st.GetIncrementorTransactionsToCheck(10, []string{sender.Hex()})
			assert.Len(t, txs, 1)
			tx := txs[0]

			work := make(chan Transaction, 1)
			work <- tx
			close(work)
			inc.watchOrIncrease(work)

			if TxStateSucceed == tx.State {
				break
			}
		}

		assert.True(t, st.inserted)
		assert.True(t, st.pulled)
		assert.True(t, c.sent, "should be sent")
		assert.True(t, c.checked, "should be checked")
		assert.True(t, sg.signed, "should be signed")
		altered, err := st.tx.getLatestTx()
		assert.NoError(t, err)

		priceIncreased := 0
		for _, s := range st.stateHistory {
			if s == TxStatePriceIncreased {
				priceIncreased++
			}
		}
		assert.Equal(t, 1, priceIncreased)
		assert.Equal(t, TxStateSucceed, st.stateHistory[len(st.stateHistory)-1])

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
		assert.NoError(t, inc.InsertInitial(org, opts, sender))

		for i := 0; i < 10; i++ {
			<-time.After(cfg.PullInterval)
			txs, _ := st.GetIncrementorTransactionsToCheck(10, []string{sender.Hex()})
			assert.Len(t, txs, 1)
			tx := txs[0]

			work := make(chan Transaction, 1)
			work <- tx
			close(work)
			inc.watchOrIncrease(work)

			if TxStateSucceed == tx.State {
				break
			}
		}

		assert.True(t, st.inserted)
		assert.True(t, st.pulled)
		assert.False(t, c.sent, "already mind, should not check")
		assert.False(t, sg.signed, "tx should not be signed")
		assert.True(t, c.checked, "should be checked")
		priceIncreased := 0
		for _, s := range st.stateHistory {
			if s == TxStatePriceIncreased {
				priceIncreased++
			}
		}
		assert.Equal(t, 0, priceIncreased)
		assert.Equal(t, TxStateSucceed, st.stateHistory[len(st.stateHistory)-1])

		altered, err := st.tx.getLatestTx()
		assert.NoError(t, err)
		assert.Equal(t, chid, st.tx.ChainID)
		assert.Equal(t, originalGasPrice, altered.GasPrice(), "original gas price should stay the same")
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

func Test_caclulateGasPrice(t *testing.T) {
	for _, test := range []struct {
		tx           Transaction
		currentPrice *big.Int

		expect *big.Int
		noErr  bool
	}{
		{
			tx: Transaction{
				Opts: TransactionOpts{
					PriceMultiplier: 2,
					MaxPrice:        big.NewInt(1000),
				},
			},
			currentPrice: big.NewInt(100),
			expect:       big.NewInt(200),
			noErr:        true,
		},
		{
			tx: Transaction{
				Opts: TransactionOpts{
					PriceMultiplier: 2,
					MaxPrice:        big.NewInt(1000),
				},
			},
			currentPrice: big.NewInt(900),
			expect:       big.NewInt(1000),
			noErr:        true,
		},
		{
			tx: Transaction{
				Opts: TransactionOpts{
					PriceMultiplier: 2,
					MaxPrice:        big.NewInt(1000),
				},
			},
			currentPrice: big.NewInt(1000),
			expect:       nil,
			noErr:        false,
		},
		{
			tx: Transaction{
				Opts: TransactionOpts{
					PriceMultiplier: 2,
					MaxPrice:        big.NewInt(1000),
				},
			},
			currentPrice: big.NewInt(1200),
			expect:       nil,
			noErr:        false,
		},
	} {
		got, err := calculateGasPrice(test.tx, test.currentPrice)
		assert.Equal(t, test.expect, got)
		assert.Equal(t, test.noErr, err == nil)
	}
}

func defaultOpts() TransactionOpts {
	return TransactionOpts{
		PriceMultiplier:  2.0,
		MaxPrice:         new(big.Int).SetInt64(100),
		IncreaseInterval: time.Millisecond * 2,
		CheckInterval:    time.Millisecond * 1,
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

func (s *mockStorage) GetIncrementorTransactionsToCheck(count int64, signers []string) (tx []Transaction, err error) {
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
