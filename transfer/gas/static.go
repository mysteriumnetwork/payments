package gas

import "math/big"

// StaticStation returns static non changing gas price.
type StaticStation struct {
	staticPrice *big.Int
}

func NewStaticStation(price *big.Int) *StaticStation {
	return &StaticStation{staticPrice: price}
}

func (s *StaticStation) GetAverageGasPrice() (*big.Int, error) {
	return new(big.Int).Set(s.staticPrice), nil
}

func (s *StaticStation) GetFastGasPrice() (*big.Int, error) {
	return new(big.Int).Set(s.staticPrice), nil
}

func (s *StaticStation) GetLowGasPrice() (*big.Int, error) {
	return new(big.Int).Set(s.staticPrice), nil
}
