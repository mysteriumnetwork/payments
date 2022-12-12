package client

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NonceTracker(t *testing.T) {
	trck := NewNonceTracker(&mockClient{})
	addr := crypto.HexToAddress("0x0")
	wg := sync.WaitGroup{}

	size := 1000
	nonces := make(chan uint64, size)
	for i := 0; i < size; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			nonce, err := trck.GetNonce(context.Background(), addr)
			assert.NoError(t, err)
			nonces <- nonce
		}()
	}

	wg.Wait()
	close(nonces)
	results := make(map[uint64]int)

	for v := range nonces {
		if prev, ok := results[v]; ok {
			results[v] = prev + 1
		} else {
			results[v] = 1
		}
	}

	for k, v := range results {
		assert.Equal(t, 1, v, fmt.Sprintf("%v is bad", k))
	}

}

type mockClient struct {
}

func (mc *mockClient) PendingNonceAt(ctx context.Context, account crypto.Address) (uint64, error) {
	return 1, nil
}
