package gas

import (
	"math/big"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mysteriumnetwork/payments/v3/units"
	"github.com/stretchr/testify/assert"
)

func TestPolygonOfficial(t *testing.T) {
	m := newMockPolygonOfficialApi()
	m.run()
	defer m.stop()

	var oneGwei int64 = 1000000000
	es := NewMaticStation("http://localhost:8183", big.NewInt(200*oneGwei))

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

type mockPolygonOfficialApi struct {
	server   *http.Server
	response GasPrices
	Calls    int
}

func newMockPolygonOfficialApi() *mockPolygonOfficialApi {
	return &mockPolygonOfficialApi{
		response: GasPrices{},
		Calls:    0,
	}
}

func (m *mockPolygonOfficialApi) reset() {
	m.response = GasPrices{}
	r := gin.Default()
	r.GET("/", m.getGas)
	m.server.Handler = r
	m.Calls = 0
}

func (m *mockPolygonOfficialApi) setResponse(response GasPrices) {
	m.response = response
}

func (m *mockPolygonOfficialApi) setHandler(handler func(c *gin.Context)) {
	r := gin.Default()
	r.GET("/", handler)
	m.server.Handler = r
}

func (m *mockPolygonOfficialApi) run() {
	r := gin.Default()
	r.GET("/", m.getGas)

	m.server = &http.Server{
		Addr:    ":8183",
		Handler: r,
	}

	go func() {
		m.server.ListenAndServe()
	}()
}

func (m *mockPolygonOfficialApi) stop() {
	m.server.Close()
}

func (m *mockPolygonOfficialApi) getGas(c *gin.Context) {
	m.Calls++

	resp := maticGasPriceResp{
		BlockNumber:      100,
		BlockTime:        10,
		EstimatedBaseFee: units.BigIntWeiToFloatGwei(m.response.BaseFee),
		SafeLow: struct {
			MaxFee         float64 "json:\"maxFee\""
			MaxPriorityFee float64 "json:\"maxPriorityFee\""
		}{
			MaxPriorityFee: units.BigIntWeiToFloatGwei(m.response.SafeLow),
		},
		Standard: struct {
			MaxFee         float64 "json:\"maxFee\""
			MaxPriorityFee float64 "json:\"maxPriorityFee\""
		}{
			MaxPriorityFee: units.BigIntWeiToFloatGwei(m.response.Average),
		},
		Fast: struct {
			MaxFee         float64 "json:\"maxFee\""
			MaxPriorityFee float64 "json:\"maxPriorityFee\""
		}{
			MaxPriorityFee: units.BigIntWeiToFloatGwei(m.response.Fast),
		},
	}
	c.JSON(http.StatusOK, resp)
}
