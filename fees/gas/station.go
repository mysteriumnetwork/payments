package gas

import "math/big"

// Station is a gas station inteface that provides methods
// to get gas prices in a network.
type Station interface {
	// GetAverageGasPrice returns the current average gas price.
	GetAverageGasPrice() (*big.Int, error)
	// GetLowGasPrice returns the lowest possible suggest gas price.
	GetLowGasPrice() (*big.Int, error)
	// GetFastGasPrice returns gas price suggested for fast transactions.
	GetFastGasPrice() (*big.Int, error)
}

func priceMaxUpperBound(price *big.Int, bound *big.Int) *big.Int {
	if price.Cmp(bound) > 0 {
		return bound
	}
	return price
}
