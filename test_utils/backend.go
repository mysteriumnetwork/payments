package test_utils

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"math/big"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"fmt"
)

type TransactionalBackend interface {
	bind.ContractBackend
	Commit()  //taken from Simulated backend
}


// Keep in mind that initial amount is actually in Gwei - you need really HUGE
// numbers here to make more than few simple transactions
func NewSimulatedBackend(genesisAddress common.Address, initialAmmount int64) (TransactionalBackend) {
	sb := backends.NewSimulatedBackend(core.GenesisAlloc{
		genesisAddress: {Balance: big.NewInt(initialAmmount)},
	})
	return sb
}

type TxLoggingBackend struct {
	TransactionalBackend
	abis * AbiList
}


func (sb * TxLoggingBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	abiName, method, err := sb.abis.Lookup(tx.Data())
	if err != nil {
		fmt.Printf("Error: %v. Tx data was: %v\n" , err, common.ToHex(tx.Data()))
	} else {
		fmt.Printf("Tx: Contract: %s, Method: %s, Gas: %d\n", abiName, method.String(), tx.Gas())
	}
	return sb.TransactionalBackend.SendTransaction(ctx, tx )
}
func LoggingBackend( sb TransactionalBackend , abis * AbiList) ( TransactionalBackend ) {

	return &TxLoggingBackend{
		sb,
		abis,
	}
}