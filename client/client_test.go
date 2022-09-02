package client

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mysteriumnetwork/payments/client/mocks"
	"github.com/mysteriumnetwork/payments/crypto"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	t.Run("client creation", func(t *testing.T) {
		pendingNonceErr := fmt.Errorf("pendingNonce error")
		cl := &mocks.EtherClientMock{PendingNonceAtFunc: func(ctx context.Context, account common.Address) (uint64, error) {
			return 0, pendingNonceErr
		}}
		getter := NewDefaultEthClientGetter(cl)

		bc := NewBlockchain(getter, time.Second)
		assert.Same(t, cl, bc.ethClient.Client())
		assert.Equal(t, time.Second, bc.bcTimeout)
		_, err := bc.nonceFunc(context.Background(), common.Address{})
		assert.Equal(t, pendingNonceErr, err)

		customErr := fmt.Errorf("custom error")
		nonceFunc := func(ctx context.Context, account common.Address) (uint64, error) {
			return 0, customErr
		}
		bcWithCustomNonceTracker := NewBlockchainWithCustomNonceTracker(getter, 2*time.Second, nonceFunc)
		assert.Same(t, cl, bcWithCustomNonceTracker.ethClient.Client())
		assert.Equal(t, 2*time.Second, bcWithCustomNonceTracker.bcTimeout)
		_, err = bcWithCustomNonceTracker.nonceFunc(context.Background(), common.Address{})
		assert.Equal(t, customErr, err)
	})

	t.Run("calls correct function", func(t *testing.T) {
		calls := make(map[string]int)
		callsLock := &sync.Mutex{}
		cl := getMockClientWithFunctionCallCounter(calls, callsLock)
		getter := NewDefaultEthClientGetter(cl)
		bc := NewBlockchain(getter, time.Second)

		_, _, err := bc.TransactionByHash(common.Hash{})
		assert.NoError(t, err)
		_, err = bc.TransactionReceipt(common.Hash{})
		assert.NoError(t, err)
		_, err = bc.PendingNonceAt(common.Address{})
		assert.NoError(t, err)
		_, err = bc.NonceAt(common.Address{}, nil)
		assert.NoError(t, err)
		_, err = bc.GetEthBalance(common.Address{})
		assert.NoError(t, err)
		_, err = bc.FilterLogs(ethereum.FilterQuery{})
		assert.NoError(t, err)
		_, err = bc.HeaderByNumber(nil)
		assert.NoError(t, err)
		_, err = bc.SuggestGasPrice()
		assert.NoError(t, err)
		_, err = bc.BlockNumber()
		assert.NoError(t, err)
		_, err = bc.NetworkID()
		assert.NoError(t, err)
		err = bc.SendTransaction(nil)
		assert.NoError(t, err)

		assert.Len(t, calls, 11)
		for _, v := range calls {
			assert.Equal(t, 1, v)
		}
	})

	t.Run("get channel id", func(t *testing.T) {
		bc := NewBlockchain(NewDefaultEthClientGetter(&mocks.EtherClientMock{}), time.Second)
		hermesId := common.HexToAddress("0x80Ed28d84792d8b153bf2F25F0C4B7a1381dE4ab")
		address := common.HexToAddress("0xEb2623a2734E272BCC07Bda954863F316f4Bd4CF")
		channel, err := bc.getProviderChannelAddressBytes(hermesId, address)
		assert.NoError(t, err)
		assert.Equal(t, common.Hex2Bytes("bff791cc2eff82736bd18305ca116dbcd569957426036d6a970b7c88dacb6af1"), channel[:])

		channel, err = bc.getProviderChannelAddressForWithdrawalBytes(hermesId, address)
		assert.NoError(t, err)
		assert.Equal(t, common.Hex2Bytes("624e678e381a4e8b1c8a782997dbada9c86303ce880a3e0f20ba5422b33ca4d2"), channel[:])
	})

	t.Run("registration request", func(t *testing.T) {
		identity := common.HexToAddress("0xEb2623a2734E272BCC07Bda954863F316f4Bd4CF")
		hermes := common.HexToAddress("0xEb2623a2734E272BCC07Bda954863F316f4Bd4C1")
		beneficiary := common.HexToAddress("0xEb2623a2734E272BCC07Bda954863F316f4Bd4C2")
		stake := big.NewInt(4)
		tf := big.NewInt(2)
		signature := common.Hex2Bytes("abcd")
		rr := RegistrationRequest{
			WriteRequest: WriteRequest{
				Identity: identity,
			},
			HermesID:      hermes,
			Stake:         stake,
			TransactorFee: tf,
			Beneficiary:   beneficiary,
			Signature:     signature,
		}
		eo := rr.toEstimateOps()
		assert.Equal(t, identity, eo.From)
		assert.EqualValues(t, []interface{}{hermes, stake, tf, beneficiary, signature}, eo.Params)
		assert.Equal(t, "registerIdentity", eo.Method)
	})

	t.Run("write request", func(t *testing.T) {
		t.Run("gas limit", func(t *testing.T) {
			gl := uint64(3)
			wr := WriteRequest{
				GasLimit: gl,
			}
			assert.Equal(t, gl, wr.getGasLimit())
		})

		t.Run("transact ops creation EIP-1559", func(t *testing.T) {
			pendingNonceErr := fmt.Errorf("pendingNonce error")
			address := common.HexToAddress("0x123")
			cl := &mocks.EtherClientMock{
				PendingNonceAtFunc: func(ctx context.Context, account common.Address) (uint64, error) {
					if account == address {
						return 2, nil
					}
					return 0, pendingNonceErr
				},
			}
			getter := NewDefaultEthClientGetter(cl)
			bc := NewBlockchain(getter, time.Second)
			wr := WriteRequest{
				GasTip:   big.NewInt(2),
				BaseFee:  big.NewInt(3),
				Identity: address,
			}
			to, err := bc.makeTransactOpts(context.Background(), &wr)
			assert.NoError(t, err)
			assert.Equal(t, wr.GasTip, to.GasTipCap)
			assert.Equal(t, big.NewInt(5), to.GasFeeCap)
			assert.Equal(t, big.NewInt(2), to.Nonce)
			assert.Equal(t, address, to.From)
		})

		t.Run("transact ops creation pre EIP-1559", func(t *testing.T) {
			pendingNonceErr := fmt.Errorf("pendingNonce error")
			address := common.HexToAddress("0x123")
			cl := &mocks.EtherClientMock{
				PendingNonceAtFunc: func(ctx context.Context, account common.Address) (uint64, error) {
					if account == address {
						return 2, nil
					}
					return 0, pendingNonceErr
				},
			}
			getter := NewDefaultEthClientGetter(cl)
			bc := NewBlockchain(getter, time.Second)
			wr := WriteRequest{
				GasPrice: big.NewInt(2),
				Identity: address,
			}
			to, err := bc.makeTransactOpts(context.Background(), &wr)
			assert.NoError(t, err)
			assert.Equal(t, wr.GasPrice, to.GasPrice)
			assert.Equal(t, big.NewInt(2), to.Nonce)
			assert.Equal(t, address, to.From)
		})
	})

	t.Run("open consumer channel request", func(t *testing.T) {
		identity := common.HexToAddress("0xEb2623a2734E272BCC07Bda954863F316f4Bd4CF")
		hermes := common.HexToAddress("0xEb2623a2734E272BCC07Bda954863F316f4Bd4C1")
		registry := common.HexToAddress("0xEb2623a2734E272BCC07Bda954863F316f4Bd4C2")
		tf := big.NewInt(2)
		signature := common.Hex2Bytes("abcd")
		rr := OpenConsumerChannelRequest{
			WriteRequest: WriteRequest{
				Identity: identity,
			},
			HermesID:        hermes,
			RegistryAddress: registry,
			TransactorFee:   tf,
			Signature:       signature,
		}
		eo := rr.toEstimateOps()
		assert.Equal(t, identity, eo.From)
		assert.EqualValues(t, []interface{}{hermes, registry, tf, signature}, eo.Params)
		assert.Equal(t, "openConsumerChannel", eo.Method)
	})

	t.Run("pay and settle request", func(t *testing.T) {
		identity := common.HexToAddress("0xEb2623a2734E272BCC07Bda954863F316f4Bd4CF")
		provider := common.HexToAddress("0xEb2623a2734E272BCC07Bda954863F316f4Bd4C2")
		beneficiary := common.HexToAddress("0xEb2623a2734E272BCC07Bda954863F316f4Bd4C3")
		promiseAmount := big.NewInt(4)
		promiseFee := big.NewInt(3)
		r := common.Hex2Bytes("123456")
		signature := common.Hex2Bytes("abcd")
		beneficiarySignature := common.Hex2Bytes("123abc")
		pasr := PayAndSettleRequest{
			WriteRequest: WriteRequest{
				Identity: identity,
			},
			ProviderID: provider,
			Promise: crypto.Promise{
				Amount:    promiseAmount,
				Fee:       promiseFee,
				R:         r,
				Signature: signature,
			},
			Beneficiary:          beneficiary,
			BeneficiarySignature: beneficiarySignature,
		}
		eo := pasr.toEstimateOps()
		assert.Equal(t, identity, eo.From)
		assert.EqualValues(t, []interface{}{provider, pasr.Promise.Amount, pasr.Promise.Fee, ToBytes32(pasr.Promise.R), pasr.Promise.Signature, pasr.Beneficiary, pasr.BeneficiarySignature}, eo.Params)
		assert.Equal(t, "payAndSettle", eo.Method)
	})
}
