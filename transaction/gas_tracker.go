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

type fees struct {
	Base *big.Int
	Tip  *big.Int
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

var errMaxPriceReached = errors.New("max price reached, cannot increase")

func NewGasTracker(gs GasStation, opts map[int64]GasIncreaseOpts, speed GasTrackerSpeed) *GasTracker {
	if speed == "" {
		speed = GasTrackerSpeedMedium
	}

	return &GasTracker{
		gs:    gs,
		opts:  opts,
		speed: speed,
	}
}

func (g *GasTracker) CanReceiveMoreGas(chainID int64, lastFillUpUTC time.Time) (bool, error) {
	opts, ok := g.opts[chainID]
	if !ok {
		return false, fmt.Errorf("no opts for chain %d", chainID)
	}

	receiveAfter := lastFillUpUTC.Add(opts.IncreaseInterval)
	return time.Now().UTC().After(receiveAfter), nil
}

func (g *GasTracker) ReceiveInitialGas(chainID int64) (*fees, error) {
	prices, err := g.gs.GetGasPrices(chainID)
	if err != nil {
		return nil, err
	}

	fees := &fees{
		Base: prices.BaseFee,
	}

	switch g.speed {
	case GasTrackerSpeedSlow:
		fees.Tip = prices.SafeLow
	case GasTrackerSpeedMedium:
		fees.Tip = prices.Average
	case GasTrackerSpeedFast:
		fees.Tip = prices.Fast
	default:
		return nil, errors.New("gas station not configured")
	}

	return fees, nil
}

func (g *GasTracker) RecalculateDeliveryGas(chainID int64, lastKnownTip *big.Int) (*fees, error) {
	if lastKnownTip == nil || lastKnownTip.Cmp(big.NewInt(0)) <= 0 {
		return g.ReceiveInitialGas(chainID)
	}

	opts, ok := g.opts[chainID]
	if !ok {
		return nil, fmt.Errorf("no opts for chain %d", chainID)
	}

	newTip := g.calculateNewPrice(chainID, lastKnownTip, opts.Multiplier)

	newFees, err := g.ReceiveInitialGas(chainID)
	if err != nil {
		return nil, fmt.Errorf("could not recalculate gas price: %w", err)
	}
	// If new tip is increased way above, use that.
	if newFees.Tip.Cmp(newTip) > 0 || newTip == nil {
		newTip = newFees.Tip
	}

	// Check that new tip is not exceeding our limits.
	if newTip.Cmp(opts.PriceLimit) > 0 {
		if lastKnownTip.Cmp(opts.PriceLimit) < 0 {
			return &fees{
				Base: newFees.Base,
				Tip:  opts.PriceLimit,
			}, nil
		}

		return nil, errMaxPriceReached
	}

	return &fees{
		Base: newFees.Base,
		Tip:  newTip,
	}, nil
}

func (g *GasTracker) calculateNewPrice(chainID int64, lastKnownTip *big.Int, multi float64) *big.Int {
	newTip, _ := new(big.Float).Mul(
		big.NewFloat(multi),
		new(big.Float).SetInt(lastKnownTip),
	).Int(nil)

	return newTip
}
