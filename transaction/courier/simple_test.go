package courier

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mysteriumnetwork/payments/v3/client"
	"github.com/mysteriumnetwork/payments/v3/transaction"
	"github.com/stretchr/testify/assert"
)

func TestSimpleCourier(t *testing.T) {
	pk, err := crypto.GenerateKey()
	assert.NoError(t, err)
	publicKeyECDSA, ok := pk.Public().(*ecdsa.PublicKey)
	assert.True(t, ok)
	senderAddr := crypto.PubkeyToAddress(*publicKeyECDSA)

	mockBCClient := mockBCClient{}
	signFactory := func(sender common.Address, chain int64) transaction.SignFunc {
		return func(signAddr common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if sender != senderAddr {
				return nil, fmt.Errorf("unknown sender")
			}
			if signAddr != senderAddr {
				return nil, fmt.Errorf("unknown signer address")
			}
			return types.SignTx(tx, types.NewLondonSigner(big.NewInt(chain)), pk)
		}
	}

	courier := NewSimpleCourier(&mockBCClient, signFactory)

	resetFunc := func() {
		mockBCClient.reset()
	}

	t.Run("myst transfer", func(t *testing.T) {
		defer resetFunc()

		to := common.HexToAddress("0x742d13F0b2A19C823bdd362b16305e4704b97A31")
		mystAddress := common.HexToAddress("0x742d13F0b2A19C823bdd362b16305e4704b97A32")
		var chainId int64 = 1
		amount := big.NewInt(100)

		var deliveryRequest transaction.DeliveryRequest
		t.Run("creates transfer", func(t *testing.T) {
			deliveryRequest, err = courier.NewMystTransferDelivery(transaction.NewSender(senderAddr, chainId), amount, to, mystAddress)
			assert.NoError(t, err)

			assert.Equal(t, chainId, deliveryRequest.ChainID)
			assert.Equal(t, senderAddr, deliveryRequest.Sender)
			actualData, err := json.Marshal(deliveryRequest.Data)
			assert.NoError(t, err)
			expectedData, err := json.Marshal(
				mystTransferData{
					transferData: transferData{
						Amount: amount,
						To:     to,
					},
					MystAddress: mystAddress,
				})
			assert.NoError(t, err)
			assert.Equal(t, expectedData, actualData)
			assert.Equal(t, deliveryTypeMystTransfer, deliveryRequest.Type)
		})

		t.Run("can deliver", func(t *testing.T) {
			canDeliver := courier.CanDeliver(deliveryRequest.Type)
			assert.True(t, canDeliver)
		})

		t.Run("delivers it", func(t *testing.T) {
			var nonce uint64 = 0
			blob, err := json.Marshal(deliveryRequest.Data)
			assert.NoError(t, err)
			now := time.Now()

			gasTip := big.NewInt(10)
			baseFee := big.NewInt(100)

			tx, err := courier.DeliverTransaction(transaction.Delivery{
				UniqueID: fmt.Sprintf("%s|%d|%d", deliveryRequest.Sender.Hex(), nonce, deliveryRequest.ChainID),
				Sender:   deliveryRequest.Sender,

				Nonce:    nonce,
				ChainID:  deliveryRequest.ChainID,
				GasPrice: nil,
				GasTip:   gasTip,
				BaseFee:  baseFee,

				Type:  deliveryRequest.Type,
				State: transaction.DeliveryStateWaiting,

				ShipmentData:    blob,
				SentTransaction: []byte{},

				CreatedUTC: now,
				UpdateUTC:  now,
			})
			assert.NoError(t, err)

			assert.Equal(t, big.NewInt(chainId), tx.ChainId())
			assert.Equal(t, new(big.Int).Add(baseFee, gasTip), tx.GasPrice())
			assert.Equal(t, []byte(fmt.Sprintf("send %s MYST to %s", amount.String(), to.Hex())), tx.Data())
			assert.Equal(t, nonce, tx.Nonce())
			assert.Equal(t, mystAddress, *tx.To())

			assert.Equal(t, 1, mockBCClient.sentMystTransfers)
			assert.Equal(t, 0, mockBCClient.sentTransfers)
		})
	})

	t.Run("native transfer", func(t *testing.T) {
		defer resetFunc()

		to := common.HexToAddress("0x742d13F0b2A19C823bdd362b16305e4704b97A31")
		var chainId int64 = 137
		amount := big.NewInt(100)

		var deliveryRequest transaction.DeliveryRequest
		t.Run("creates transfer", func(t *testing.T) {
			deliveryRequest, err = courier.NewNetworkTransferDelivery(transaction.NewSender(senderAddr, chainId), amount, to)
			assert.NoError(t, err)

			assert.Equal(t, chainId, deliveryRequest.ChainID)
			assert.Equal(t, senderAddr, deliveryRequest.Sender)
			actualData, err := json.Marshal(deliveryRequest.Data)
			assert.NoError(t, err)
			expectedData, err := json.Marshal(
				transferData{
					Amount: amount,
					To:     to,
				})
			assert.NoError(t, err)
			assert.Equal(t, expectedData, actualData)
			assert.Equal(t, deliveryTypeNetworkTransfer, deliveryRequest.Type)
		})

		t.Run("can deliver", func(t *testing.T) {
			canDeliver := courier.CanDeliver(deliveryRequest.Type)
			assert.True(t, canDeliver)
		})

		t.Run("delivers it", func(t *testing.T) {
			var nonce uint64 = 1
			blob, err := json.Marshal(deliveryRequest.Data)
			assert.NoError(t, err)
			now := time.Now()

			gasTip := big.NewInt(10)
			baseFee := big.NewInt(100)

			tx, err := courier.DeliverTransaction(transaction.Delivery{
				UniqueID: fmt.Sprintf("%s|%d|%d", deliveryRequest.Sender.Hex(), nonce, deliveryRequest.ChainID),
				Sender:   deliveryRequest.Sender,

				Nonce:    nonce,
				ChainID:  deliveryRequest.ChainID,
				GasPrice: nil,
				GasTip:   gasTip,
				BaseFee:  baseFee,

				Type:  deliveryRequest.Type,
				State: transaction.DeliveryStateWaiting,

				ShipmentData:    blob,
				SentTransaction: []byte{},

				CreatedUTC: now,
				UpdateUTC:  now,
			})
			assert.NoError(t, err)

			assert.Equal(t, big.NewInt(chainId), tx.ChainId())
			assert.Equal(t, new(big.Int).Add(baseFee, gasTip), tx.GasPrice())
			assert.Equal(t, []byte(fmt.Sprintf("send %s to %s", amount.String(), to.Hex())), tx.Data())
			assert.Equal(t, nonce, tx.Nonce())
			assert.Equal(t, to, *tx.To())

			assert.Equal(t, 0, mockBCClient.sentMystTransfers)
			assert.Equal(t, 1, mockBCClient.sentTransfers)
		})
	})

	t.Run("errors", func(t *testing.T) {
		defer resetFunc()

		t.Run("unknown type", func(t *testing.T) {
			defer resetFunc()

			var chainId int64 = 137
			var deliveryRequest transaction.DeliveryRequest
			t.Run("creates transfer", func(t *testing.T) {
				deliveryRequest = transaction.DeliveryRequest{
					ChainID: chainId,
					Sender:  senderAddr,
					Type:    transaction.DeliverableType("unknown"),
					Data:    struct{}{},
				}

				assert.Equal(t, chainId, deliveryRequest.ChainID)
				assert.Equal(t, senderAddr, deliveryRequest.Sender)
				actualData, err := json.Marshal(deliveryRequest.Data)
				assert.NoError(t, err)
				assert.NoError(t, err)
				assert.Equal(t, []byte("{}"), actualData)
				assert.Equal(t, "unknown", string(deliveryRequest.Type))
			})

			t.Run("cannot deliver", func(t *testing.T) {
				canDeliver := courier.CanDeliver(deliveryRequest.Type)
				assert.False(t, canDeliver)
			})

			t.Run("deliver returns error", func(t *testing.T) {
				var nonce uint64 = 1
				blob, err := json.Marshal(deliveryRequest.Data)
				assert.NoError(t, err)
				now := time.Now()

				gasTip := big.NewInt(10)
				baseFee := big.NewInt(100)

				_, err = courier.DeliverTransaction(transaction.Delivery{
					UniqueID: fmt.Sprintf("%s|%d|%d", deliveryRequest.Sender.Hex(), nonce, deliveryRequest.ChainID),
					Sender:   deliveryRequest.Sender,

					Nonce:    nonce,
					ChainID:  deliveryRequest.ChainID,
					GasPrice: nil,
					GasTip:   gasTip,
					BaseFee:  baseFee,

					Type:  deliveryRequest.Type,
					State: transaction.DeliveryStateWaiting,

					ShipmentData:    blob,
					SentTransaction: []byte{},

					CreatedUTC: now,
					UpdateUTC:  now,
				})
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "type \"unknown\" is impossible to handle")
			})
		})

		t.Run("blockchain error", func(t *testing.T) {
			defer resetFunc()

			to := common.HexToAddress("0x742d13F0b2A19C823bdd362b16305e4704b97A31")
			var chainId int64 = 137
			amount := big.NewInt(100)

			var deliveryRequest transaction.DeliveryRequest
			t.Run("creates transfer", func(t *testing.T) {
				deliveryRequest, err = courier.NewNetworkTransferDelivery(transaction.NewSender(senderAddr, chainId), amount, to)
				assert.NoError(t, err)

				assert.Equal(t, chainId, deliveryRequest.ChainID)
				assert.Equal(t, senderAddr, deliveryRequest.Sender)
				actualData, err := json.Marshal(deliveryRequest.Data)
				assert.NoError(t, err)
				expectedData, err := json.Marshal(
					transferData{
						Amount: amount,
						To:     to,
					})
				assert.NoError(t, err)
				assert.Equal(t, expectedData, actualData)
				assert.Equal(t, deliveryTypeNetworkTransfer, deliveryRequest.Type)
			})

			t.Run("can deliver", func(t *testing.T) {
				canDeliver := courier.CanDeliver(deliveryRequest.Type)
				assert.True(t, canDeliver)
			})

			t.Run("fails to deliver because of blockchain error", func(t *testing.T) {
				blockchainError := fmt.Errorf("blockchain error")
				mockBCClient.errToReturn = blockchainError

				var nonce uint64 = 2
				blob, err := json.Marshal(deliveryRequest.Data)
				assert.NoError(t, err)
				now := time.Now()

				_, err = courier.DeliverTransaction(transaction.Delivery{
					UniqueID: fmt.Sprintf("%s|%d|%d", deliveryRequest.Sender.Hex(), nonce, deliveryRequest.ChainID),
					Sender:   deliveryRequest.Sender,

					Nonce:    nonce,
					ChainID:  deliveryRequest.ChainID,
					GasPrice: nil,
					GasTip:   big.NewInt(10),
					BaseFee:  big.NewInt(100),

					Type:  deliveryRequest.Type,
					State: transaction.DeliveryStateWaiting,

					ShipmentData:    blob,
					SentTransaction: []byte{},

					CreatedUTC: now,
					UpdateUTC:  now,
				})
				assert.Error(t, err)
				assert.ErrorIs(t, err, blockchainError)
			})
		})

		t.Run("missing data", func(t *testing.T) {
			defer resetFunc()

			var chainId int64 = 137

			t.Run("no amount", func(t *testing.T) {
				var nonce uint64 = 2
				to := common.HexToAddress("0x742d13F0b2A19C823bdd362b16305e4704b97A31")
				blob, err := json.Marshal(transferData{
					Amount: nil,
					To:     to,
				})
				assert.NoError(t, err)
				now := time.Now()

				_, err = courier.DeliverTransaction(transaction.Delivery{
					UniqueID: fmt.Sprintf("%s|%d|%d", senderAddr, nonce, chainId),
					Sender:   senderAddr,

					Nonce:    nonce,
					ChainID:  chainId,
					GasPrice: nil,
					GasTip:   big.NewInt(10),
					BaseFee:  big.NewInt(100),

					Type:  deliveryTypeNetworkTransfer,
					State: transaction.DeliveryStateWaiting,

					ShipmentData:    blob,
					SentTransaction: []byte{},

					CreatedUTC: now,
					UpdateUTC:  now,
				})
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "amount cannot be 0")
				assert.ErrorIs(t, err, transaction.ErrImpossibleToDeliver)
			})

			t.Run("no recipient", func(t *testing.T) {
				var nonce uint64 = 3
				blob, err := json.Marshal(transferData{
					Amount: big.NewInt(100),
					To:     common.Address{},
				})
				assert.NoError(t, err)
				now := time.Now()

				_, err = courier.DeliverTransaction(transaction.Delivery{
					UniqueID: fmt.Sprintf("%s|%d|%d", senderAddr, nonce, chainId),
					Sender:   senderAddr,

					Nonce:    nonce,
					ChainID:  chainId,
					GasPrice: nil,
					GasTip:   big.NewInt(10),
					BaseFee:  big.NewInt(100),

					Type:  deliveryTypeNetworkTransfer,
					State: transaction.DeliveryStateWaiting,

					ShipmentData:    blob,
					SentTransaction: []byte{},

					CreatedUTC: now,
					UpdateUTC:  now,
				})
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "recipient address cannot be empty")
				assert.ErrorIs(t, err, transaction.ErrImpossibleToDeliver)
			})

			t.Run("no myst address", func(t *testing.T) {
				var nonce uint64 = 4
				blob, err := json.Marshal(mystTransferData{
					transferData{
						Amount: big.NewInt(100),
						To:     common.HexToAddress("0x742d13F0b2A19C823bdd362b16305e4704b97A31"),
					},
					common.Address{},
				})
				assert.NoError(t, err)
				now := time.Now()

				_, err = courier.DeliverTransaction(transaction.Delivery{
					UniqueID: fmt.Sprintf("%s|%d|%d", senderAddr, nonce, chainId),
					Sender:   senderAddr,

					Nonce:    nonce,
					ChainID:  chainId,
					GasPrice: nil,
					GasTip:   big.NewInt(10),
					BaseFee:  big.NewInt(100),

					Type:  deliveryTypeMystTransfer,
					State: transaction.DeliveryStateWaiting,

					ShipmentData:    blob,
					SentTransaction: []byte{},

					CreatedUTC: now,
					UpdateUTC:  now,
				})
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "myst address cannot be empty")
				assert.ErrorIs(t, err, transaction.ErrImpossibleToDeliver)
			})
		})
	})
}

