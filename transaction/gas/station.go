package gas

import "math/big"

// Station is a gas station inteface that provides methods
// to get gas prices in a network.
type Station interface {
	// GetGasPrice returns the gas prices (Low, Average and Fast).
	GetGasPrices() (*GasPrices, error)
}

type GasPrices struct {
	SafeLow *big.Int
	Average *big.Int
	Fast    *big.Int

	BaseFee *big.Int
}

func priceMaxUpperBound(price *big.Int, bound *big.Int) *big.Int {
	if price.Cmp(bound) > 0 {
		return bound
	}
	return price
}
