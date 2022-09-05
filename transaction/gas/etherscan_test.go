package gas

import (
	"fmt"
	"math/big"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mysteriumnetwork/payments/units"
	"github.com/stretchr/testify/assert"
)

func TestEtherscan(t *testing.T) {
	m := newMockEtherscanApi()
	m.run()
	defer m.stop()

	var oneGwei int64 = 1000000000
	es := NewEtherscanStation(time.Second, "", "http://localhost:8182", big.NewInt(200*oneGwei))

	t.Run("get gas", func(t *testing.T) {
		m.setResponse(GasPrices{
			BaseFee: big.NewInt(30 * oneGwei),
			SafeLow: big.NewInt(10 * oneGwei),
			Average: big.NewInt(20 * oneGwei),
			Fast:    big.NewInt(25 * oneGwei),
		})
		gp, err := es.GetGasPrices()
		assert.NoError(t, err)

		assert.Equal(t, big.NewInt(30*oneGwei), gp.BaseFee)
		assert.Equal(t, big.NewInt(10*oneGwei), gp.SafeLow)
		assert.Equal(t, big.NewInt(20*oneGwei), gp.Average)
		assert.Equal(t, big.NewInt(25*oneGwei), gp.Fast)
		assert.Equal(t, 1, m.Calls)

		m.setResponse(GasPrices{
			BaseFee: big.NewInt(50 * oneGwei),
			SafeLow: big.NewInt(150 * oneGwei),
			Average: big.NewInt(190 * oneGwei),
			Fast:    big.NewInt(220 * oneGwei),
		})
		gp, err = es.GetGasPrices()
		assert.NoError(t, err)

		assert.Equal(t, big.NewInt(50*oneGwei), gp.BaseFee)
		assert.Equal(t, big.NewInt(150*oneGwei), gp.SafeLow)
		assert.Equal(t, big.NewInt(190*oneGwei), gp.Average)
		//uppperbound
		assert.Equal(t, big.NewInt(200*oneGwei), gp.Fast)
		assert.Equal(t, 2, m.Calls)

		m.reset()
	})

	t.Run("handle error", func(t *testing.T) {
		m.setHandler(func(c *gin.Context) {
			c.AbortWithStatus(http.StatusInternalServerError)
		})
		_, err := es.GetGasPrices()
		assert.Error(t, err)

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
