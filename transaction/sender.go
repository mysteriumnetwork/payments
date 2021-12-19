package transaction

import "github.com/ethereum/go-ethereum/common"

type Sender struct {
	Address common.Address
	ChainID int64
}

func NewSender(addr common.Address, chain int64) Sender {
	return Sender{
		Address: addr,
		ChainID: chain,
	}
}
