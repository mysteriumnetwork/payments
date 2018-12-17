package test_utils

import (
	"context"
	"fmt"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
)

type TransactionalBackend interface {
	bind.ContractBackend
	Commit()                                                                            //taken from Simulated backend
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) //taken from Simulated backend
}

// Keep in mind that initial amount is actually in Gwei - you need really HUGE
// numbers here to make more than few simple transactions
func NewSimulatedBackend(genesisAddress common.Address, initialAmmount int64) TransactionalBackend {
	sb := backends.NewSimulatedBackend(core.GenesisAlloc{
		genesisAddress: {Balance: big.NewInt(initialAmmount)},
	}, 10000000)
	return sb
}

type TxLoggingBackend struct {
	TransactionalBackend
	abis *AbiList
}

func (sb *TxLoggingBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	abiName, method, err := sb.abis.LookupByTx(tx)
	if err != nil {
		fmt.Printf("Error: %v. Tx data was: %v\n", err, common.ToHex(tx.Data()))
	} else {
		fmt.Printf("Tx: Contract: %s, Method: %s (%s), Gas: %d\n", abiName, method.Name, common.ToHex(method.Id()), tx.Gas())
	}
	return sb.TransactionalBackend.SendTransaction(ctx, tx)
}

func (sb *TxLoggingBackend) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	var methodName = ""
	abiName, method, err := sb.abis.InputLookup(call.Data)
	if err != nil {
		fmt.Printf("Uknown method id. Input: %v \n", call.Data)
		methodName = "<undefined>"
		abiName = "<undefined>"
	} else {
		methodName = method.String()
	}

	fmt.Printf("%v state read:  %v. Payload: %v\n", abiName, methodName, common.ToHex(call.Data[4:]))

	return sb.TransactionalBackend.CallContract(ctx, call, blockNumber)
}

func LoggingBackend(sb TransactionalBackend, abis *AbiList) TransactionalBackend {

	return &TxLoggingBackend{
		sb,
		abis,
	}
}
