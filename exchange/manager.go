package exchange

import (
	"context"
	"errors"

	"github.com/mysteriumnetwork/msk/ctxl"
)

// MultiManager manages multi exchange APIs.
type MultiManager struct {
	apis []API
	log  ctxl.Logger
}

func NewMultiManager(exhangeAPIs []API) *MultiManager {
	return &MultiManager{
		apis: exhangeAPIs,
		log:  ctxl.New("exchange-multi-manager"),
	}
}

func (e *MultiManager) GetRateCacheWithFallback(ctx context.Context, coins []Coin, vsCurrencies []Currency) (PriceResponse, error) {
	for _, provider := range e.apis {
		ctx = ctxl.SetField(ctx, "provider", provider.GetName())

		res, err := provider.GetRateCacheWithFallback(ctx, coins, vsCurrencies)
		if err != nil {
			e.log(ctx).Err(err).Msg("requesting a rate from cache with fallback failed")
			continue
		}

		e.log(ctx).Info().Interface("result", res).Msg("get rate cache with fallback returned result")
		return res, nil
	}

	return nil, errors.New("all providers failed, exchange rate can't be provided")
}

func (e *MultiManager) GetRate(ctx context.Context, coins []Coin, vsCurrencies []Currency) (PriceResponse, error) {
	for _, provider := range e.apis {
		ctx = ctxl.SetField(ctx, "provider", provider.GetName())

		res, err := provider.GetRate(ctx, coins, vsCurrencies)
		if err != nil {
			e.log(ctx).Err(err).Msg("requesting a rate failed")
			continue
		}

		e.log(ctx).Info().Interface("result", res).Msg("get rate returned result")
		return res, nil
	}

	return nil, errors.New("all providers failed, exchange rate can't be provided")
}

func (e *MultiManager) GetRateCache(ctx context.Context, coins []Coin, vsCurrencies []Currency) (PriceResponse, error) {
	for _, provider := range e.apis {
		ctx = ctxl.SetField(ctx, "provider", provider.GetName())

		res, err := provider.GetRateCache(ctx, coins, vsCurrencies)
		if err != nil {
			e.log(ctx).Err(err).Msg("requesting a rate failed")
			continue
		}

		e.log(ctx).Info().Interface("result", res).Msg("get rate cache returned result")
		return res, nil
	}

	return nil, errors.New("all providers failed, exchange rate can't be provided")
}