type mockBCClient struct {
	sentTransfers     int
	sentMystTransfers int
	lock              sync.Mutex
	errToReturn       error
}

func (m *mockBCClient) TransferMyst(chainID int64, req client.TransferRequest) (*types.Transaction, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.errToReturn != nil {
		return nil, m.errToReturn
	}

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(chainID),
		To:        &req.MystAddress,
		Nonce:     req.Nonce.Uint64(),
		GasTipCap: req.GasTip,
		GasFeeCap: new(big.Int).Add(req.GasTip, req.BaseFee),
		Value:     big.NewInt(0),
		Gas:       req.GasLimit,
		Data:      []byte(fmt.Sprintf("send %s MYST to %s", req.Amount.String(), req.Recipient.Hex())),
	})
	signedTx, err := req.Signer(req.Identity, tx)
	if err != nil {
		return nil, fmt.Errorf("could not sign tx: %w", err)
	}

	m.sentMystTransfers++
	return signedTx, nil
}

func (m *mockBCClient) TransferEth(chainID int64, etr client.EthTransferRequest) (*types.Transaction, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.errToReturn != nil {
		return nil, m.errToReturn
	}

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(chainID),
		To:        &etr.To,
		Nonce:     etr.Nonce.Uint64(),
		GasTipCap: etr.GasTip,
		GasFeeCap: new(big.Int).Add(etr.GasTip, etr.BaseFee),
		Value:     etr.Amount,
		Gas:       etr.GasLimit,
		Data:      []byte(fmt.Sprintf("send %s to %s", etr.Amount.String(), etr.To.Hex())),
	})
	signedTx, err := etr.Signer(etr.Identity, tx)
	if err != nil {
		return nil, fmt.Errorf("could not sign tx: %w", err)
	}

	m.sentTransfers++
	return signedTx, nil
}

func (m *mockBCClient) reset() {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.sentTransfers = 0
	m.sentMystTransfers = 0
	m.errToReturn = nil
}
