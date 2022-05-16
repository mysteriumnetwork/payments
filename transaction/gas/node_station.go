package gas

import (
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
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

	blockNumber, err := n.bc.BlockNumber(n.chainID)
	if err != nil {
		return nil, err
	}

	header, err := n.bc.HeaderByNumber(n.chainID, big.NewInt(int64(blockNumber)))
	if err != nil {
		return nil, err
	}
	config := params.MainnetChainConfig
	baseFee := misc.CalcBaseFee(config, header)

	return &GasPrices{
		SafeLow: suggestGasPrice,
		Average: suggestGasPrice,
		Fast:    suggestGasPrice,
		BaseFee: baseFee,
	}, nil
}
