package matic

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
)

type BlockIncludedResponse struct {
	HeaderBlockNumber string `json:"headerBlockNumber"`
	BlockNumber       string `json:"blockNumber"`
	Start             string `json:"start"`
	End               string `json:"end"`
	Proposer          string `json:"proposer"`
	Root              string `json:"root"`
	CreatedAt         string `json:"createdAt"`
	Message           string `json:"message"`
}

type BlockIncludedResponseError struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// API performs requests to various matic APIS.
type API struct {
	network Network
	client  *http.Client
}

type Network string

const (
	NetworkMumbai  Network = "MUMBAI"
	NetworkMainnet Network = "MAINNET"
)

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

var ErrBlockNotIncluded = errors.New("block not yet included")

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
		// this means we might have gotten an error response, try parsing that
		res := BlockIncludedResponseError{}
		err2 := json.Unmarshal(body, &res)
		if err2 != nil {
			return BlockIncludedResponse{}, fmt.Errorf("could not unmarshal error response %w", fmt.Errorf("could not unmarshal response from block include %w", err))
		}
		return BlockIncludedResponse{}, ErrBlockNotIncluded
	}
	return res, nil
}

func (a *API) getMerkleProofURI(start, end, number *big.Int) string {
	if a.network == NetworkMainnet {
		return fmt.Sprintf("https://apis.matic.network/api/v1/matic/fast-merkle-proof?start=%v&end=%v&number=%v", start.String(), end.String(), number.String())
	}
	return fmt.Sprintf("https://apis.matic.network/api/v1/mumbai/fast-merkle-proof?start=%v&end=%v&number=%v", start.String(), end.String(), number.String())

}

type FastMerkleProofResponse struct {
	Proof string `json:"proof"`
}

func (a *API) GetFastMerkleProof(blockStart, blockEnd, blockNumber *big.Int) (FastMerkleProofResponse, error) {
	resp, err := a.client.Get(a.getMerkleProofURI(blockStart, blockEnd, blockNumber))
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
	if err != nil {
		// this means we might have gotten an error response, try parsing that
		res := FastMerkleProofResponse{}
		err2 := json.Unmarshal(body, &res)
		if err2 != nil {
			return FastMerkleProofResponse{}, fmt.Errorf("could not unmarshal error response %w", fmt.Errorf("could not unmarshal response from fast merkle proof %w", err))
		}
		return FastMerkleProofResponse{}, ErrBlockNotIncluded
	}
	return res, nil
}
