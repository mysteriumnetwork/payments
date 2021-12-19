package courier

import (
	"fmt"
	"math/big"

	"github.com/mysteriumnetwork/payments/transaction"

	"github.com/ethereum/go-ethereum/common"
)

type transferData struct {
	Amount *big.Int       `json:"amount"`
	To     common.Address `json:"to"`
}

type mystTransferData struct {
	transferData
	MystAddress common.Address `json:"myst_address"`
}

func (d *transferData) validate() error {
	if d.Amount == nil || d.Amount.Cmp(big.NewInt(0)) == 0 {
		return fmt.Errorf("amount cannot be 0: %w", transaction.ErrImpossibleToDeliver)
	}
	if d.To.Hex() == common.HexToAddress("").Hex() {
		return fmt.Errorf("recipient address cannot be empty: %w", transaction.ErrImpossibleToDeliver)
	}

	return nil
}

func (d *mystTransferData) validate() error {
	if err := d.transferData.validate(); err != nil {
		return err
	}
	if d.MystAddress.Hex() == common.HexToAddress("").Hex() {
		return fmt.Errorf("myst address cannot be empty: %w", transaction.ErrImpossibleToDeliver)
	}

	return nil
}
