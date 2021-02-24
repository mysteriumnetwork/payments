package matic

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
)

// BlockIncludedResponse represents the block included response.
type BlockIncludedResponse struct {
	HeaderBlockNumber string `json:"headerBlockNumber"`
	BlockNumber       string `json:"blockNumber"`
	Start             string `json:"start"`
	End               string `json:"end"`
	Proposer          string `json:"proposer"`
	Root              string `json:"root"`
	CreatedAt         string `json:"createdAt"`
	Message           string `json:"message"`
	Error             bool   `json:"error"`
}

// API performs requests to various matic APIS.
type API struct {
	network Network
	client  *http.Client
}

// Network represents the various matic networks
type Network string

const (
	// NetworkMumbai represents the matic mumbai testnet.
	NetworkMumbai Network = "MUMBAI"
	// NetworkMainnet represents the matic mainnet.
	NetworkMainnet Network = "MAINNET"
)

// NewAPI creates a new instance of matic API client.
func NewAPI(network Network, client *http.Client) *API {
	return &API{
		network: network,
		client:  client,
	}
}

func (a *API) getBlockIncludedURI(block *big.Int) string {
	if a.network == NetworkMainnet {
		return fmt.Sprintf("https://apis.matic.network/api/v1/matic/block-included/%v", block.String())
	}
	return fmt.Sprintf("https://apis.matic.network/api/v1/mumbai/block-included/%v", block.String())
}

// ErrBlockNotIncluded indicates that the block has not been included yet.
var ErrBlockNotIncluded = errors.New("block not yet included")

// GetBlockIncludedResponse fetches theblock inclusion response from matic api.
func (a *API) GetBlockIncludedResponse(block *big.Int) (BlockIncludedResponse, error) {
	resp, err := a.client.Get(a.getBlockIncludedURI(block))
	if err != nil {
		return BlockIncludedResponse{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return BlockIncludedResponse{}, err
	}

	res := BlockIncludedResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return res, err
	}
	
	if res.Error {
		return res, ErrBlockNotIncluded
	}

	return res, nil
}

func (a *API) getMerkleProofURI(start, end, number *big.Int) string {
	url := url.URL{
		Scheme: "https",
		Host:   "apis.matic.network",
		Path:   "api/v1/mumbai/fast-merkle-proof",
	}

	q := url.Query()
	q.Set("start", start.String())
	q.Set("end", end.String())
	q.Set("number", number.String())
	url.RawQuery = q.Encode()

	if a.network == NetworkMainnet {
		url.Path = "api/v1/matic/fast-merkle-proof"
		return url.String()
	}
	return url.String()
}

// FastMerkleProofResponse represent the fast merkle proof response.
type FastMerkleProofResponse struct {
	Proof string `json:"proof"`
}

// GetFastMerkleProof fetches the fast merkle proof from matic API.
func (a *API) GetFastMerkleProof(blockStart, blockEnd, blockNumber *big.Int) (FastMerkleProofResponse, error) {
	uri := a.getMerkleProofURI(blockStart, blockEnd, blockNumber)
	resp, err := a.client.Get(uri)
	if err != nil {
		return FastMerkleProofResponse{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return FastMerkleProofResponse{}, err
	}

	res := FastMerkleProofResponse{}
	err = json.Unmarshal(body, &res)
	return res, err
}
