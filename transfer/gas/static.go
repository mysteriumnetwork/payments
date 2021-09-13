package gas

import "math/big"

// StaticStation returns static non changing gas price.
type StaticStation struct {
	staticPrice *big.Int
}

func NewStaticStation(price *big.Int) *StaticStation {
	return &StaticStation{staticPrice: price}
}

func (s *StaticStation) GetGasPrices() (*GasPrices, error) {
	prices := GasPrices{
		SafeLow: new(big.Int).Set(s.staticPrice),
		Average: new(big.Int).Set(s.staticPrice),
		Fast:    new(big.Int).Set(s.staticPrice),
	}
	return &prices, nil
}
