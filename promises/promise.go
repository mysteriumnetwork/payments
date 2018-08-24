package promises

import (
	"math/big"

	"github.com/MysteriumNetwork/payments/identity"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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

func (p *Promise) Bytes() []byte {
	slices := [][]byte{
		[]byte(issuerPrefix),
		p.ConsumerHash(),
		p.Receiver.Bytes(),
		abi.U256(big.NewInt(p.SeqNo)),
		abi.U256(big.NewInt(p.Amount)),
	}
	var res []byte
	for _, slice := range slices {
		res = append(res, slice...)
	}
	return res
}

type IssuedPromise struct {
	Promise
	IssuerSignature []byte
}

type ReceivedPromise struct {
	IssuedPromise
	ReceiverSignature []byte
}

func SignByPayer(promise *Promise, payer identity.Signer) (*IssuedPromise, error) {
	signature, err := payer.Sign(promise.Bytes())
	if err != nil {
		return nil, err
	}

	return &IssuedPromise{
		*promise,
		signature,
	}, nil
}

const receiverPrefix = "Receiver prefix:"

func SignByReceiver(promise *IssuedPromise, receiver identity.Signer) (*ReceivedPromise, error) {
	publicKey, err := crypto.Ecrecover(crypto.Keccak256(promise.Bytes()), promise.IssuerSignature)
	if err != nil {
		return nil, err
	}
	pubKey, err := crypto.UnmarshalPubkey(publicKey)
	if err != nil {
		return nil, err
	}
	payerAddr := crypto.PubkeyToAddress(*pubKey)
	sig, err := receiver.Sign([]byte(receiverPrefix), crypto.Keccak256(promise.Bytes()), payerAddr.Bytes())
	return &ReceivedPromise{
		*promise,
		sig,
	}, nil
}
