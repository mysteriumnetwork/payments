package courier

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/mysteriumnetwork/payments/client"
	"github.com/mysteriumnetwork/payments/transaction"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Simple is the default simple courier that implements
// `transaction.DeliveryCourier` interface. It can handle
// only the most basic transactions.
type Simple struct {
	bc SimpleBCClient
	sf SignerFactory
}

type SimpleBCClient interface {
	TransferMyst(chainID int64, req client.TransferRequest) (*types.Transaction, error)
	TransferEth(chainID int64, etr client.EthTransferRequest) (*types.Transaction, error)
}

// SignerFactory given a sender and chain should produce a signature func
// which can be used to sign transactions.
type SignerFactory func(sender common.Address, chain int64) transaction.SignFunc

const (
	deliveryTypeNetworkTransfer transaction.DeliverableType = "network-transfer"
	deliveryTypeMystTransfer    transaction.DeliverableType = "myst-transfer"
)

func NewSimpleCourier(bc SimpleBCClient, sf SignerFactory) *Simple {
	return &Simple{
		bc: bc,
		sf: sf,
	}
}

func (s *Simple) DeliverTransaction(td transaction.Delivery) (*types.Transaction, error) {
	fn := s.getDeliveryFunc(td.Type)
	if fn == nil {
		return nil, fmt.Errorf("type %q is impossible to handle", td.Type)
	}

	return fn(td)
}

func (s *Simple) CanDeliver(typ transaction.DeliverableType) bool {
	return s.getDeliveryFunc(typ) != nil
}

func (c *Simple) NewNetworkTransferDelivery(sender transaction.Sender, amount *big.Int, to common.Address) (transaction.DeliveryRequest, error) {
	nt := transferData{
		Amount: amount,
		To:     to,
	}

	return transaction.DeliveryRequest{
		ChainID: sender.ChainID,
		Sender:  sender.Address,
		Type:    deliveryTypeNetworkTransfer,
		Data:    nt,
	}, nt.validate()
}

func (c *Simple) NewMystTransferDelivery(sender transaction.Sender, amount *big.Int, to common.Address, mystAddr common.Address) (transaction.DeliveryRequest, error) {
	mt := mystTransferData{
		transferData: transferData{
			Amount: amount,
			To:     to,
		},
		MystAddress: mystAddr,
	}

	return transaction.DeliveryRequest{
		ChainID: sender.ChainID,
		Sender:  sender.Address,
		Type:    deliveryTypeMystTransfer,
		Data:    mt,
	}, mt.validate()
}

func (s *Simple) getDeliveryFunc(typ transaction.DeliverableType) func(transaction.Delivery) (*types.Transaction, error) {
	switch typ {
	case deliveryTypeNetworkTransfer:
		return s.networkTransfer
	case deliveryTypeMystTransfer:
		return s.mystTransfer
	default:
		return nil
	}
}

func (c *Simple) networkTransfer(td transaction.Delivery) (*types.Transaction, error) {
	wr := td.ToWriteRequest(c.sf(td.Sender, td.ChainID), 50000)

	var extra transferData
	if err := json.Unmarshal(td.ShipmentData, &extra); err != nil {
		return nil, err
	}

	if err := extra.validate(); err != nil {
		return nil, err
	}

	request := client.EthTransferRequest{
		WriteRequest: wr,
		Amount:       extra.Amount,
		To:           extra.To,
	}

	return c.bc.TransferEth(td.ChainID, request)
}

func (c *Simple) mystTransfer(td transaction.Delivery) (*types.Transaction, error) {
	wr := td.ToWriteRequest(c.sf(td.Sender, td.ChainID), 100000)

	var extra mystTransferData
	if err := json.Unmarshal(td.ShipmentData, &extra); err != nil {
		return nil, err
	}

	if err := extra.validate(); err != nil {
		return nil, err
	}

	request := client.TransferRequest{
		WriteRequest: wr,
		Amount:       extra.Amount,
		Recipient:    extra.To,
		MystAddress:  extra.MystAddress,
	}

	return c.bc.TransferMyst(td.ChainID, request)
}
