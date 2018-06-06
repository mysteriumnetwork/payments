package promises

import (
	"github.com/MysteriumNetwork/payments/registry"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type Promise struct {
	ServiceConsumer common.Address
	Receiver        common.Address
	SeqNo           int64
	Amount          int64
}

const issuerPrefix = "Issuer prefix:"

func (p *Promise) ConsumerHash() []byte {
	return crypto.Keccak256(p.ServiceConsumer.Bytes())
}

func (p *Promise) HashBytes() []byte {
	return crypto.Keccak256([]byte(issuerPrefix), p.ConsumerHash(), p.Receiver.Bytes(), abi.U256(big.NewInt(p.SeqNo)), abi.U256(big.NewInt(p.Amount)))
}

type IssuedPromise struct {
	Promise
	IssuerSignature []byte
}

type ReceivedPromise struct {
	IssuedPromise
	ReceiverSignature []byte
}

func SignByPayer(promise *Promise, payer *registry.MystIdentity) (*IssuedPromise, error) {
	hash := promise.HashBytes()
	signature, err := crypto.Sign(hash, payer.PrivateKey)
	if err != nil {
		return nil, err
	}

	return &IssuedPromise{
		*promise,
		signature,
	}, nil
}

const receiverPrefix = "Receiver prefix:"

func SignByReceiver(promise *IssuedPromise, receiver *registry.MystIdentity) (*ReceivedPromise, error) {
	publicKey, err := crypto.Ecrecover(promise.HashBytes(), promise.IssuerSignature)
	if err != nil {
		return nil, err
	}
	pubKey := crypto.ToECDSAPub(publicKey)
	payerAddr := crypto.PubkeyToAddress(*pubKey)
	hash := crypto.Keccak256([]byte(receiverPrefix), promise.HashBytes(), payerAddr.Bytes())
	sig, err := crypto.Sign(hash, receiver.PrivateKey)
	return &ReceivedPromise{
		*promise,
		sig,
	}, nil
}
