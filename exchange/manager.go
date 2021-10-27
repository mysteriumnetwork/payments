package exchange

import (
	"errors"

	"github.com/rs/zerolog/log"
)

// MultiManager manages multi exchange APIs.
type MultiManager struct {
	apis []API
}

func NewMultiManager(exhangeAPIs []API) *MultiManager {
	return &MultiManager{
		apis: exhangeAPIs,
	}
}

func (e *MultiManager) GetRateCacheWithFallback(coins []Coin, vsCurrencies []Currency) (PriceResponse, error) {
	for _, provider := range e.apis {
		res, err := provider.GetRateCacheWithFallback(coins, vsCurrencies)
		if err != nil {
			log.Err(err).Str("provider", provider.GetName()).Msg("requesting a rate from cache with fallback failed")
			continue
		}

		return res, nil
	}

	return nil, errors.New("all providers failed, exchange rate can't be provided")
}

func (e *MultiManager) GetRate(coins []Coin, vsCurrencies []Currency) (PriceResponse, error) {
	for _, provider := range e.apis {
		res, err := provider.GetRate(coins, vsCurrencies)
		if err != nil {
			log.Err(err).Str("provider", provider.GetName()).Msg("requesting a rate failed")
			continue
		}

		return res, nil
	}

	return nil, errors.New("all providers failed, exchange rate can't be provided")
}

func (e *MultiManager) GetRateCache(coins []Coin, vsCurrencies []Currency) (PriceResponse, error) {
	for _, provider := range e.apis {
		res, err := provider.GetRateCache(coins, vsCurrencies)
		if err != nil {
			log.Err(err).Str("provider", provider.GetName()).Msg("requesting a rate from cache failed")
			continue
		}

		return res, nil
	}

	return nil, errors.New("all providers failed, exchange rate can't be provided")
}
