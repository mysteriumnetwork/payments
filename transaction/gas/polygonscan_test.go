package gas

import (
	"fmt"
	"math/big"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mysteriumnetwork/payments/v3/units"
	"github.com/stretchr/testify/assert"
)

func TestPolygonscan(t *testing.T) {
	m := newMockPolygonscanApi()
	m.run()
	defer m.stop()

	var oneGwei int64 = 1000000000
	es := NewPolygonscanStation(time.Second, "", "http://localhost:8184", big.NewInt(200*oneGwei))

	t.Run("get gas", func(t *testing.T) {
		m.setResponse(GasPrices{
			BaseFee: big.NewInt(35 * oneGwei),
			SafeLow: big.NewInt(31 * oneGwei),
			Average: big.NewInt(32 * oneGwei),
			Fast:    big.NewInt(33 * oneGwei),
		})
		gp, err := es.GetGasPrices()
		assert.NoError(t, err)

		assert.Equal(t, big.NewInt(35*oneGwei), gp.BaseFee)
		assert.Equal(t, big.NewInt(31*oneGwei), gp.SafeLow)
		assert.Equal(t, big.NewInt(32*oneGwei), gp.Average)
		assert.Equal(t, big.NewInt(33*oneGwei), gp.Fast)
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

type mockPolygonscanApi struct {
	server   *http.Server
	response GasPrices
	Calls    int
}

func newMockPolygonscanApi() *mockPolygonscanApi {
	return &mockPolygonscanApi{
		response: GasPrices{},
		Calls:    0,
	}
}

func (m *mockPolygonscanApi) reset() {
	m.response = GasPrices{}
	r := gin.Default()
	r.GET("/api", m.getGas)
	m.server.Handler = r
	m.Calls = 0
}

func (m *mockPolygonscanApi) setResponse(response GasPrices) {
	m.response = response
}

func (m *mockPolygonscanApi) setHandler(handler func(c *gin.Context)) {
	r := gin.Default()
	r.GET("/api", handler)
	m.server.Handler = r
}

func (m *mockPolygonscanApi) run() {
	r := gin.Default()
	r.GET("/api", m.getGas)

	m.server = &http.Server{
		Addr:    ":8184",
		Handler: r,
	}

	go func() {
		m.server.ListenAndServe()
	}()
}

func (m *mockPolygonscanApi) stop() {
	m.server.Close()
}

func (m *mockPolygonscanApi) getGas(c *gin.Context) {
	m.Calls++

	module := c.Query("module")
	if module != "gastracker" {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	action := c.Query("action")
	if action != "gasoracle" {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	resp := polygonscanGasPriceResponse{
		Status:  "1",
		Message: "hello",
		Result: polygonscanPriceResult{
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
