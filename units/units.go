package units

import "math/big"

var (
	// singleStep is the amount of units between steps
	// in eth currency representation
	singleStep = new(big.Int).SetInt64(1000000000)
	// eth is a single eth represented in eth.
	eth = new(big.Int).SetInt64(1)
	// gwei is a single eth represented in gwei.
	gwei = new(big.Int).Mul(eth, singleStep)
	// wei is a single eth represented in wei.
	wei = new(big.Int).Mul(gwei, singleStep)
)

// SingleEthInGwei returns amount of gwei in one eth.
func SingleEthInGwei() *big.Int {
	return new(big.Int).Set(gwei)
}

// SingleEthInWei returns amount of wei in one eth.
func SingleEthInWei() *big.Int {
	return new(big.Int).Set(wei)
}

// SingleGweiInWei returns amount of wei in one gwei.
func SingleGweiInWei() *big.Int {
	return new(big.Int).Set(gwei)
}

// FloatGweiToBigIntWei returns the given gwei as wei.
func FloatGweiToBigIntWei(gwei float64) *big.Int {
	wei, _ := new(big.Float).Mul(
		big.NewFloat(gwei),
		new(big.Float).SetInt(SingleGweiInWei()),
	).Int(nil)

	return wei
}
