package units

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnits(t *testing.T) {
	t.Run("eth to wei", func(t *testing.T) {
		t.Run("from float", func(t *testing.T) {
			one, ok := new(big.Int).SetString("1000000000000000000", 10)
			require.True(t, ok)
			type args struct {
				input float64
			}
			tests := []struct {
				name string
				args args
				want *big.Int
			}{
				{
					name: "calculates a single eth correctly",
					args: args{
						input: 1.0,
					},
					want: one,
				},
				{
					name: "calculates half eth",
					args: args{
						input: 0.5,
					},
					want: big.NewInt(0).Div(one, big.NewInt(2)),
				},
				{
					name: "calculates a small amount",
					args: args{
						input: 0.000000005,
					},
					want: big.NewInt(0).SetUint64(5000_000_000),
				},
			}
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := FloatEthToBigIntWei(tt.args.input); got.Cmp(tt.want) != 0 {
						t.Errorf("FloatEthToBigIntWei() = %v, want %v", got, tt.want)
					}
				})
			}
		})

		t.Run("from decimal", func(t *testing.T) {
			for _, test := range []struct {
				input decimal.Decimal
				get   string
			}{
				{
					input: decimal.RequireFromString("12312.312312312311231231"),
					get:   "12312312312312311231231",
				},
				{
					input: decimal.RequireFromString("12312.31231231231123123123"),
					get:   "12312312312312311231231",
				},
				{
					input: decimal.RequireFromString("12312.3123"),
					get:   "12312312300000000000000",
				},
				{
					input: decimal.RequireFromString("12312"),
					get:   "12312000000000000000000",
				},
				{
					input: decimal.RequireFromString("0.0000000000001"),
					get:   "100000",
				},
				{
					input: decimal.RequireFromString("0.000000000000000001"),
					get:   "1",
				},
				{
					input: decimal.RequireFromString("0"),
					get:   "0",
				},
				{
					input: decimal.RequireFromString("0.0000000000000000000000000000001"),
					get:   "0",
				},
			} {
				result := DecimalEthToBigIntWei(test.input)
				assert.Equal(t, test.get, result.String())
			}
		})
	})

	t.Run("gwei to wei", func(t *testing.T) {
		t.Run("float to big int", func(t *testing.T) {
			oneGweiInWei, ok := new(big.Int).SetString("1000000000", 10)
			require.True(t, ok)
			type args struct {
				input float64
			}
			tests := []struct {
				name string
				args args
				want *big.Int
			}{
				{
					name: "calculates a single gwei correctly",
					args: args{
						input: 1.0,
					},
					want: oneGweiInWei,
				},
				{
					name: "calculates half gwei",
					args: args{
						input: 0.5,
					},
					want: big.NewInt(0).Div(oneGweiInWei, big.NewInt(2)),
				},
				{
					name: "calculates a small amount",
					args: args{
						input: 0.000000001,
					},
					want: big.NewInt(1),
				},
				{
					name: "calculates a big amount",
					args: args{
						input: 1000000000000,
					},
					want: big.NewInt(0).Mul(oneGweiInWei, big.NewInt(1000000000000)),
				},
			}
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := FloatGweiToBigIntWei(tt.args.input); got.Cmp(tt.want) != 0 {
						t.Errorf("FloatGweiToBigIntWei() = %v, want %v", got, tt.want)
					}
				})
			}
		})
	})

	t.Run("wei to eth", func(t *testing.T) {
		t.Run("to float", func(t *testing.T) {
			oneEthInWei, ok := new(big.Int).SetString("1000000000000000000", 10)
			require.True(t, ok)
			type args struct {
				input *big.Int
			}
			tests := []struct {
				name string
				args args
				want float64
			}{
				{
					name: "calculates a single eth correctly",
					args: args{
						input: oneEthInWei,
					},
					want: 1.0,
				},
				{
					name: "calculates half eth",
					args: args{
						input: big.NewInt(0).Div(oneEthInWei, big.NewInt(2)),
					},
					want: 0.5,
				},
				{
					name: "calculates a small amount",
					args: args{
						input: big.NewInt(0).SetUint64(5000_000_000),
					},
					want: 0.000000005,
				},
			}
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := BigIntWeiToFloatEth(tt.args.input); got != tt.want {
						t.Errorf("BigIntWeiToFloatEth() = %v, want %v", got, tt.want)
					}
				})
			}
		})

		t.Run("to decimal", func(t *testing.T) {
			mustBig := func(s string) *big.Int {
				res, ok := big.NewInt(0).SetString(s, 10)
				if !ok {
					panic("failed big int")
				}
				return res
			}
			for i, test := range []struct {
				input *big.Int
				get   decimal.Decimal
			}{
				{
					input: mustBig("12312312312312311231231"),
					get:   decimal.RequireFromString("12312.312312312311231231"),
				},
				{
					input: mustBig("1231231231231231123123112"),
					get:   decimal.RequireFromString("1231231.231231231123123112"),
				},
				{
					input: mustBig("12312312300000000000000"),
					get:   decimal.RequireFromString("12312.3123"),
				},
				{
					input: mustBig("12312000000000000000000"),
					get:   decimal.RequireFromString("12312"),
				},
				{
					input: mustBig("100000"),
					get:   decimal.RequireFromString("0.0000000000001"),
				},
				{
					input: mustBig("1"),
					get:   decimal.RequireFromString("0.000000000000000001"),
				},
				{
					input: mustBig("0"),
					get:   decimal.RequireFromString("0"),
				},
			} {
				result := BigIntWeiToDecimalEth(test.input)
				assert.Equal(t, test.get.String(), result.String(), fmt.Sprintf("test %d failed", i+1))
			}
		})
	})
}
