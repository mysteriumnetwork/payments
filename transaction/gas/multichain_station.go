package gas

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

// MultichainStation is a station that can hold multiple station
// and call them depending on the chain given.
type MultichainStation map[int64][]Station

func (m MultichainStation) GetGasPrices(chainID int64) (*GasPrices, error) {
	stations, ok := m[chainID]
	if !ok {
		return nil, fmt.Errorf("no gas stations for chain %d", chainID)
	}

	for i, station := range stations {
		prices, err := station.GetGasPrices()
		if err != nil {
			log.Error().Err(err).Int64("chainID", chainID).Int("stationIndex", i).Msg("failed to get gas prices")
			continue
		}
		return prices, nil
	}

	return nil, fmt.Errorf("all gas station for chain %d failed", chainID)
}
