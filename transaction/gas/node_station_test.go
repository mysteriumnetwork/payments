package gas

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/mysteriumnetwork/payments/v3/client"
	"github.com/mysteriumnetwork/payments/v3/client/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNodeStation(t *testing.T) {
	cl := &mocks.EtherClientMock{
		SuggestGasPriceFunc: func(ctx context.Context) (*big.Int, error) {
			return big.NewInt(200), nil
		},
		HeaderByNumberFunc: func(ctx context.Context, number *big.Int) (*types.Header, error) {
			return &types.Header{
				Number: big.NewInt(20),
			}, nil
		},
	}
	getter := client.NewDefaultAddressableEthClientGetter("", cl)
	mbc := client.NewMultichainBlockchainClient(map[int64]client.BC{
		1: client.NewBlockchain(getter, time.Second),
	})
	ns := NewNodeStation(mbc, 1)
	t.Run("get gas", func(t *testing.T) {
		gp, err := ns.GetGasPrices()
		assert.NoError(t, err)
		assert.Equal(t, big.NewInt(200), gp.SafeLow)
		assert.Equal(t, big.NewInt(200), gp.Average)
		assert.Equal(t, big.NewInt(200), gp.Fast)
		assert.Equal(t, big.NewInt(1000000000), gp.BaseFee)
	})

	t.Run("handle error", func(t *testing.T) {
		cl.HeaderByNumberFunc = func(ctx context.Context, number *big.Int) (*types.Header, error) {
			return nil, fmt.Errorf("error")
		}
		_, err := ns.GetGasPrices()
		assert.Error(t, err)
	})
}
