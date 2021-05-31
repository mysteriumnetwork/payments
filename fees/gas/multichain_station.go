package gas

import (
	"fmt"
	"math/big"
)

// MultichainStation is a station that can hold multiple station
// and call them depending on the chain given.
type MultichainStation map[int64]Station

func (m MultichainStation) GetAveragePrice(chainID int64) (*big.Int, error) {
	station, ok := m[chainID]
	if !ok {
		return nil, fmt.Errorf("no gas station for chain %d", chainID)
	}

	return station.GetAverageGasPrice()
}

func (m MultichainStation) GetFastPrice(chainID int64) (*big.Int, error) {
	station, ok := m[chainID]
	if !ok {
		return nil, fmt.Errorf("no gas station for chain %d", chainID)
	}

	return station.GetFastGasPrice()
}

func (m MultichainStation) GetLowPrice(chainID int64) (*big.Int, error) {
	station, ok := m[chainID]
	if !ok {
		return nil, fmt.Errorf("no gas station for chain %d", chainID)
	}

	return station.GetLowGasPrice()
}
