package transaction

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/mysteriumnetwork/payments/transaction/gas"
)

type GasTracker struct {
	gs   GasStation
	opts map[int64]GasIncreaseOpts

	speed GasTrackerSpeed
}

type GasIncreaseOpts struct {
	Multiplier       float64
	PriceLimit       *big.Int
	IncreaseInterval time.Duration
}

type GasStation interface {
	GetGasPrices(chainID int64) (*gas.GasPrices, error)
}

type GasTrackerSpeed string

const (
	GasTrackerSpeedSlow   GasTrackerSpeed = "slow"
	GasTrackerSpeedMedium GasTrackerSpeed = "medium"
	GasTrackerSpeedFast   GasTrackerSpeed = "fast"
)

func NewGasTracker(gs GasStation, opts map[int64]GasIncreaseOpts, speed GasTrackerSpeed) *GasTracker {
	return &GasTracker{
		gs:    gs,
		opts:  opts,
		speed: speed,
	}
}

func (g *GasTracker) initialDeliveryGas(chainID int64) (*big.Int, error) {
	prices, err := g.gs.GetGasPrices(chainID)
	if err != nil {
		return nil, err
	}

	switch g.speed {
	case GasTrackerSpeedSlow:
		return prices.SafeLow, nil
	case GasTrackerSpeedMedium:
		return prices.Average, nil
	case GasTrackerSpeedFast:
		return prices.Fast, nil
	}

	return nil, errors.New("gas station not configured")
}

func (g *GasTracker) CanReceiveMoreGas(chainID int64, lastFillUpUTC time.Time) (bool, error) {
	opts, ok := g.opts[chainID]
	if !ok {
		return false, fmt.Errorf("no opts for chain %d", chainID)
	}

	receiveAfter := lastFillUpUTC.Add(opts.IncreaseInterval)
	return time.Now().UTC().After(receiveAfter), nil
}

func (g *GasTracker) RecalculateDeliveryGas(chainID int64, lastKnownGas *big.Int) (*big.Int, error) {
	if lastKnownGas == nil || lastKnownGas.Cmp(big.NewInt(0)) <= 0 {
		return g.initialDeliveryGas(chainID)
	}

	opts, ok := g.opts[chainID]
	if !ok {
		return nil, fmt.Errorf("no opts for chain %d", chainID)
	}

	newGasPrice, _ := new(big.Float).Mul(
		big.NewFloat(opts.Multiplier),
		new(big.Float).SetInt(lastKnownGas),
	).Int(nil)
	if newGasPrice == nil {
		return opts.PriceLimit, nil
	}

	if newGasPrice.Cmp(opts.PriceLimit) > 0 {
		if lastKnownGas.Cmp(opts.PriceLimit) < 0 {
			return opts.PriceLimit, nil
		}

		return nil, errors.New("max price reached, cannot increase")
	}

	return newGasPrice, nil
}