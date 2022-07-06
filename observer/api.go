package observer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// API is object which exposes observer API.
type API struct {
	client *http.Client
	url    string
}

const (
	hermesesEndpoint = "api/v1/observed/hermes"
)

// NewAPI returns a new API instance.
func NewAPI(url string, timeout time.Duration) *API {
	if timeout == 0 {
		timeout = time.Second * 10
	}

	return &API{
		client: &http.Client{
			Timeout: timeout,
		},
		url: strings.TrimSuffix(url, "/"),
	}
}

// HermesesResponse is returned from the observer hermeses endpoints and maps a slice of hermeses to its chain id.
type HermesesResponse map[int64][]HermesResponse

// HermesResponse describes a hermes.
type HermesResponse struct {
	HermesAddress common.Address `json:"hermes_address"`
	Operator      common.Address `json:"operator"`
	ChannelImpl   common.Address `json:"channel_impl"`

	Version int    `json:"version"`
	Fee     uint   `json:"fee"`
	URL     string `json:"url"`

	Active   bool `json:"active"`
	Approved bool `json:"approved"`
}

// HermesFilter can be used to filter values which would be returned from the observer.
type HermesFilter struct {
	Active   *bool
	Approved *bool
}

func (h *HermesFilter) toQuery() string {
	query := url.Values{}
	if h.Active != nil {
		query.Set("active", fmt.Sprint(*h.Active))
	}
	if h.Approved != nil {
		query.Set("approved", fmt.Sprint(*h.Approved))
	}

	return query.Encode()
}

// GetHermeses returns a map by chain of all known hermeses.
// An optional value `HermesFilter` is accepted which can be used to filter data on observer.
// `f` can be equal to `nil`, it will then be ignored.
func (a *API) GetHermeses(f *HermesFilter) (HermesesResponse, error) {
	if a.url == "" {
		return nil, fmt.Errorf("no observer url set")
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", a.url, hermesesEndpoint), nil)
	if err != nil {
		return nil, err
	}
	if f != nil {
		req.URL.RawQuery = f.toQuery()
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusOK:
		res := HermesesResponse{}
		err := json.NewDecoder(resp.Body).Decode(&res)
		return res, err
	default:
		return nil, fmt.Errorf("unknown status code from observer: %d", resp.StatusCode)
	}
}

// GetApprovedHermesAdresses returns a map by chain of all approved hermes addresses.
func (a *API) GetApprovedHermesAdresses() (map[int64][]common.Address, error) {
	t := true
	return a.GetHermesAdresses(&HermesFilter{
		Approved: &t,
	})
}

// GetHermesAdresses returns a map by chain of hermes addresses filtered by the given query.
func (a *API) GetHermesAdresses(f *HermesFilter) (map[int64][]common.Address, error) {
	resp, err := a.GetHermeses(f)

	if err != nil {
		return nil, err
	}
	res := make(map[int64][]common.Address)
	for chainId, hermesesResp := range resp {
		hermeses := make([]common.Address, len(hermesesResp))
		for i, h := range hermesesResp {
			hermeses[i] = h.HermesAddress
		}
		res[chainId] = hermeses
	}
	return res, nil
}

func (a *API) GetHermesData(chainId int64, hermesAddress common.Address) (*HermesResponse, error) {
	hermesesData, err := a.GetHermeses(nil)
	if err != nil {
		return nil, err
	}
	hermesesForChainId, ok := hermesesData[chainId]
	if !ok {
		return nil, fmt.Errorf("no hermeses for chain id %d", chainId)
	}

	for _, h := range hermesesForChainId {
		if h.HermesAddress == hermesAddress {
			return &h, nil
		}
	}

	return nil, fmt.Errorf("no data for hermes %s", hermesAddress.Hex())
}
