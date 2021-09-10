/* Mysterium network payment library.
 *
 * Copyright (C) 2020 BlockDev AG
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package crypto

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestBigMystToFloat(t *testing.T) {
	singleMyst := big.NewInt(0).SetUint64(Myst)
	type args struct {
		input *big.Int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "calculates a single myst correctly",
			args: args{
				input: singleMyst,
			},
			want: 1.0,
		},
		{
			name: "calculates half a myst",
			args: args{
				input: big.NewInt(0).Div(singleMyst, big.NewInt(2)),
			},
			want: 0.5,
		},
		{
			name: "calculates a small amount of a myst",
			args: args{
				input: big.NewInt(0).SetUint64(5000_000_000),
			},
			want: 0.000000005,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BigMystToFloat(tt.args.input); got != tt.want {
				t.Errorf("BigMystToFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloatToBigMyst(t *testing.T) {
	singleMyst, _ := new(big.Float).Set(bigMyst).Int(nil)
	type args struct {
		input float64
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "calculates a single myst correctly",
			args: args{
				input: 1.0,
			},
			want: singleMyst,
		},
		{
			name: "calculates half a myst",
			args: args{
				input: 0.5,
			},
			want: big.NewInt(0).Div(singleMyst, big.NewInt(2)),
		},
		{
			name: "calculates a small amount of a myst",
			args: args{
				input: 0.000000005,
			},
			want: big.NewInt(0).SetUint64(5000_000_000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatToBigMyst(tt.args.input); got.Cmp(tt.want) != 0 {
				t.Errorf("FloatToBigMyst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBigMystToDecimal(t *testing.T) {
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
		result := BigMystToDecimal(test.input)
		assert.Equal(t, test.get.String(), result.String(), fmt.Sprintf("test %d failed", i+1))
	}
}

func TestDecimalToBigMyst(t *testing.T) {
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
		result := DecimalToBigMyst(test.input)
		assert.Equal(t, test.get, result.String())
	}
}
