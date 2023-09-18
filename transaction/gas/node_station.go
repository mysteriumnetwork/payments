package gas

import (
	"math/big"

	"github.com/ethereum/go-ethereum/consensus/misc/eip1559"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

type NodeStation struct {
	bc      BCClient
	chainID int64
}

func NewNodeStation(bc BCClient, chainID int64) *NodeStation {
	return &NodeStation{bc: bc, chainID: chainID}
}

type BCClient interface {
	SuggestGasPrice(chainID int64) (*big.Int, error)
	HeaderByNumber(chainID int64, number *big.Int) (*types.Header, error)
	BlockNumber(chainID int64) (uint64, error)
}

func (n *NodeStation) GetGasPrices() (*GasPrices, error) {
	suggestGasPrice, err := n.bc.SuggestGasPrice(n.chainID)
	if err != nil {
		return nil, err
	}

	header, err := n.bc.HeaderByNumber(n.chainID, nil)
	if err != nil {
		return nil, err
	}
	config := params.MainnetChainConfig
	baseFee := eip1559.CalcBaseFee(config, header)

	return &GasPrices{
		SafeLow: suggestGasPrice,
		Average: suggestGasPrice,
		Fast:    suggestGasPrice,
		BaseFee: baseFee,
	}, nil
}
