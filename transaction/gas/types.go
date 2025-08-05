package gas

import (
	"errors"
	"strings"
)

const EthereumMainnetChainID = 1
const PolygonMainnetChainID = 137
const AmoyPolygonTestnetChainID = 80002

const DefaultEtherscanEndpointURI = "https://api.etherscan.io/v2/"
const DefaultPolygonscanEndpointURI = "https://api.etherscan.io/v2/"

type EtherscanApiConfig struct {
	URL     string `json:"url"`
	ChainID uint64 `json:"chain_id"`
	ApiKey  string `json:"api_key"`
}

func NewEtherscanApiUrl(url string, chainID uint64, apiKey string) (EtherscanApiConfig, error) {
	cfg := EtherscanApiConfig{
		URL:     url,
		ChainID: chainID,
		ApiKey:  apiKey,
	}

	if err := cfg.Validate(); err != nil {
		return cfg, err
	}

	if !strings.HasSuffix(cfg.URL, "/") {
		cfg.URL += "/"
	}

	return cfg, nil
}

func NewMustEtherscanApiUrl(url string, chainID uint64, apiKey string) EtherscanApiConfig {
	cfg, err := NewEtherscanApiUrl(url, chainID, apiKey)
	if err != nil {
		panic(err)
	}
	return cfg
}

func (api EtherscanApiConfig) Validate() error {
	if api.URL == "" {
		return errors.New("url can't be empty")
	}

	if api.ChainID == 0 {
		return errors.New("chain_id can't be 0")
	}

	return nil
}
