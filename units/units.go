package units

import (
	"math/big"

	"github.com/shopspring/decimal"
)

var (
	// singleStep is the amount of units between steps
	// in eth currency representation
	singleStep = big.NewInt(1_000_000_000)
	// oneEth is a single eth represented in eth.
	oneEth = big.NewInt(1)
	// oneEthInGwei is a single eth represented in gwei.
	oneEthInGwei = new(big.Int).Mul(oneEth, singleStep)
	// oneEthInWei is a single eth represented in wei.
	oneEthInWei = new(big.Int).Mul(oneEthInGwei, singleStep)
	// ethDecimals is the amount of decimal places of eth.
	ethDecimals int32 = 18
)

// FloatGweiToBigIntWei returns the given gwei as wei.
func FloatGweiToBigIntWei(gwei float64) *big.Int {
	wei, _ := new(big.Float).Mul(
		big.NewFloat(gwei),
		new(big.Float).SetInt(oneEthInGwei),
	).Int(nil)
	return wei
}

// FloatEthToBigIntWei returns the given eth as wei. Precision might be lost.
func FloatEthToBigIntWei(eth float64) *big.Int {
	wei, _ := new(big.Float).Mul(
		big.NewFloat(eth),
		new(big.Float).SetInt(oneEthInWei),
	).Int(nil)
	return wei
}

// BigIntWeiToFloatEth takes in a big int wei value and returns a float64 ether representation.
func BigIntWeiToFloatEth(input *big.Int) float64 {
	f := new(big.Float).SetInt(input)
	ethInGwei := new(big.Float).SetInt(oneEthInWei)
	res, _ := f.Quo(f, ethInGwei).Float64()
	return res
}

// DecimalEthToBigIntWei is an accurate conversion between a decimal eth value and big int wei value.
// If number has greater precision than `ethDecimals` it will get truncated.
func DecimalEthToBigIntWei(input decimal.Decimal) *big.Int {
	cleaned := input.Truncate(ethDecimals)
	if cleaned.IsZero() {
		return big.NewInt(0)
	}

	shifted := input.Shift(ethDecimals)
	return shifted.BigInt()
}

// BigIntWeiToDecimal converts a bg int wei value to a decimal eth.
func BigIntWeiToDecimalEth(input *big.Int) decimal.Decimal {
	return decimal.NewFromBigInt(input, -ethDecimals)
}
