package crypto

import (
	"math/big"

	"github.com/shopspring/decimal"
)

// Myst represents a single myst ERC 777 token.
//
// Deprecated: this constant is deprecated and will be removed.
const Myst uint64 = 1000_000_000_000_000_000

// MystSize defines MYST size - precision for floating points.
//
// Deprecated: this constant is deprecated and will be removed.
const MystSize uint = 18

const mystSizeInt32 int32 = int32(MystSize)

var bigMyst = big.NewFloat(0).SetUint64(Myst)

// BigMystToFloat takes in a big int and returns a float64 representation of the myst.
//
// Deprecated: all unit conversion functions are now in the units package. Use BigIntWeiToFloatEth instead.
func BigMystToFloat(input *big.Int) float64 {
	f := new(big.Float).SetInt(input)
	divided := f.Quo(f, bigMyst)
	r, _ := divided.Float64()
	return r
}

// DecimalToBigMyst is an accurate conversion between a decimal and big int.
// If number has greater precision than `MystSize` it will get truncated.
//
// Deprecated: all unit conversion functions are now in the units package. Use DecimalEthToBigIntWei instead.
func DecimalToBigMyst(input decimal.Decimal) *big.Int {
	cleaned := input.Truncate(mystSizeInt32)
	if cleaned.IsZero() {
		return new(big.Int)
	}

	shifted := input.Shift(mystSizeInt32)
	return shifted.BigInt()
}

// BigMystToDecimal converts a BigMyst to a decimal.
//
// Deprecated: all unit conversion functions are now in the units package. Use BigIntWeiToDecimalEth instead.
func BigMystToDecimal(input *big.Int) decimal.Decimal {
	return decimal.NewFromBigInt(input, -mystSizeInt32)
}

// FloatToBigMyst takes in a float converts it to a big int representation.
// For example, 1.5 becomes 1500_000_000_000_000_000.
//
// Deprecated: all unit conversion functions are now in the units package. Use FloatEthToBigIntWei instead.
func FloatToBigMyst(input float64) *big.Int {
	multiplied := new(big.Float).Mul(new(big.Float).SetFloat64(input), bigMyst)
	res, _ := multiplied.Int(nil)
	return res
}
