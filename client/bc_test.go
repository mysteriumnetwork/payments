package client

import (
	"context"
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBC(t *testing.T) {
	t.Run("getter creation", func(t *testing.T) {
		cl := &EthMultiClient{}
		getter := NewDefaultEthClientGetter(cl)
		assert.Same(t, cl, getter.Client())
	})

	t.Run("connection failure error", func(t *testing.T) {
		assert.True(t, IsErrConnectionFailed(fmt.Errorf("some error: %w", context.DeadlineExceeded)))
		assert.True(t, IsErrConnectionFailed(fmt.Errorf("some error: %w", context.Canceled)))
		assert.False(t, IsErrConnectionFailed(fmt.Errorf("some error")))
		var err error = &net.DNSError{IsTimeout: false}
		assert.False(t, IsErrConnectionFailed(fmt.Errorf("some error: %w", err)))
		err = &net.DNSError{IsTimeout: true}
		assert.True(t, IsErrConnectionFailed(fmt.Errorf("some error: %w", err)))
	})
}
