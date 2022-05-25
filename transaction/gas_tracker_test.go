package transaction

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calculateOverpay(t *testing.T) {
	var (
		customType      DeliverableType = "some_type"
		customTypeUpper DeliverableType = "SOME_TYPE"
		customType2     DeliverableType = "some_type2"
	)

	for _, test := range []struct {
		name string
		opts map[int64]GasIncreaseOpts

		giveChainID int64
		giveType    DeliverableType
		givePrice   *big.Int

		getPrice *big.Int
	}{
		{
			name: "types match, chain matches, overpay by 50%",
			opts: map[int64]GasIncreaseOpts{
				1: {
					OverpayFor: []DeliverableType{
						customType,
					},
					OverpayByMul: 1.5,
				},
			},
			giveChainID: 1,
			giveType:    customType,
			givePrice:   big.NewInt(10005005000),
			getPrice:    big.NewInt(15007507500),
		},
		{
			name: "if types don't match because of uppercase letters it should still increase",
			opts: map[int64]GasIncreaseOpts{
				1: {
					OverpayFor: []DeliverableType{
						customTypeUpper,
					},
					OverpayByMul: 1.5,
				},
			},
			giveChainID: 1,
			giveType:    customType,
			givePrice:   big.NewInt(10005005000),
			getPrice:    big.NewInt(15007507500),
		},
		{
			name: "if multiplier is less than 1 it should be ignored",
			opts: map[int64]GasIncreaseOpts{
				1: {
					OverpayFor: []DeliverableType{
						customType,
					},
					OverpayByMul: 0.99,
				},
			},
			giveChainID: 1,
			giveType:    customType,
			givePrice:   big.NewInt(10005005000),
			getPrice:    big.NewInt(10005005000),
		},
		{
			name: "if type in transaction doesn't match opts it should be ignored",
			opts: map[int64]GasIncreaseOpts{
				1: {
					OverpayFor: []DeliverableType{
						customType,
					},
					OverpayByMul: 1.5,
				},
			},
			giveChainID: 1,
			giveType:    customType2,
			givePrice:   big.NewInt(10005005000),
			getPrice:    big.NewInt(10005005000),
		},
		{
			name: "if type in opts doesn't match it should be ignored",
			opts: map[int64]GasIncreaseOpts{
				1: {
					OverpayFor: []DeliverableType{
						customType2,
					},
					OverpayByMul: 1.5,
				},
			},
			giveChainID: 1,
			giveType:    customType,
			givePrice:   big.NewInt(10005005000),
			getPrice:    big.NewInt(10005005000),
		},
		{
			name:        "if no opts exist it should be ignored",
			opts:        map[int64]GasIncreaseOpts{},
			giveChainID: 1,
			giveType:    customType,
			givePrice:   big.NewInt(10005005000),
			getPrice:    big.NewInt(10005005000),
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			gt := &GasTracker{
				opts: test.opts,
			}
			got := gt.calculateOverpay(test.giveChainID, test.giveType, test.givePrice)
			assert.Equal(t, test.getPrice, got, "prices should match")
		})
	}
}
