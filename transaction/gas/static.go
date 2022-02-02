package gas

import (
	"math/big"
)

// StaticStation returns static non changing gas price.
type StaticStation struct {
	tip  *big.Int
	base *big.Int
}

func NewStaticStation(tip, base *big.Int) *StaticStation {
	return &StaticStation{
		tip:  tip,
		base: base,
	}
}

func (s *StaticStation) GetGasPrices() (*GasPrices, error) {
	prices := GasPrices{
		SafeLow: new(big.Int).Set(s.tip),
		Average: new(big.Int).Set(s.tip),
		Fast:    new(big.Int).Set(s.tip),
		BaseFee: s.base,
	}
	return &prices, nil
}
