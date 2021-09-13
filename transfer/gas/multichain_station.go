package gas

import (
	"fmt"
)

// MultichainStation is a station that can hold multiple station
// and call them depending on the chain given.
type MultichainStation map[int64]Station

func (m MultichainStation) GetGasPrices(chainID int64) (*GasPrices, error) {
	station, ok := m[chainID]
	if !ok {
		return nil, fmt.Errorf("no gas station for chain %d", chainID)
	}

	return station.GetGasPrices()
}
