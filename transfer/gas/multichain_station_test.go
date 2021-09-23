package gas

import (
	"errors"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultichain(t *testing.T) {
	t.Run("green path", func(t *testing.T) {
		defaultPrice := big.NewInt(10)
		mq := MultichainStation{
			1: []Station{NewStaticStation(defaultPrice), NewStaticStation(big.NewInt(600))},
		}

		prices, err := mq.GetGasPrices(1)
		assert.NoError(t, err)
		assert.Equal(t, defaultPrice, prices.Average)
		assert.Equal(t, defaultPrice, prices.SafeLow)
		assert.Equal(t, defaultPrice, prices.Fast)
	})
	t.Run("fallback", func(t *testing.T) {
		defaultPrice := big.NewInt(10)
		mq := MultichainStation{
			1: []Station{NewFailingStationMock(), NewStaticStation(defaultPrice)},
		}

		prices, err := mq.GetGasPrices(1)
		assert.NoError(t, err)
		assert.Equal(t, defaultPrice, prices.Average)
		assert.Equal(t, defaultPrice, prices.SafeLow)
		assert.Equal(t, defaultPrice, prices.Fast)
	})
	t.Run("chain does not exist", func(t *testing.T) {
		defaultPrice := big.NewInt(10)
		mq := MultichainStation{
			1: []Station{NewFailingStationMock(), NewStaticStation(defaultPrice)},
		}
		prices, err := mq.GetGasPrices(2)

		assert.Equal(t, fmt.Sprint(err), "no gas stations for chain 2")
		assert.Nil(t, prices)
	})
}

type FailingStationMock struct {
}

func NewFailingStationMock() *FailingStationMock {
	return &FailingStationMock{}
}

func (mock *FailingStationMock) GetGasPrices() (*GasPrices, error) {
	return nil, errors.New("mock error")
}
