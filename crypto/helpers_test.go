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
	"math/big"
	"testing"
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
