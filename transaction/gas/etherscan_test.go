package gas

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"math/big"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mysteriumnetwork/payments/v3/units"
)

func TestEtherscan(t *testing.T) {
	m := newMockEtherscanApi()
	m.run()
	defer m.stop()

	var oneGwei int64 = 1000000000
	es := NewEtherscanStation(time.Second, NewMustEtherscanApiUrl("http://localhost:8182", 1, "fake"), big.NewInt(200*oneGwei))

	t.Run("get gas", func(t *testing.T) {
		m.setResponse(GasPrices{
			BaseFee: big.NewInt(30 * oneGwei),
			SafeLow: big.NewInt(10 * oneGwei),
			Average: big.NewInt(20 * oneGwei),
			Fast:    big.NewInt(25 * oneGwei),
		})
		gp, err := es.GetGasPrices()
		require.NoError(t, err)

		require.Equal(t, big.NewInt(30*oneGwei), gp.BaseFee)
		require.Equal(t, big.NewInt(10*oneGwei), gp.SafeLow)
		require.Equal(t, big.NewInt(20*oneGwei), gp.Average)
		require.Equal(t, big.NewInt(25*oneGwei), gp.Fast)
		require.Equal(t, 1, m.Calls)

		m.setResponse(GasPrices{
			BaseFee: big.NewInt(50 * oneGwei),
			SafeLow: big.NewInt(150 * oneGwei),
			Average: big.NewInt(190 * oneGwei),
			Fast:    big.NewInt(220 * oneGwei),
		})
		gp, err = es.GetGasPrices()
		require.NoError(t, err)

		require.Equal(t, big.NewInt(50*oneGwei), gp.BaseFee)
		require.Equal(t, big.NewInt(150*oneGwei), gp.SafeLow)
		require.Equal(t, big.NewInt(190*oneGwei), gp.Average)
		//uppperbound
		require.Equal(t, big.NewInt(200*oneGwei), gp.Fast)
		require.Equal(t, 2, m.Calls)

		m.reset()
	})

	t.Run("handle error", func(t *testing.T) {
		m.setHandler(func(c *gin.Context) {
			c.AbortWithStatus(http.StatusInternalServerError)
		})
		_, err := es.GetGasPrices()
		require.Error(t, err)

		m.reset()
	})
}

type mockEtherscanApi struct {
	server   *http.Server
	response GasPrices
	Calls    int
}

func newMockEtherscanApi() *mockEtherscanApi {
	return &mockEtherscanApi{
		response: GasPrices{},
		Calls:    0,
	}
}

func (m *mockEtherscanApi) reset() {
	m.response = GasPrices{}
	r := gin.Default()
	r.GET("/api", m.getGas)
	m.server.Handler = r
	m.Calls = 0
}

func (m *mockEtherscanApi) setResponse(response GasPrices) {
	m.response = response
}

func (m *mockEtherscanApi) setHandler(handler func(c *gin.Context)) {
	r := gin.Default()
	r.GET("/api", handler)
	m.server.Handler = r
}

func (m *mockEtherscanApi) run() {
	r := gin.Default()
	r.GET("/api", m.getGas)

	m.server = &http.Server{
		Addr:    ":8182",
		Handler: r,
	}

	go func() {
		m.server.ListenAndServe()
	}()
}

func (m *mockEtherscanApi) stop() {
	m.server.Close()
}

func (m *mockEtherscanApi) getGas(c *gin.Context) {
	m.Calls++

	module := c.Query("module")
	if module != "gastracker" {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	resp := etherscanGasPriceResponse{
		Status:  "1",
		Message: "hello",
		Result: gasPriceResult{
			LastBlock:    "100",
			GasUsedRatio: "0.95",

			SafeGasPrice:    fmt.Sprintf("%f", units.BigIntWeiToFloatGwei(m.response.SafeLow)),
			ProposeGasPrice: fmt.Sprintf("%f", units.BigIntWeiToFloatGwei(m.response.Average)),
			SuggestBaseFee:  fmt.Sprintf("%f", units.BigIntWeiToFloatGwei(m.response.BaseFee)),
			FastGasPrice:    fmt.Sprintf("%f", units.BigIntWeiToFloatGwei(m.response.Fast)),
		},
	}
	c.JSON(http.StatusOK, resp)
}
