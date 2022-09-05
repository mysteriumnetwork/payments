package observer

import (
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestObserverApi(t *testing.T) {
	m := newMockObserverApi()
	m.run()
	defer m.stop()

	obs := NewAPI("http://localhost:8185", time.Second)
	defaultResponse := HermesesResponse{
		1: []HermesResponse{
			{
				HermesAddress: common.HexToAddress("0x1"),
				Operator:      common.HexToAddress("0x2"),
				ChannelImpl:   common.HexToAddress("0x3"),
				Version:       1,
				Fee:           200,
				URL:           "http://hello.com",
				Active:        true,
				Approved:      true,
			},
			{
				HermesAddress: common.HexToAddress("0x4"),
				Operator:      common.HexToAddress("0x5"),
				ChannelImpl:   common.HexToAddress("0x6"),
				Version:       1,
				Fee:           200,
				URL:           "http://hello2.com",
				Active:        false,
				Approved:      true,
			},
			{
				HermesAddress: common.HexToAddress("0x7"),
				Operator:      common.HexToAddress("0x8"),
				ChannelImpl:   common.HexToAddress("0x9"),
				Version:       1,
				Fee:           200,
				URL:           "http://hello3.com",
				Active:        true,
				Approved:      false,
			},
			{
				HermesAddress: common.HexToAddress("0x10"),
				Operator:      common.HexToAddress("0x11"),
				ChannelImpl:   common.HexToAddress("0x12"),
				Version:       1,
				Fee:           200,
				URL:           "http://hello4.com",
				Active:        false,
				Approved:      false,
			},
		},
	}

	t.Run("get approved", func(t *testing.T) {
		m.setResponse(defaultResponse)
		resp, err := obs.GetApprovedHermesAdresses()
		assert.NoError(t, err)
		assert.Len(t, resp, 1)
		assert.Len(t, resp[1], 2)
		assert.Equal(t, common.HexToAddress("0x1"), resp[1][0])
		assert.Equal(t, common.HexToAddress("0x4"), resp[1][1])
		assert.Equal(t, 1, m.Calls)
		m.reset()
	})

	t.Run("get addresses", func(t *testing.T) {
		m.setResponse(defaultResponse)
		active := true
		resp, err := obs.GetHermesAdresses(&HermesFilter{
			Active: &active,
		})
		assert.NoError(t, err)
		assert.Len(t, resp, 1)
		assert.Len(t, resp[1], 2)
		assert.Equal(t, common.HexToAddress("0x1"), resp[1][0])
		assert.Equal(t, common.HexToAddress("0x7"), resp[1][1])
		assert.Equal(t, 1, m.Calls)
		m.reset()
	})

	t.Run("get data", func(t *testing.T) {
		m.setResponse(defaultResponse)
		resp, err := obs.GetHermesData(1, common.HexToAddress("0x4"))
		assert.NoError(t, err)
		assert.False(t, resp.Active)
		assert.True(t, resp.Approved)
		assert.Equal(t, common.HexToAddress("0x4"), resp.HermesAddress)
		assert.Equal(t, common.HexToAddress("0x5"), resp.Operator)
		assert.Equal(t, common.HexToAddress("0x6"), resp.ChannelImpl)
		assert.Equal(t, 1, resp.Version)
		assert.Equal(t, uint(200), resp.Fee)
		assert.Equal(t, "http://hello2.com", resp.URL)
		assert.Equal(t, 1, m.Calls)
		m.reset()
	})

	t.Run("get hermeses", func(t *testing.T) {
		m.setResponse(HermesesResponse{
			1: defaultResponse[1],
			2: []HermesResponse{
				{
					HermesAddress: common.HexToAddress("0x20"),
					Operator:      common.HexToAddress("0x21"),
					ChannelImpl:   common.HexToAddress("0x22"),
					Version:       4,
					Fee:           250,
					URL:           "http://hello20.com",
					Active:        false,
					Approved:      true,
				},
			},
		})
		approved := true
		active := false
		resp, err := obs.GetHermeses(&HermesFilter{
			Approved: &approved,
			Active:   &active,
		})
		assert.NoError(t, err)
		assert.Len(t, resp, 2)
		assert.Len(t, resp[1], 1)
		assert.Len(t, resp[2], 1)

		assert.False(t, resp[1][0].Active)
		assert.True(t, resp[1][0].Approved)
		assert.Equal(t, common.HexToAddress("0x4"), resp[1][0].HermesAddress)
		assert.Equal(t, common.HexToAddress("0x5"), resp[1][0].Operator)
		assert.Equal(t, common.HexToAddress("0x6"), resp[1][0].ChannelImpl)
		assert.Equal(t, 1, resp[1][0].Version)
		assert.Equal(t, uint(200), resp[1][0].Fee)
		assert.Equal(t, "http://hello2.com", resp[1][0].URL)

		assert.False(t, resp[2][0].Active)
		assert.True(t, resp[2][0].Approved)
		assert.Equal(t, common.HexToAddress("0x20"), resp[2][0].HermesAddress)
		assert.Equal(t, common.HexToAddress("0x21"), resp[2][0].Operator)
		assert.Equal(t, common.HexToAddress("0x22"), resp[2][0].ChannelImpl)
		assert.Equal(t, 4, resp[2][0].Version)
		assert.Equal(t, uint(250), resp[2][0].Fee)
		assert.Equal(t, "http://hello20.com", resp[2][0].URL)

		assert.Equal(t, 1, m.Calls)
		m.reset()
	})

	t.Run("handles errors", func(t *testing.T) {
		m.setHandler(func(c *gin.Context) {
			c.AbortWithStatus(http.StatusInternalServerError)
		})
		_, err := obs.GetHermeses(nil)
		assert.Error(t, err)
		m.reset()
	})
}

type mockObserverApi struct {
	server   *http.Server
	response HermesesResponse
	Calls    int
}

func newMockObserverApi() *mockObserverApi {
	return &mockObserverApi{
		response: make(HermesesResponse),
		Calls:    0,
	}
}

func (m *mockObserverApi) reset() {
	m.response = make(HermesesResponse)
	r := gin.Default()
	r.GET("/api/v1/observed/hermes", m.getHermeses)
	m.server.Handler = r
	m.Calls = 0
}

func (m *mockObserverApi) setResponse(response HermesesResponse) {
	m.response = response
}

func (m *mockObserverApi) setHandler(handler func(c *gin.Context)) {
	r := gin.Default()
	r.GET("/api/v1/observed/hermes", handler)
	m.server.Handler = r
}

func (m *mockObserverApi) run() {
	r := gin.Default()
	r.GET("/api/v1/observed/hermes", m.getHermeses)

	m.server = &http.Server{
		Addr:    ":8185",
		Handler: r,
	}

	go func() {
		m.server.ListenAndServe()
	}()
}

func (m *mockObserverApi) stop() {
	m.server.Close()
}

func (m *mockObserverApi) getHermeses(c *gin.Context) {
	m.Calls++

	var activeBool, approvedBool *bool = nil, nil
	active, ok := c.GetQuery("active")
	if ok {
		activeBoolVal, err := strconv.ParseBool(active)
		if err == nil {
			activeBool = &activeBoolVal
		}
	}
	approved, ok := c.GetQuery("approved")
	if ok {
		approvedBoolVal, err := strconv.ParseBool(approved)
		if err == nil {
			approvedBool = &approvedBoolVal
		}
	}

	resp := make(HermesesResponse)
	for chainId, hermeses := range m.response {
		hermesesResp := make([]HermesResponse, 0)
		for _, h := range hermeses {
			if activeBool != nil && *activeBool != h.Active {
				continue
			}
			if approvedBool != nil && *approvedBool != h.Approved {
				continue
			}
			hermesesResp = append(hermesesResp, h)
		}
		if len(hermesesResp) > 0 {
			resp[chainId] = hermesesResp
		}
	}

	c.JSON(http.StatusOK, resp)
}
